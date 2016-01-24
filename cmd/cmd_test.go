package cmd

import (
	"io/ioutil"
	"log"
	"os"
	"testing"
)

const (
	storyWithTriggers string = `{
      "included": ["daemon", "tech"],
      "excluded": ["satan"],
      "triggers": [{
        "command": "touch",
        "arguments": ["pass"],
        "wait": false
      }]
    }`

	storyWithoutTriggers string = `{
      "included": ["daemon", "tech"],
      "excluded": ["satan"],
      "triggers": [{
        "command": "touch",
        "arguments": ["pass"],
        "wait": false
      }]
    }`
)

var stdin *os.File

func TestMain(m *testing.M) {
	if stdin == nil {
		stdin = pushToStdin(storyWithTriggers)
	}
	os.Stdin = stdin

	r := m.Run()

	cleanupStdin()
	os.Exit(r)
}

func pushToStdin(contents string) *os.File {
	file, _ := ioutil.TempFile(os.TempDir(), "stdin")

	if file != nil {
		file.WriteString(contents)
		file.Sync()
		return file
	}

	return os.Stdin
}

func cleanupStdin() {
	stats, err := os.Stdin.Stat()
	if err != nil {
		log.Println(err)
	}

	if stats.Size() > 0 {
		os.Remove(os.Stdin.Name())
	}
}
