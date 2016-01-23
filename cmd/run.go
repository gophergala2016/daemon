package cmd

import (
	"log"
	"os"
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
	// At an interval scour the interwebs for stories to trigger events
	doEvery(interval, func() {
		log.Println("Scanning")
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
func feedsPath() (string, error) {
	path := os.Getenv("DAEMON_FEEDS")

	// Check for feeds directory
	if _, err := os.Stat(path); os.IsNotExist(err) {
		_, err = os.Create(path)
		if err != nil {
			return EnvironmentVariables["DAEMON_FEEDS"], err
		}
	}

	return path, nil
}

// Find the directory where all of the stories will be stored.
// If the directory does not exist then attempt to create it.
func storyDirectory() (string, error) {
	path := os.Getenv("DAEMON_STORY_DIR")

	// Check for story directory
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err = os.MkdirAll(path, 0755)
		if err != nil {
			return EnvironmentVariables["DAEMON_STORY_DIR"], err
		}
	}

	return path, nil
}
