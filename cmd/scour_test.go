package cmd

import (
	"bytes"
	"os"
	"os/exec"
	"testing"

	"github.com/gophergala2016/daemon/common"
)

func TestScourNoArguments(t *testing.T) {
	slash := string(os.PathSeparator)
	dir, _ := os.Getwd()
	base := dir + slash + ".." + slash
	pipeline := common.Pipeline{
		exec.Command(base+"daemon", "scour"),
	}

	contents := []byte(`{"included": ["daemon", "tech"], "excluded": ["satan"], "triggers": [{"command": "touch", "arguments": ["pass"], "wait": false}]}`)
	reader := bytes.NewReader(contents)
	_, err := pipeline.Run(reader)
	if err != nil {
		t.Error(err)
	}
}

func TestScourInvalidFeed(t *testing.T) {
	slash := string(os.PathSeparator)
	dir, _ := os.Getwd()
	base := dir + slash + ".." + slash
	pipeline := common.Pipeline{
		exec.Command(base+"daemon", "scour", "--feed", base+"does_not_exist.txt"),
	}

	contents := []byte(`{"included": ["daemon", "tech"], "excluded": ["satan"], "triggers": [{"command": "touch", "arguments": ["pass"], "wait": false}]}`)
	reader := bytes.NewReader(contents)
	_, err := pipeline.Run(reader)
	if err != nil {
		t.Error(err)
	}
}

func TestScourValidFeed(t *testing.T) {
	slash := string(os.PathSeparator)
	dir, _ := os.Getwd()
	base := dir + slash + ".." + slash
	pipeline := common.Pipeline{
		exec.Command(base+"daemon", "scour", "--feed", base+"feeds.txt"),
	}

	contents := []byte(`{"included": ["daemon", "tech"], "excluded": ["satan"], "triggers": [{"command": "touch", "arguments": ["pass"], "wait": false}]}`)
	reader := bytes.NewReader(contents)
	_, err := pipeline.Run(reader)
	if err != nil {
		t.Error(err)
	}
}
