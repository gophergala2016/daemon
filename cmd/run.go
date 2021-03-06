package cmd

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/codegangsta/cli"
	"github.com/gophergala2016/daemon/common"
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
		Flags: []cli.Flag{
			cli.IntFlag{
				Name:  "interval",
				Value: 15,
				Usage: "",
			},
		},
	}
)

// Run the main daemon processes.
func run(context *cli.Context) {
	initEnvironment()
	exe := common.CurrentExecutable()
	interval := time.Duration(context.Int("interval")) * time.Minute

	// At an interval scour the interwebs for stories to trigger events
	doEvery(interval, func() {
		log.Println("Scanning")

		feeds := feedsPath()
		stories := getStories(storyDirectory())
		for story := range stories {
			go func(s string) {
				// Read the story contents
				contents, err := ioutil.ReadFile(s)
				if err != nil {
					log.Printf("  [Fail] (%s) %s\n", s, err)
					return
				}

				// Create a chain of piped commands
				pipeline := common.Pipeline{
					exec.Command(exe, "scour", "--feeds", feeds),
					exec.Command(exe, "scourge"),
				}

				// Execute the chain while supplying the story as input
				reader := bytes.NewReader(contents)
				_, err = pipeline.Run(reader)
				if err != nil {
					log.Printf("  [Skip] (%s)\n", s)
					return
				}

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
