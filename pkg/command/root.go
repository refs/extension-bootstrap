package command

import (
	"os"

	"github.com/micro/cli"
)

// RootCommand undocummented
func RootCommand() error {
	app := cli.App{
		Name:  "ocis-yeller",
		Usage: "contains a series of services",
		Commands: []cli.Command{
			ServerCommand(),
		},
	}

	return app.Run(os.Args)
}
