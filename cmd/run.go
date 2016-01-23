package cmd

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/codegangsta/cli"
)

const (
	// The interval between scouring the interwebs
	interval = 15 * time.Minute
)

var (
	// EnvironmentVariables is a key/value pairing storing default values.
	EnvironmentVariables = map[string]string{
		"DAEMON_STORY_DIR": "./stories",
		"DAEMON_FEEDS":     "./feeds.txt",
	}

	// CommandRun is the command to execute the main daemon processes.
	CommandRun = cli.Command{
		Name:      "run",
		ShortName: "",
		Usage:     "",
		Action:    run,
		Flags:     []cli.Flag{},
	}
)

// Run the main daemon processes.
func run(context *cli.Context) {
	initEnvironment()

	// At an interval scour the interwebs for stories to trigger events
	doEvery(interval, func() {
		log.Println("Scanning")

		stories := getStories(storyDirectory())
		for story := range stories {
			go func(s string) {
				// Read the story contents
				// TODO: Store file contents in `contents` variable
				_, err := ioutil.ReadFile(s)
				if err != nil {
					log.Printf("  [Fail] (%s) %s\n", s, err)
					return
				}

				// TODO: Create a chain of piped commands
				// TODO: Execute the chain while supplying the story as input

				log.Printf("  [Pass] (%s)\n", s)
			}(story)
		}
	})
}

// Repeat a piece of logic on a scheduled interval.
func doEvery(d time.Duration, f func()) {
	for _ = range time.Tick(d) {
		f()
	}
}

// Initialize the environment variables.
func initEnvironment() {
	log.Println("Initializing the Daemon")

	// Setup the environment variables with defaults if they are missing
	for key, value := range EnvironmentVariables {
		env := os.Getenv(key)
		if env == "" {
			os.Setenv(key, value)
		}
	}
}

// Find the file listing a feed per line to crawl.
// If the file does not exist then attempt to create it.
func feedsPath() string {
	path := os.Getenv("DAEMON_FEEDS")

	// Check for feeds directory
	if _, err := os.Stat(path); os.IsNotExist(err) {
		_, err = os.Create(path)
		if err != nil {
			return EnvironmentVariables["DAEMON_FEEDS"]
		}
	}

	return path
}

// Find the directory where all of the stories will be stored.
// If the directory does not exist then attempt to create it.
func storyDirectory() string {
	path := os.Getenv("DAEMON_STORY_DIR")

	// Check for story directory
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err = os.MkdirAll(path, 0755)
		if err != nil {
			return EnvironmentVariables["DAEMON_STORY_DIR"]
		}
	}

	return path
}

// Find all story files from the given directory.
func getStories(path string) chan string {
	c := make(chan string)

	go func() {
		filepath.Walk(path, func(p string, fi os.FileInfo, e error) error {
			// Bubble any errors which may have occured in a previous iteration
			if e != nil {
				return e
			}

			// Only use files with the ".story" suffix
			name := fi.Name()
			if strings.HasSuffix(name, ".story") {
				c <- p
			}

			return nil
		})

		defer close(c)
	}()

	return c
}
