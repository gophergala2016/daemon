package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/codegangsta/cli"
	"github.com/gophergala2016/daemon/common"
)

var (
	// CommandScour will go through a list of feeds searching for contextual matches.
	CommandScour = cli.Command{
		Name:      "scour",
		ShortName: "",
		Usage:     "",
		Action:    scour,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "feeds",
				Value: "feeds.txt",
				Usage: "",
			},
		},
	}
)

// Scour a collection of feeds looking for a match to the given story.
func scour(context *cli.Context) {
	// Get the path to the file containing the feeds
	path := context.String("feeds")

	// Read our current story file
	story, err := common.ReadStory(os.Stdin)
	if err != nil {
		log.Println(err)
		return
	}

	// Open the feed
	f, err := os.Open(path)
	if err != nil {
		log.Println(err)
		return
	}
	defer f.Close()

	// Digest the feeds
	urls := []string{}
	bf := bufio.NewScanner(f)
	for bf.Scan() {
		u := bf.Text()
		comment := strings.Index(u, "//")

		if u != "" && comment != 0 {
			urls = append(urls, u)
		}
	}
	if err := bf.Err(); err != nil {
		log.Println(err)
		return
	}

	// Loop through each feed searching for a matching story
	c := make(chan common.Story)
	e := make(chan error)
	for _, url := range urls {
		go func(u string) {
			s, err := story.Find(u)
			if err != nil {
				e <- err
				return
			}
			c <- s
		}(url)
	}

	// Make sure everything has completed prior before quitting
	for i := 0; i < len(urls); i++ {
		select {
		case incoming := <-c:
			// Ignore JSON marshalling errors and move on to the next
			// potential match
			json, err := incoming.ToJSON()
			if err == nil {
				fmt.Fprintf(os.Stdout, "%s", json)
				return
			}

		case <-e:
			// TODO: Record the error
		}
	}

	// If nothing hit then report a failure
	log.Println("Could not find a story match")
}
