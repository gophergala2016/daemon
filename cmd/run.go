package cmd

import "github.com/codegangsta/cli"

var (
	executable string
	feeds      string
	CommandRun = cli.Command{
		Name:      "run",
		ShortName: "",
		Usage:     "",
		Action:    run,
		Flags:     []cli.Flag{},
	}
)

func run(context *cli.Context) {
	// Stub
}
