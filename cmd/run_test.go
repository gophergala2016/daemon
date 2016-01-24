package cmd

import (
	"os"
	"testing"
)

func TestInitEnvironemnt(t *testing.T) {
	initEnvironment()
	for key, value := range EnvironmentVariables {
		if os.Getenv(key) != value {
			t.Fail()
		}
	}
}

func TestFeedsPath(t *testing.T) {
	path := feedsPath()
	if _, err := os.Stat(path); os.IsNotExist(err) {
		t.Fail()
	}
}

func TestStoryDirectory(t *testing.T) {
	path := storyDirectory()
	if _, err := os.Stat(path); os.IsNotExist(err) {
		t.Fail()
	}
}

func TestDigestingStories(t *testing.T) {
	stories := getStories("../stories")

	found := 0
	for _ = range stories {
		found++
	}

	if found == 0 {
		t.Fail()
	}
}
