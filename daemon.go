package main

import (
	"os"
	"github.com/codegangsta/cli"
)

// Application variables for versioning
const (
	APP_NAME = "daemon"
	APP_USAGE = ""
	APP_VERSION = "0.0.0"
)

func main() {
	app := cli.NewApp()
	app.Name = APP_NAME
	app.Usage = APP_USAGE
	app.Version = APP_VERSION
	app.Commands = []cli.Command{}

	app.Run(os.Args)
}
