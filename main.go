//go:generate protoc -I=. -I=$GOPATH/src --gofast_out=. ./server/protocol/internal.proto

package main

import (
	"os"

	"github.com/urfave/cli"

	"github.com/liftbridge-io/liftbridge/server"
)

func main() {
	app := cli.NewApp()
	app.Name = "liftbridge"
	app.Usage = "Lightweight, fault-tolerant message streams"
	app.Version = server.Version
	app.Flags = getFlags()
	app.Action = start
	if err := app.Run(os.Args); err != nil {
		panic(err)
	}
}

func start(c *cli.Context) error {
	_ = "STUB: not implemented"
	// Read config from file if present.
	return nil
}

func overrideFromFlags(c *cli.Context, config *server.Config) error {
	_ = "STUB: not implemented"
	// Override with flags.
	return nil
}

func getFlags() []cli.Flag { _ = "STUB: not implemented"; return nil }

// NOTE: cannot use Value here as urfave/cli has another bug
// where it does not replace this value with the specified values but appends them:-(
// Value: &cli.StringSlice{nats.DefaultURL},

func normalizeNatsServers(natsServers []string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil,

		// urlfave.cli has issues with *Slice flags - it doesn't yet parse
		// command-line entries the same way as env vars, see
		// https://github.com/urfave/cli/pull/605
		// It has been around since Mar 2017 so don't hold your breath for a fix!
		// ... so we are manually splitting here for now.
		// We also need to handle possible multiple --nats-servers on the cli as this is supported.
		nil
}

// TODO: validate the server URL and return error?
