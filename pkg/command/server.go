package command

import (
	"context"
	"log"
	"syscall"

	"github.com/micro/cli"
	"github.com/oklog/run"
	srv "github.com/refs/extension-bootstrap/pkg/server"
)

// ServerCommand returns a server subcommand
func ServerCommand() cli.Command {
	return cli.Command{
		Name:  "start",
		Usage: "starts the yeller service",
		Action: func(c *cli.Context) error {
			var (
				gr          = run.Group{}
				ctx, cancel = context.WithCancel(context.Background())
			)

			defer cancel()

			server, err := srv.Service(ctx)
			if err != nil {
				log.Panic(err)
			}

			gr.Add(func() error {
				return server.Run()
			}, func(_ error) {
				cancel()
			})

			gr.Add(run.SignalHandler(ctx, syscall.SIGKILL))
			return gr.Run()
		},
	}
}
