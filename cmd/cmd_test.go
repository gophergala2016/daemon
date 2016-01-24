package cmd

import (
	"io/ioutil"
	"log"
	"os"
)

var (
	stdin *os.File

	storyWithTriggers = []byte(`{
      "included": ["gopher", "gala"],
      "excluded": [""],
      "triggers": [{
        "command": "echo",
        "arguments": ["pass"],
        "wait": false
      }]
    }`)

	storyWithoutTriggers = []byte(`{
      "included": ["gopher", "gala"],
      "excluded": [""],
      "triggers": []
    }`)
)

func pushToStdin(contents []byte) *os.File {
	file, _ := ioutil.TempFile(os.TempDir(), "stdin")

	if file != nil {
		os.Stdin = file
	}

	os.Stdin.Write(contents)
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
