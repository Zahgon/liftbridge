package main

import (
	"fmt"
	"os"

	lift "github.com/liftbridge-io/go-liftbridge/v2"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "liftbridge-bench-consumer"
	app.Usage = "Benchmark tool for Liftbridge message consumption"
	app.Version = "1.0.0"
	app.Flags = getFlags()
	app.Action = run
	if err := app.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func getFlags() []cli.Flag { _ = "STUB: not implemented"; return nil }

func run(c *cli.Context) error {
	_ = "STUB: not implemented"
	// Parse servers
	return nil
}

// Connect to Liftbridge

// Parse start option

// Setup stats

// Setup cancellation

// Message counter

// Start benchmark

// Subscribe

// Wait for completion

// Print results

func parseStartOption(pos string) (lift.SubscriptionOption, error) {
	_ = "STUB: not implemented"
	return *new(lift.SubscriptionOption), nil
}

func normalizeServers(servers []string) []string { _ = "STUB: not implemented"; return nil }
