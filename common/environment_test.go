package common

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func TestCurrentExecutable(t *testing.T) {
	exe := CurrentExecutable()
	if !strings.HasSuffix(exe, ".test") {
		t.Fail()
	}
}

func TestStandardInput(t *testing.T) {
	in := StandardInput()
	contents, err := ioutil.ReadAll(in)
	if err != nil {
		t.Error(err)
	}
	if strings.TrimSpace(string(contents)) != "" {
		t.Fail()
	}
}

func TestStandardInputAsFile(t *testing.T) {
	os.Stdin = pushToStdin(storyWithTriggers)
	defer cleanupStdin()

	in := StandardInput()
	contents, err := ioutil.ReadAll(in)
	if err != nil {
		t.Error(err)
	}
	if strings.TrimSpace(string(contents)) != string(storyWithTriggers) {
		t.Fail()
	}
}
