package command

import (
	"context"
	"log"
	"os"
	"os/signal"
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

			stop := make(chan os.Signal, 1)

			gr.Add(func() error {
				signal.Notify(stop, syscall.SIGTERM)

				<-stop // empty the channel

				return nil
			}, func(_ error) {
				close(stop)
				cancel() // cancel context so the server stops
			})

			return gr.Run()
		},
	}
}
