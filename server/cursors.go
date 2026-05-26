package server

import (
	"context"
	"sync"
	"time"

	lru "github.com/hashicorp/golang-lru"
	"google.golang.org/grpc/status"
)

const (
	defaultCursorTimeout = 5 * time.Second
	cursorCacheSize      = 512
)

// cursorManager provides an API for managing consumer cursor positions for
// stream partitions. A cursorManager can only accept operations for requests
// that map to internal cursor partitions of which this server is the leader.
// Otherwise, it will return an error.
type cursorManager struct {
	*Server
	mu           sync.RWMutex
	cache        *lru.Cache
	disableCache bool // Used for testing purposes only
}

func newCursorManager(s *Server) *cursorManager {
	_ = "STUB: not implemented"
	// Ignoring error here because it's only returned if size is <= 0.
	return nil
}

// Initialize cursor management by creating the cursors stream if it doesn't
// yet exist and the configured number of partitions is greater than 0. This
// should be called when this node has been elected metadata leader to ensure
// the cursors stream exists.
func (c *cursorManager) Initialize() error { _ = "STUB: not implemented"; return nil }

// BecomePartitionLeader should be called when this server becomes the leader
// for any cursor partitions.
func (c *cursorManager) BecomePartitionLeader() {
	_ = "STUB: not implemented"
	// Clear the cache when we become leader to avoid serving potentially stale
	// cursors.
	return
}

// SetCursor stores a cursor position for a particular stream partition
// uniquely identified by an opaque string. This returns an error if persisting
// the cursor failed.
func (c *cursorManager) SetCursor(ctx context.Context, streamName, cursorID string, partitionID int32, offset int64) *status.Status {
	_ = "STUB: not implemented"
	return nil
}

// We lock on write to ensure ordering is consistent between the partition
// and in-memory cache even though the cache itself is thread-safe.

// Cache the offset.

// GetCursor returns the latest partition offset for the given cursor, if it
// exists.
func (c *cursorManager) GetCursor(ctx context.Context, streamName, cursorID string, partitionID int32) (int64, *status.Status) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Find the latest offset for the cursor in the log.

// Cache the offset.

func (c *cursorManager) getCursorsPartitionID(cursorKey []byte) (int32, *status.Status) {
	_ = "STUB: not implemented"
	return 0, nil
}

// TODO: Attempt to forward to partition leader.

func (c *cursorManager) getCursorKey(cursorID, streamName string, partitionID int32) []byte {
	_ = "STUB: not implemented"
	return nil
}

func (c *cursorManager) getLatestCursorOffset(ctx context.Context, cursorKey []byte, partition *partition) (
	int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// No cursors have been committed or the cursors partition is now empty so
// return -1.

// Use reverse subscription to find the latest cursor value efficiently.
// By reading from newest to oldest, the first matching key is the latest
// cursor value, allowing early exit instead of scanning all messages.

// When reading in reverse, the first match is the latest value.

// Reached the oldest message without finding the cursor key.

// ResourceExhausted means we've read all messages (reached end of
// reverse iteration).
