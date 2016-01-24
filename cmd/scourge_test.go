package cmd

import (
	"bytes"
	"os"
	"os/exec"
	"testing"

	"github.com/gophergala2016/daemon/common"
)

func TestScourgeWithTriggers(t *testing.T) {
	slash := string(os.PathSeparator)
	dir, _ := os.Getwd()
	base := dir + slash + ".." + slash
	pipeline := common.Pipeline{
		exec.Command(base+"daemon", "scourge"),
	}

	contents := []byte(`{"included": ["daemon", "tech"], "excluded": ["satan"], "triggers": [{"command": "touch", "arguments": ["pass"], "wait": false}]}`)
	reader := bytes.NewReader(contents)
	_, err := pipeline.Run(reader)
	if err != nil {
		t.Error(err)
	}
}

func TestScourgeWithNoTriggers(t *testing.T) {
	slash := string(os.PathSeparator)
	dir, _ := os.Getwd()
	base := dir + slash + ".." + slash
	pipeline := common.Pipeline{
		exec.Command(base+"daemon", "scourge"),
	}

	contents := []byte(`{"included": ["daemon", "tech"], "excluded": ["satan"], "triggers": []}`)
	reader := bytes.NewReader(contents)
	_, err := pipeline.Run(reader)
	if err != nil {
		t.Error(err)
	}
}
