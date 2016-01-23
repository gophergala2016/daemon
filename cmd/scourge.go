package cmd

import "github.com/codegangsta/cli"

var (
	// CommandScourge will react to a story hit executing linear processes.
	CommandScourge = cli.Command{
		Name:      "scourge",
		ShortName: "",
		Usage:     "",
		Action:    scourge,
		Flags:     []cli.Flag{},
	}
)

// Scourge executes the reaction to a story hit.
func scourge(context *cli.Context) {
	// TODO: Read the story
	// TODO: Check if there is anything to execute
	// TODO: Loop through the story triggers
}
