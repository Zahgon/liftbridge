package commitlog

import (
	"sync"

	"github.com/liftbridge-io/liftbridge/server/logger"
)

const defaultCompactMaxGoroutines = 10

// compactCleanerOptions contains configuration settings for the
// compactCleaner.
type compactCleanerOptions struct {
	Logger        logger.Logger
	Name          string
	MaxGoroutines int
}

// compactCleaner implements the compaction policy which replaces segments with
// compacted ones, i.e. retaining only the last message for a given key.
type compactCleaner struct {
	compactCleanerOptions
}

// NewCompactCleaner returns a new cleaner which performs log compaction by
// rewriting segments such that they contain only the last message for a given
// key.
func newCompactCleaner(opts compactCleanerOptions) *compactCleaner {
	_ = "STUB: not implemented"
	return nil
}

// Compact performs log compaction by rewriting segments such that they contain
// only the last message for a given key. Compaction is applied to all segments
// up to but excluding the active (last) segment or the provided HW, whichever
// comes first. This returns the compacted segments and a leaderEpochCache
// containing the earliest offsets for each leader epoch or nil if nothing was
// compacted.
func (c *compactCleaner) Compact(hw int64, segments []*segment) ([]*segment,
	*leaderEpochCache, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

type keyOffset struct {
	sync.RWMutex
	offset int64
}

func (k *keyOffset) set(offset int64) { _ = "STUB: not implemented"; return }

func (k *keyOffset) get() int64 { _ = "STUB: not implemented"; return 0 }

func (c *compactCleaner) compact(hw int64, segments []*segment) ([]*segment,
	*leaderEpochCache, int, error) {
	_ = "STUB: not implemented"

	// Compact messages up to the last segment or HW, whichever is first, by
	// scanning keys and retaining only the latest.
	// TODO: Implement option for configuring minimum compaction lag.
	return nil, nil, 0, nil
}

// Write new segments. Skip the last segment since we will not compact it.
// TODO: Join segments that are below the bytes limit.

// Add the last segment back in to the compacted list.

// Maintain start offset for each new leader epoch for the last segment.

func (c *compactCleaner) cleanSegment(seg *segment, keyOffsets *sync.Map, hw int64,
	epochCache *leaderEpochCache) (*segment, int, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

// Retain all messages with no keys and last message for each key.
// Also retain all messages after the HW.

// Maintain start offset for each new leader epoch.

// If the new segment is empty, remove it along with the old one.

// Otherwise replace the old segment with the compacted one.

func (c *compactCleaner) scanKeys(hw int64, segments []*segment) *sync.Map {
	_ = "STUB: not implemented"
	return nil
}

func (c *compactCleaner) scanSegments(hw int64, ch <-chan *segment, wg *sync.WaitGroup, keyOffsets *sync.Map) {
	_ = "STUB: not implemented"
	return
}

func cleanupEmptySegment(new, old *segment) error {
	_ = "STUB: not implemented"
	// Delete the new segment if it's empty.
	return nil
}

// Also delete the old segment since it's been compacted. Set the replaced
// flag since this is in the read path.
