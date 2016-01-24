package common

import (
	"os"
	"testing"
)

func TestFindingAStory(t *testing.T) {
	os.Stdin = pushToStdin(storyWithTriggers)
	defer cleanupStdin()

	story, err := ReadStory(StandardInput())
	if err != nil {
		t.Error(err)
	}

	story, err = story.Find("http://feeds.reuters.com/news/artsculture?format=xml")
	if err != nil && err.Error() != "Not found" {
		t.Error(err)
	}
}

func TestMatchAStory(t *testing.T) {
	os.Stdin = pushToStdin(storyWithTriggers)
	defer cleanupStdin()

	story, err := ReadStory(StandardInput())
	if err != nil {
		t.Error(err)
	}

	if !story.Match("the daemon in technology has started") {
		t.Fail()
	}
	if story.Match("satan daemon in technology has started") {
		t.Fail()
	}
}

func TestReadingAStory(t *testing.T) {
	os.Stdin = pushToStdin(storyWithTriggers)
	defer cleanupStdin()

	story, err := ReadStory(StandardInput())
	if err != nil {
		t.Error(err)
	}
	if len(story.Triggers) != 1 {
		t.Errorf("Expecting 1 trigger but found %q\n", len(story.Triggers))
	}
}
