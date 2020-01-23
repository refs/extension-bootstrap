package command

import (
	"fmt"

	"github.com/micro/cli"
)

// ServerCommand returns a server subcommand
func ServerCommand() cli.Command {
	return cli.Command{
		Action: func(c *cli.Context) error {
			fmt.Println("oi")
			return nil
		},
	}
}
