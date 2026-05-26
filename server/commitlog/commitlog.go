// Package commitlog provides an implementation for a file-backed write-ahead log.
package commitlog

import (
	"os"
	"sync"
	"time"

	"github.com/pkg/errors"

	"github.com/liftbridge-io/liftbridge/server/logger"
)

// ErrSegmentNotFound is returned if the segment could not be found.
var ErrSegmentNotFound = errors.New("segment not found")

// ErrIncorrectOffset is returned if the offset is incorrect. This is used in case Optimistic
// Concurrency Control is activated.
var ErrIncorrectOffset = errors.New("incorrect offset")

const (
	logFileSuffix               = ".log"
	indexFileSuffix             = ".index"
	hwFileName                  = "replication-offset-checkpoint"
	defaultMaxSegmentBytes      = 1073741824
	defaultHWCheckpointInterval = 5 * time.Second
	defaultCleanerInterval      = 5 * time.Minute
)

// commitLog implements the CommitLog interface, which is a durable write-ahead
// log.
type commitLog struct {
	readonly         int32 // Atomic flag
	deleteCleaner    *deleteCleaner
	compactCleaner   *compactCleaner
	name             string
	mu               sync.RWMutex
	hw               int64
	closed           chan struct{}
	segments         []*segment
	vActiveSegment   *segment
	hwWaiters        map[contextReader]chan bool
	leaderEpochCache *leaderEpochCache
	deleted          bool
	Options
}

// Options contains settings for configuring a commitLog.
type Options struct {
	Name                 string        // commitLog name
	Path                 string        // Path to log directory
	MaxSegmentBytes      int64         // Max bytes a Segment can contain before creating a new one
	MaxSegmentAge        time.Duration // Max time before a new log segment is rolled out.
	MaxLogBytes          int64         // Retention by bytes
	MaxLogMessages       int64         // Retention by messages
	MaxLogAge            time.Duration // Retention by age
	Compact              bool          // Run compaction on log clean
	CompactMaxGoroutines int           // Max number of goroutines to use in a log compaction
	CleanerInterval      time.Duration // Frequency to enforce retention policy
	HWCheckpointInterval time.Duration // Frequency to checkpoint HW to disk
	ConcurrencyControl   bool          // Optimistic Concurrency Control
	Logger               logger.Logger
}

// New creates a new CommitLog and starts a background goroutine which
// periodically checkpoints the high watermark to disk.
func New(opts Options) (CommitLog, error) { _ = "STUB: not implemented"; return *new(CommitLog), nil }

// After an unclean shutdown, the leader epoch checkpoint file could be
// ahead of the log (as the log is flushed asynchronously by default). To
// account for this, remove all entries from the leader epoch checkpoint
// file where the offset is greater than the log end offset.

// The earliest leader epoch may not be flushed during a hard failure.
// Recover it here.

func (l *commitLog) init() error {
	err := os.MkdirAll(l.Path, 0755)
	if err != nil {
		return errors.Wrap(err, "mkdir failed")
	}
	return nil
}

func (l *commitLog) open() error { _ = "STUB: not implemented"; return nil }

// If this file is an index file, make sure it has a corresponding .log
// file.

// Recover high watermark.

