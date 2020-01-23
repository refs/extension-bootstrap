package command

import (
	"fmt"

	"github.com/micro/cli"
)

// ServerCommand returns a server subcommand
func ServerCommand() cli.Command {
	return cli.Command{
		Name:  "start",
		Usage: "starts the yeller service",
		Action: func(c *cli.Context) error {
			// do all the channel magic to block
			fmt.Println("oi")
			return nil
		},
	}
}
