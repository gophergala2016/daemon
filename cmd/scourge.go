package cmd

import (
	"log"
	"os"
	"os/exec"

	"github.com/codegangsta/cli"
	"github.com/gophergala2016/daemon/common"
)

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
	// Read the story
	story, err := common.ReadStory(os.Stdin)
	if err != nil {
		log.Println(err)
	}

	// Check if there is anything to execute
	if len(story.Triggers) == 0 {
		return
	}

	// Loop through the story triggers
	for _, t := range story.Triggers {
		cmd := exec.Command(t.Command, t.Arguments...)
		err := cmd.Start()

		// If there was a problem starting the command and it is set
		// to wait then we fail hard
		if err != nil && t.Wait {
			log.Fatalf("Failed to execute [%s]", t.ToString())
		}

		// If we are to wait then we need to check its status and
		// fail hard if there was an issue
		if t.Wait {
			if err = cmd.Wait(); err != nil {
				log.Fatalf("Error executing [%s]", t.ToString())
			}
		}
	}
}
