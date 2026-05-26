package main

import (
	"context"
	"fmt"
	"os"

	lift "github.com/liftbridge-io/go-liftbridge/v2"
	"github.com/urfave/cli"

	"github.com/liftbridge-io/liftbridge/bench/common"
)

func main() {
	app := cli.NewApp()
	app.Name = "liftbridge-bench-producer"
	app.Usage = "Benchmark tool for Liftbridge message ingestion"
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

// Validate

// Connect to Liftbridge

// Create stream if requested

// Pre-generate messages (NOT timed)

// Setup stats

// Warn about pub-batch with concurrency

// Run benchmark

// Print results

// Delete stream if requested

func runBenchmark(
	ctx context.Context,
	client lift.Client,
	stream string,
	messages []common.PreparedMessage,
	concurrent int,
	pubBatch int,
	ackPolicy string,
	stats *common.Stats,
) error {
	_ = "STUB: not implemented"
	return nil
}

// Get ack policy option

// Progress counter

// Progress reporter

// Synchronous mode: one message at a time

// Async batch mode: send pubBatch messages, then wait for all acks

// Wait for all acks in this batch before sending next batch

func getAckPolicyOption(policy string) lift.MessageOption {
	_ = "STUB: not implemented"
	return *new(lift.MessageOption)
}

func normalizeServers(servers []string) []string { _ = "STUB: not implemented"; return nil }
