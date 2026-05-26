package commitlog

import (
	"io"
	"os"
	"sync"
	"time"

	"github.com/pkg/errors"
)

const (
	fileFormat      = "%020d%s"
	logSuffix       = ".log"
	cleanedSuffix   = ".cleaned"
	truncatedSuffix = ".truncated"
	indexSuffix     = ".index"
)

var (
	// ErrEntryNotFound is returned when a segment search cannot find a
	// specific entry.
	ErrEntryNotFound = errors.New("entry not found")

	// ErrSegmentClosed is returned on reads/writes to a closed segment.
	ErrSegmentClosed = errors.New("segment has been closed")

	// ErrSegmentExists is returned when attempting to create a segment that
	// already exists.
	ErrSegmentExists = errors.New("segment already exists")

	// ErrSegmentReplaced is returned when attempting to read from a segment
	// that has been replaced due to log compaction. When this error is
	// encountered, operations should be retried in order to run against the
	// new segment.
	ErrSegmentReplaced = errors.New("segment was replaced")

	// ErrCommitLogDeleted is returned when attempting to read from a commit
	// log that has been deleted.
	ErrCommitLogDeleted = errors.New("commit log was deleted")

	// ErrCommitLogClosed is returned when attempting to read from a commit
	// log that has been closed.
	ErrCommitLogClosed = errors.New("commit log was closed")

	// timestamp returns the current time in Unix nanoseconds. This function
	// exists for mocking purposes.
	timestamp = func() int64 { return time.Now().UnixNano() }
)

type segment struct {
	writer         io.Writer
	reader         io.Reader
	log            *os.File
	Index          *index
	BaseOffset     int64
	firstOffset    int64
	lastOffset     int64
	firstWriteTime int64
	lastWriteTime  int64
	position       int64
	maxBytes       int64
	path           string
	suffix         string
	waiters        map[interface{}]chan struct{}
	sealed         bool
	closed         bool
	replaced       bool
	deleted        bool // marked for deletion, excluded from read path

	sync.RWMutex
}

