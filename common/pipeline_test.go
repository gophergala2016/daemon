package common

import (
	"bytes"
	"os/exec"
	"strings"
	"testing"
)

func TestPipelineRunWithOneProgram(t *testing.T) {
	expect := "blah"
	pipeline := Pipeline{
		exec.Command("echo", expect),
	}

	reader := bytes.NewReader([]byte("test input"))
	output, err := pipeline.Run(reader)
	if err != nil {
		t.Error(err)
	}
	if strings.TrimSpace(string(output)) != expect {
		t.Errorf("Expecting %q but was %q\n", expect, output)
	}
}

func TestPipelineRunWithMultiplePrograms(t *testing.T) {
	expect := "8"
	pipeline := Pipeline{
		exec.Command("echo", "8\n2\n3"),
		exec.Command("grep", "8"),
	}

	reader := bytes.NewReader([]byte("test input"))
	output, err := pipeline.Run(reader)
	if err != nil {
		t.Error(err)
	}
	if strings.TrimSpace(string(output)) != expect {
		t.Errorf("Expecting %q but was %q\n", expect, output)
	}
}
