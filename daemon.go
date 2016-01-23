package main

import (
	"os"

	"github.com/codegangsta/cli"
	"github.com/gophergala2016/daemon/cmd"
)

// Application variables for versioning
const (
	AppName    = "daemon"
	AppUsage   = ""
	AppVersion = "0.0.0"
)

func main() {
	app := cli.NewApp()
	app.Name = AppName
	app.Usage = AppUsage
	app.Version = AppVersion
	app.Commands = []cli.Command{
		cmd.CommandRun,
		cmd.CommandScour,
	}

	app.Run(os.Args)
}
