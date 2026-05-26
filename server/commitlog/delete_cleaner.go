package commitlog

import (
	"time"

	"github.com/liftbridge-io/liftbridge/server/logger"
)

// computeTTL calculates the age cutoff for messages when there is an age
// retention policy. This function exists for mocking purposes.
var computeTTL = func(age time.Duration) int64 {
	return time.Now().Add(-age).UnixNano()
}

// deleteCleanerOptions contains configuration settings for the DeleteCleaner.
type deleteCleanerOptions struct {
	Retention struct {
		Bytes    int64
		Messages int64
		Age      time.Duration
	}
	Logger logger.Logger
	Name   string
}

// deleteCleaner implements the delete cleanup policy which deletes old log
// segments based on the retention policy.
type deleteCleaner struct {
	deleteCleanerOptions
}

// newDeleteCleaner returns a new cleaner which enforces log retention
// policies by deleting segments.
func newDeleteCleaner(opts deleteCleanerOptions) *deleteCleaner {
	_ = "STUB: not implemented"
	return nil
}

// Clean will enforce the log retention policy by deleting old segments.
// Deletion only occurs at the segment granularity.
func (c *deleteCleaner) Clean(segments []*segment) ([]*segment, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Limit by age first.

// Next limit by number of messages.

// Lastly limit by number of bytes.

func (c *deleteCleaner) noRetentionLimits() bool { _ = "STUB: not implemented"; return false }

func (c *deleteCleaner) applyMessagesLimit(segments []*segment) ([]*segment, error) {
	_ = "STUB: not implemented"
	// We must retain at least the active segment.
	return nil, nil
}

// We start at the most recent segment and work our way backwards until we
// meet the retention size.

// Collect segments to delete

// Delete segments using mark-then-delete for consistency

func (c *deleteCleaner) applyBytesLimit(segments []*segment) ([]*segment, error) {
	_ = "STUB: not implemented"
	// We must retain at least the active segment.
	return nil, nil
}

// We start at the most recent segment and work our way backwards until we
// meet the retention size.

// Collect segments to delete

// Delete segments using mark-then-delete for consistency

func (c *deleteCleaner) applyAgeLimit(segments []*segment) ([]*segment, error) {
	_ = "STUB: not implemented"
	// We must retain at least the active segment.
	return nil, nil
}

// Collect all segments whose last-written timestamp is less than the TTL
// with the exception of the active (last) segment.

// Delete segments using mark-then-delete for consistency

// deleteSegments deletes the given segments using a mark-then-delete approach.
// This ensures that if deletion fails partway through, the segments are already
// removed from the read path (marked as deleted) and won't cause inconsistency.
// The actual file deletion can be retried on the next cleanup cycle.
func (c *deleteCleaner) deleteSegments(segments []*segment) error {
	_ = "STUB: not implemented"
	// Phase 1: Mark all segments as deleted to remove them from read path.
	// This is atomic per-segment and ensures readers won't see these segments.
	return nil
}

// Phase 2: Actually delete the files. If this fails partway through,
// the segments are already marked deleted and won't be visible to readers.
// Remaining files will be cleaned up on the next cleanup cycle.

// Continue trying to delete other segments
