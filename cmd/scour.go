package cmd

import "github.com/codegangsta/cli"

var (
	// CommandScour will go through a list of feeds searching for contextual matches.
	CommandScour = cli.Command{
		Name:      "scour",
		ShortName: "",
		Usage:     "",
		Action:    scour,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "feed",
				Value: "feeds.txt",
				Usage: "",
			},
		},
	}
)

// Scour a collection of feeds looking for a match to the given story.
func scour(context *cli.Context) {
	// TODO: Get the path to the file containing the feeds
	// TODO: Read our current story file
	// TODO: Open the feed
	// TODO: Digest the feeds
	// TODO: Loop through each feed searching for a matching story
	// TODO: Make sure everything has completed prior before quitting
}
