package cmd

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/codegangsta/cli"
)

func TestScourNoArguments(t *testing.T) {
	os.Args = []string{"test", "scour"}
	os.Stdin = pushToStdin(storyWithTriggers)
	defer cleanupStdin()

	app := cli.NewApp()
	app.Name = "test"
	app.Commands = []cli.Command{
		CommandScour,
	}

	err := app.Run(os.Args)
	if err != nil {
		t.Error(err)
	}
}

func TestScourInvalidFeed(t *testing.T) {
	os.Args = []string{"test", "scour", "--feeds", "nonexistent.txt"}
	os.Stdin = pushToStdin(storyWithTriggers)
	defer cleanupStdin()

	app := cli.NewApp()
	app.Name = "test"
	app.Commands = []cli.Command{
		CommandScour,
	}

	err := app.Run(os.Args)
	if err != nil {
		t.Error(err)
	}
}

func TestScourValidFeed(t *testing.T) {
	file, err := ioutil.TempFile(os.TempDir(), "feeds.txt")
	if err != nil {
		t.Fail()
	}

	defer os.Remove(file.Name())
	file.WriteString("http://golangweekly.com/rss/17nm799j")

	os.Args = []string{"test", "scour", "--feeds", file.Name()}
	os.Stdin = pushToStdin(storyWithTriggers)
	defer cleanupStdin()

	app := cli.NewApp()
	app.Name = "test"
	app.Commands = []cli.Command{
		CommandScour,
	}

	err = app.Run(os.Args)
	if err != nil {
		t.Error(err)
	}
}
