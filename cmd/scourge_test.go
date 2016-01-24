package cmd

import (
	"os"
	"testing"

	"github.com/codegangsta/cli"
)

func TestScourgeWithTriggers(t *testing.T) {
	os.Args = []string{"test", "scourge"}
	app := cli.NewApp()
	app.Name = "test"
	app.Commands = []cli.Command{
		CommandScourge,
	}

	err := app.Run(os.Args)
	if err != nil {
		t.Error(err)
	}
}

func TestScourgeWithNoTriggers(t *testing.T) {
	os.Args = []string{"test", "scourge"}
	app := cli.NewApp()
	app.Name = "test"
	app.Commands = []cli.Command{
		CommandScourge,
	}

	err := app.Run(os.Args)
	if err != nil {
		t.Error(err)
	}
}
