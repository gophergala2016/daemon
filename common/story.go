package common

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"strings"

	"github.com/gophergala2016/daemon/reader"
)

// Trigger represents a command to be executed on a story hit.
type Trigger struct {
	Command   string   `json:"command"`
	Arguments []string `json:"arguments"`
	Wait      bool     `json:"wait"`
}

// ToString pretty prints the trigger.
func (t *Trigger) ToString() string {
	if t == nil || t.Command == "" {
		return ""
	}

	return fmt.Sprintf("%s %s", t.Command, strings.Join(t.Arguments, " "))
}

// Story contains the criterion to search feeds and find matches based
// on simple inclusion/exclusion keywords. It also lists the events to
// be triggered once a match is found.
//
// A story is what provides the flexibility to the solution giving it
// a wide range of possible usages.
type Story struct {
	Included []string    `json:"included"`
	Excluded []string    `json:"excluded"`
	Article  reader.Item `json:"article,omitempty"`
	Triggers []Trigger   `json:"triggers,omitempty"`
}

// FromJSON parses JSON in to a story.
func (s *Story) FromJSON(content string) error {
	return json.Unmarshal([]byte(content), s)
}

// ToJSON converts a story to JSON format.
func (s *Story) ToJSON() ([]byte, error) {
	return json.Marshal(s)
}

// Find will read the feed and check if the story matches the criterion.
func (s Story) Find(url string) (Story, error) {
	feed, err := reader.GetFeed(url)
	if err != nil {
		log.Println(err)
		return s, err
	}

	for _, article := range feed.Items {
		t := fmt.Sprintf("%s %s", article.Title, string(article.Description))
		if s.Match(t) {
			s.Article = article
			return s, nil
		}
	}

	return s, errors.New("Not found")
}

// Match validates whether the content is within the story parameters.
func (s Story) Match(content string) bool {
	return s.validateInclusions(content) && s.validateExclusions(content)
}

// validateInclusions determines if all `included` words are present.
func (s Story) validateInclusions(content string) bool {
	hits := 0

	content = strings.ToLower(content)
	for _, word := range s.Included {
		if strings.Contains(content, strings.ToLower(word)) {
			hits++
		}
	}

	return hits == len(s.Included)
}

// validateExclusions determines if all `excluded` words are not present.
func (s Story) validateExclusions(content string) bool {
	hits := 0

	content = strings.ToLower(content)
	for _, word := range s.Excluded {
		if !strings.Contains(content, strings.ToLower(word)) {
			hits++
		}
	}

	return hits == len(s.Excluded)
}

// ReadStory reads from a stream of JSON in to a story.
func ReadStory(reader io.Reader) (Story, error) {
	// Create a new scanner to read the JSON
	json := ""
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		json += scanner.Text()
	}

	// Check for any scanner errors
	if err := scanner.Err(); err != nil {
		return Story{}, err
	}

	// Parse the JSON in to a story
	story := Story{}
	err := story.FromJSON(json)
	return story, err
}
