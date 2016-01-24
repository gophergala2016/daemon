package cmd

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/codegangsta/cli"
)

func TestScourNoArguments(t *testing.T) {
	os.Args = []string{"test", "scour"}
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
	os.Args = []string{"test", "scour"}
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
	os.Args = []string{"test", "scour"}
	app := cli.NewApp()
	app.Name = "test"
	app.Commands = []cli.Command{
		CommandScour,
	}

	file, err := ioutil.TempFile(os.TempDir(), "feeds.txt")
	if err != nil {
		t.Fail()
	}
	defer os.Remove(file.Name())
	file.WriteString("http://feeds.reuters.com/news/artsculture?format=xml")

	err = app.Run(os.Args)
	if err != nil {
		t.Error(err)
	}
}
