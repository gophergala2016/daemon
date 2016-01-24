package cmd

import "testing"

func TestScourNoArguments(t *testing.T) {
	// app := cli.NewApp()
	// app.Writer = ioutil.Discard
	// app.Commands = []cli.Command{
	// 	CommandScour,
	// }

	// err := app.Run([]string{common.CurrentExecutable(), CommandScour.Name})
	// if err != nil {
	// 	t.Error(err)
	// }
}

func TestScourInvalidFeed(t *testing.T) {
	// app := cli.NewApp()
	// app.Writer = ioutil.Discard
	// app.Commands = []cli.Command{
	// 	CommandScour,
	// }

	// err := app.Run([]string{common.CurrentExecutable(), CommandScour.Name, "--feed", "babble.io"})
	// if err != nil {
	// 	t.Error(err)
	// }
}

func TestScourValidFeed(t *testing.T) {
	// app := cli.NewApp()
	// app.Writer = ioutil.Discard
	// app.Commands = []cli.Command{
	// 	CommandScour,
	// }

	// err := app.Run([]string{common.CurrentExecutable(), CommandScour.Name, "--feeds", "http://rss.cnn.com/rss/cnn_latest.rss"})
	// if err != nil {
	// 	t.Error(err)
	// }
}