// Append writes the given batch of messages to the log and returns their
// corresponding offsets in the log. This will return ErrCommitLogReadonly if
// the log is in readonly mode.
func (l *commitLog) Append(msgs []*Message) ([]int64, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// AppendMessageSet writes the given message set data to the log and returns
// the corresponding offsets in the log. This can be called even if the log is
// in readonly mode to allow for reconciliation, e.g. when replicating from
// another log.
func (l *commitLog) AppendMessageSet(ms []byte) ([]int64, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (l *commitLog) append(segment *segment, ms []byte, entries []*entry) ([]int64, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Check if message is in a new leader epoch.

// If it is, we need to assign the epoch offset.

// NewestOffset returns the offset of the last message in the log or -1 if
// empty.
func (l *commitLog) NewestOffset() int64 { _ = "STUB: not implemented"; return 0 }

// OldestOffset returns the offset of the first message in the log or -1 if
// empty.
func (l *commitLog) OldestOffset() int64 { _ = "STUB: not implemented"; return 0 }

// EarliestOffsetAfterTimestamp returns the earliest offset whose timestamp is
// greater than or equal to the given timestamp.
func (l *commitLog) EarliestOffsetAfterTimestamp(timestamp int64) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Find the first segment whose base timestamp is greater than the given
// timestamp.

// EOF indicates there is no such segment, meaning the timestamp is
// beyond the end of the log so return the next assignable offset.

// Search the previous segment for the first entry whose timestamp is
// greater than or equal to the given timestamp. If this is the first
// segment, just search it.

// This indicates there are no entries in the segment whose timestamp
// is greater than or equal to the target timestamp. In this case, search
// the next segment if there is one. If there isn't, the timestamp is
// beyond the end of the log so return the next assignable offset.

// LatestOffsetBeforeTimestamp returns the latest offset whose timestamp is less
// than or equal to the given timestamp.
func (l *commitLog) LatestOffsetBeforeTimestamp(timestamp int64) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Find the first segment whose base timestamp is greater than the given
// timestamp.

// Search the previous segment for the first entry whose timestamp is
// greater than or equal to the given timestamp. If this is the first
// segment, just search it.

// if the given timestamp is before the start of the stream return an
// error.

// Find entry equal to or greater than the given timestamp.

// If it's an exact match, return the offset.

// Otherwise we want the previous offset.

// SetHighWatermark sets the high watermark on the log. All messages up to and
// including the high watermark are considered committed.
func (l *commitLog) SetHighWatermark(hw int64) { _ = "STUB: not implemented"; return }

// TODO: should we flush the HW to disk here?

// OverrideHighWatermark sets the high watermark on the log using the given
// value, even if the value is less than the current HW. This is used for unit
// testing purposes.
func (l *commitLog) OverrideHighWatermark(hw int64) { _ = "STUB: not implemented"; return }

// notifyHWChange signals all HW waiters to wake up because the HW has changed.
// This must be called within the log mutex.
func (l *commitLog) notifyHWChange() { _ = "STUB: not implemented"; return }

// notifyReadonly signals all HW waiters to wake up if the HW is caught up to
// the LEO because the log has become readonly. This must be called within the
// log mutex.
func (l *commitLog) notifyReadonly() { _ = "STUB: not implemented"; return }

// HW is caught up to LEO so notify HW waiters.

// waitForHW registers an HW waiter and returns a channel which will receive a
// bool either when the HW changes (false) or the log has become readonly
// (true).
func (l *commitLog) waitForHW(r contextReader, hw int64) <-chan bool {
	_ = "STUB: not implemented"
	return nil
}

// HW has changed since reader last checked so they can unblock now.

// Log is readonly and HW is caught up to LEO so return an error to reader.

// Reader needs to wait for HW to advance.

func (l *commitLog) removeHWWaiter(r contextReader) { _ = "STUB: not implemented"; return }

// HighWatermark returns the high watermark for the log.
func (l *commitLog) HighWatermark() int64 { _ = "STUB: not implemented"; return 0 }

// NewLeaderEpoch indicates the log is entering a new leader epoch.
func (l *commitLog) NewLeaderEpoch(epoch uint64) error { _ = "STUB: not implemented"; return nil }

// LastOffsetForLeaderEpoch returns the start offset of the first leader epoch
// larger than the provided one or the log end offset if the current epoch
// equals the provided one.
func (l *commitLog) LastOffsetForLeaderEpoch(epoch uint64) int64 {
	_ = "STUB: not implemented"
	return 0
}

// LastLeaderEpoch returns the latest leader epoch for the log.
func (l *commitLog) LastLeaderEpoch() uint64 { _ = "STUB: not implemented"; return 0 }

func (l *commitLog) activeSegment() *segment { _ = "STUB: not implemented"; return nil }

func (l *commitLog) close() error { _ = "STUB: not implemented"; return nil }

// Close closes each log segment file and stops the background goroutine
// checkpointing the high watermark to disk.
func (l *commitLog) Close() error { _ = "STUB: not implemented"; return nil }

// Delete closes the log and removes all data associated with it from the
// filesystem.
func (l *commitLog) Delete() error { _ = "STUB: not implemented"; return nil }

// IsDeleted returns true if the commit log has been deleted.
func (l *commitLog) IsDeleted() bool { _ = "STUB: not implemented"; return false }

// IsClosed returns true if the commit log was closed.
func (l *commitLog) IsClosed() bool { _ = "STUB: not implemented"; return false }

// Truncate removes all messages from the log starting at the given offset.
func (l *commitLog) Truncate(offset int64) error { _ = "STUB: not implemented"; return nil }

// Nothing to truncate.

// Delete all following segments.

// Delete the segment if its base offset is the target offset, provided
// it's not the first segment.

// Retain all preceding segments.

// Replace segment containing offset with truncated segment.

func (l *commitLog) Segments() []*segment { _ = "STUB: not implemented"; return nil }

// NotifyLEO registers and returns a channel which is closed when messages past
// the given log end offset are added to the log. If the given offset is no
// longer the log end offset, the channel is closed immediately. Waiter is an
// opaque value that uniquely identifies the entity waiting for data.
func (l *commitLog) NotifyLEO(waiter interface{}, expectedLEO int64) <-chan struct{} {
	_ = "STUB: not implemented"
	return nil
}

// SetReadonly marks the log as readonly. When in readonly mode, new messages
// cannot be added to the log with Append and committed readers will read up to
// the log end offset (LEO), if the HW allows so, and then will receive an
// ErrCommitLogReadonly error. This will unblock committed readers waiting for
// data if they are at the LEO. Readers will continue to block if the HW is
// less than the LEO. This does not affect uncommitted readers. Messages can
// still be written to the log with AppendMessageSet for reconciliation
// purposes, e.g. when replicating from another log.
func (l *commitLog) SetReadonly(readonly bool) { _ = "STUB: not implemented"; return }

// IsReadonly indicates if the log is in readonly mode.
func (l *commitLog) IsReadonly() bool { _ = "STUB: not implemented"; return false }

// IsConcurrencyControlEnabled indicates if the log should check for concurrency before appending messages
func (l *commitLog) IsConcurrencyControlEnabled() bool { _ = "STUB: not implemented"; return false }

// checkAndPerformSplit determines if a new log segment should be rolled out
// either because the active segment is full or MaxSegmentAge has passed since
// the first message was written to it. It then performs the split if eligible,
// returning any error resulting from the split. The returned bool indicates if
// a split was performed.
func (l *commitLog) checkAndPerformSplit() (bool, error) {
	_ = "STUB: not implemented"
	// Do this in a loop because segment splitting may fail due to a competing
	// thread performing the split at the same time. If this happens, we just
	// retry the check on the new active segment.
	return false, nil
}

// ErrSegmentExists indicates another thread has already performed
// the segment split, so reload the new active segment and check
// again.

func (l *commitLog) split(oldActiveSegment *segment) error { _ = "STUB: not implemented"; return nil }

// Do a CAS on the active segment to ensure no other threads have replaced
// it already. If this fails, it means another thread has already replaced
// it, so delete the new segment and return ErrSegmentExists.

// nolint: errcheck

func (l *commitLog) cleanerLoop() { _ = "STUB: not implemented"; return }

// Check to see if the active segment should be split.

// If we rolled a new segment, we don't need to run the cleaner since
// it already ran.

// Clean applies retention and compaction rules against the log, if applicable.
func (l *commitLog) Clean() error { _ = "STUB: not implemented"; return nil }

// New segments were added while cleaning. Rebase the new segments onto
// the cleaned ones.

// Update the leader epoch offset cache to account for deleted segments. If
// compaction ran, we need to regenerate the cache using the one returned
// from compaction.

// rebaseSegments adds the segments in from to the end of the slice of segments
// in to and adds any leader epoch offsets to the given leaderEpochCache.
func (l *commitLog) rebaseSegments(from, to []*segment, epochCache *leaderEpochCache) []*segment {
	_ = "STUB: not implemented"
	return nil

	// Rebase any leader epoch offsets also. We don't check the error returned
	// here because Rebase can't return an error since epochCache is not
	// file-backed. The epoch cache is nil if compaction didn't run, in which
	// case skip this.
}

// nolint: errcheck

// clean returns the cleaned segments and, if compaction ran, a
// *leaderEpochCache maintaining the start offset for each new leader epoch. If
// compaction did not run, the leaderEpochCache will be nil.
func (l *commitLog) clean(segments []*segment) ([]*segment, *leaderEpochCache, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (l *commitLog) checkpointHWLoop() { _ = "STUB: not implemented"; return }

func (l *commitLog) checkpointHW() error { _ = "STUB: not implemented"; return nil }
