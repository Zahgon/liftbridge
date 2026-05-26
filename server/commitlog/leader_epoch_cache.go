package commitlog

import (
	"io"
	"sync"

	"github.com/liftbridge-io/liftbridge/server/logger"
)

const (
	leaderEpochFileName = "leader-epoch-checkpoint"
	leaderEpochFileV0   = 0
)

// epochOffset contains the start offset for a given leader epoch.
type epochOffset struct {
	leaderEpoch uint64
	startOffset int64
}

type leaderEpochCache struct {
	epochOffsets   []*epochOffset
	mu             sync.RWMutex
	checkpointFile string
	name           string
	log            logger.Logger
}

func newLeaderEpochCacheNoFile(name string, log logger.Logger) *leaderEpochCache {
	_ = "STUB: not implemented"
	return nil
}

func newLeaderEpochCache(name, path string, log logger.Logger) (*leaderEpochCache, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Load epoch offsets from file.

// Assign the given leader epoch to the given offset. Once assigned, an epoch
// cannot be reassigned.
func (l *leaderEpochCache) Assign(epoch uint64, offset int64) error {
	_ = "STUB: not implemented"
	return nil
}

// LastOffsetForLeaderEpoch returns the start offset of the first leader epoch
// larger than the provided one or -1 if the current epoch equals the provided
// one.
func (l *leaderEpochCache) LastOffsetForLeaderEpoch(epoch uint64) int64 {
	_ = "STUB: not implemented"
	return 0
}

// LastLeaderEpoch returns the latest leader epoch for the log.
func (l *leaderEpochCache) LastLeaderEpoch() uint64 { _ = "STUB: not implemented"; return 0 }

// ClearLatest removes all leader epoch entries from the cache with start
// offsets greater than or equal to the given offset.
func (l *leaderEpochCache) ClearLatest(offset int64) error { _ = "STUB: not implemented"; return nil }

// ClearEarliest searches for the oldest leader epoch < offset, updates the
// saved epoch offset to the given offset, then removes any previous epoch
// entries.
func (l *leaderEpochCache) ClearEarliest(offset int64) error { _ = "STUB: not implemented"; return nil }

// If the offset is less than the earliest offset remaining, add
// previous epoch back but with an updated offset.

// Rebase adds the leader epoch offsets from the given leaderEpochCache
// starting at the given offset.
func (l *leaderEpochCache) Rebase(from *leaderEpochCache, offset int64) error {
	_ = "STUB: not implemented"
	return nil
}

// Replace the contents of the cache using the provided one.
func (l *leaderEpochCache) Replace(from *leaderEpochCache) error {
	_ = "STUB: not implemented"
	return nil
}

func (l *leaderEpochCache) earliestOffset() int64 { _ = "STUB: not implemented"; return 0 }

func (l *leaderEpochCache) latestEpoch() uint64 { _ = "STUB: not implemented"; return 0 }

func (l *leaderEpochCache) latestOffset() int64 { _ = "STUB: not implemented"; return 0 }

func (l *leaderEpochCache) findEpoch(epoch uint64) *epochOffset {
	_ = "STUB: not implemented"
	return nil
}

func (l *leaderEpochCache) assign(epoch uint64, offset int64) error {
	_ = "STUB: not implemented"
	return nil
}

// flush writes the cached epoch offsets to disk in the following format:
//
// v0:
// version
// num_entries
// leader_epoch start_offset
// leader_epoch start_offset
// ...
func (l *leaderEpochCache) flush() error { _ = "STUB: not implemented"; return nil }

func (l *leaderEpochCache) warn(epoch, latestEpoch uint64, offset, latestOffset int64) {
	_ = "STUB: not implemented"
	return
}

func (l *leaderEpochCache) epochChangeMsg(newEpoch, lastEpoch uint64, newOffset, lastOffset int64) string {
	_ = "STUB: not implemented"
	return ""
}

// readLeaderEpochOffsets reads the contents of the leader epoch checkpoint
// file, which is of the following form:
//
// v0:
// version
// num_entries
// leader_epoch start_offset
// leader_epoch start_offset
// ...
func readLeaderEpochOffsets(file io.Reader) ([]*epochOffset, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Duplicate entry.