func newSegment(path string, baseOffset, maxBytes int64, isNew bool, suffix string) (*segment, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// If this is a new segment, ensure the file doesn't already exist.

// setupIndex creates and initializes an index.
// Initialization is:
// - Initialize index position
// - Initialize firstOffset/lastOffset
// - Initialize firstWriteTime/lastWriteTime
// If the index is corrupt, it will attempt to rebuild it from the log file.
func (s *segment) setupIndex() (err error) { _ = "STUB: not implemented"; return nil }

// Index is corrupt, attempt to rebuild from log file

// Re-initialize after rebuild

// If lastEntry is nil, the index is empty.

// Read the first entry to get firstOffset and firstWriteTime.

// rebuildIndex rebuilds the index by scanning the log file.
// This is called when a corrupt index is detected.
func (s *segment) rebuildIndex() error {
	_ = "STUB: not implemented"
	// Close and remove the corrupt index
	return nil
}

// Ignore close errors on corrupt index

// Create a fresh index

// Reset index position to 0 so we write from the beginning.
// newIndex() sets position = file size (10MB pre-allocated), but we need
// to write from the start. We can't call InitializePosition() here because
// it would fail on ReadAt due to position bounds checking. After we rebuild
// the entries, setupIndex will call InitializePosition() to finalize.

// If log file is empty, we're done

// Scan the log file and rebuild index entries

// Read message set header

// Partial header, stop here

// Validate the entry looks reasonable
// Max 100MB message
// Invalid size, stop here

// Check we have enough data for the full message

// Incomplete message, stop here

// Create index entry

// After rebuilding, set position to file size so InitializePosition() can
// read all entries during its binary search. The entries we wrote are
// non-zero, and the rest of the pre-allocated file is zeros (empty entries).

// CheckSplit determines if a new log segment should be rolled out either
// because this segment is full or LogRollTime has passed since the first
// message was written to the segment.
func (s *segment) CheckSplit(logRollTime time.Duration) bool {
	_ = "STUB: not implemented"
	return false
}

// Don't roll a new segment if there have been no writes to the segment
// or LogRollTime is disabled.

// Check if LogRollTime has passed since first write.

// Seal a segment from being written to. This is called on the former active
// segment after a new segment is rolled or when the segment is closed. This is
// a no-op if the segment is already sealed.
func (s *segment) Seal() { _ = "STUB: not implemented"; return }

func (s *segment) seal() { _ = "STUB: not implemented"; return }

// Notify any readers waiting for data.

// nolint: errcheck

func (s *segment) NextOffset() int64 { _ = "STUB: not implemented"; return 0 }

// If the segment hasn't been written to, the next offset should be the
// base offset.

func (s *segment) FirstOffset() int64 { _ = "STUB: not implemented"; return 0 }

func (s *segment) FirstWriteTime() int64 { _ = "STUB: not implemented"; return 0 }

func (s *segment) LastOffset() int64 { _ = "STUB: not implemented"; return 0 }

func (s *segment) Position() int64 { _ = "STUB: not implemented"; return 0 }

func (s *segment) IsEmpty() bool { _ = "STUB: not implemented"; return false }

func (s *segment) MessageCount() int64 { _ = "STUB: not implemented"; return 0 }

func (s *segment) WriteMessageSet(ms []byte, entries []*entry) error {
	_ = "STUB: not implemented"
	return nil
}

// write a byte slice to the log at the current position. This increments the
// offset as well as sets the position to the new tail.
func (s *segment) write(p []byte, entries []*entry) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (s *segment) ReadAt(p []byte, off int64) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (s *segment) notifyWaiters() { _ = "STUB: not implemented"; return }

func (s *segment) WaitForLEO(waiter interface{}, expectedLEO, actualLEO int64) <-chan struct{} {
	_ = "STUB: not implemented"
	return nil

	// Check expected LEO against last known LEO and against the current
	// (active) segment's last offset in case the LEO changed since we last
	// checked it. If the current segment's last offset is -1, this means the
	// segment is empty and we should wait for data.
}

// LEO has since changed so close channel immediately.

func (s *segment) WaitForData(waiter interface{}, pos int64) <-chan struct{} {
	_ = "STUB: not implemented"
	return nil
}

func (s *segment) waitForData(waiter interface{}, pos int64) <-chan struct{} {
	_ = "STUB: not implemented"
	// Check if we're already registered.
	return nil
}

// Check if data has been written and/or the segment was filled.

func (s *segment) removeWaiter(waiter interface{}) { _ = "STUB: not implemented"; return }

// Close a segment such that it can no longer be read from or written to. This
// operation is idempotent.
func (s *segment) Close() error { _ = "STUB: not implemented"; return nil }

func (s *segment) close() error { _ = "STUB: not implemented"; return nil }

// Cleaned creates a cleaned segment for this segment.
func (s *segment) Cleaned() (*segment, error) { _ = "STUB: not implemented"; return nil, nil }

// Truncated creates a truncated segment for this segment.
func (s *segment) Truncated() (*segment, error) { _ = "STUB: not implemented"; return nil, nil }

// Replace replaces the given segment with the callee.
func (s *segment) Replace(old *segment) error { _ = "STUB: not implemented"; return nil }

// findEntry returns the first entry whose offset is greater than or equal to
// the given offset.
func (s *segment) findEntry(offset int64) (*entry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// findEntryByTimestamp returns the first entry whose timestamp is greater than
// or equal to the given timestamp.
func (s *segment) findEntryByTimestamp(timestamp int64) (*entry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Delete closes the segment and then deletes its log and index files.
func (s *segment) Delete() error { _ = "STUB: not implemented"; return nil }

// MarkDeleted marks the segment as deleted, removing it from the read path.
// This should be called before actually deleting files to ensure readers
// don't see the segment while deletion is in progress.
func (s *segment) MarkDeleted() { _ = "STUB: not implemented"; return }

// IsDeleted returns true if the segment has been marked for deletion.
func (s *segment) IsDeleted() bool { _ = "STUB: not implemented"; return false }

type segmentScanner struct {
	s  *segment
	is *indexScanner
}

func newSegmentScanner(segment *segment) *segmentScanner { _ = "STUB: not implemented"; return nil }

// Scan should be called repeatedly to iterate over the messages in the
// segment, it will return io.EOF when there are no more messages.
func (s *segmentScanner) Scan() (messageSet, *entry, error) {
	_ = "STUB: not implemented"
	return *new(messageSet), nil, nil
}

// reverseSegmentScanner is used to iterate over messages in a segment in
// reverse order (newest to oldest).
type reverseSegmentScanner struct {
	s   *segment
	ris *reverseIndexScanner
}

// newReverseSegmentScanner creates a scanner that iterates from the given
// offset backwards.
func newReverseSegmentScanner(segment *segment, startOffset int64) *reverseSegmentScanner {
	_ = "STUB: not implemented"
	// Convert log offset to index entry offset
	return nil
}

// newReverseSegmentScannerFromEnd creates a scanner that starts at the last
// message in the segment and iterates backwards.
func newReverseSegmentScannerFromEnd(segment *segment) *reverseSegmentScanner {
	_ = "STUB: not implemented"
	return nil
}

// Scan reads the current message and moves to the previous one.
// Returns io.EOF when there are no more messages.
func (s *reverseSegmentScanner) Scan() (messageSet, *entry, error) {
	_ = "STUB: not implemented"
	return *new(messageSet), nil, nil
}

func (s *segment) logPath() string { _ = "STUB: not implemented"; return "" }

func (s *segment) indexPath() string { _ = "STUB: not implemented"; return "" }
