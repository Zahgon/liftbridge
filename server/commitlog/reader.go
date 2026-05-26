package commitlog

import (
	"context"
	"errors"
	"sync"
)

// ErrCommitLogReadonly is returned when the end of a readonly CommitLog has
// been reached.
var ErrCommitLogReadonly = errors.New("end of readonly log")

// MessageReader is the interface implemented by both Reader and ReverseReader.
// It allows reading messages from a CommitLog either forwards or backwards.
type MessageReader interface {
	ReadMessage(ctx context.Context, headersBuf []byte) (SerializedMessage, int64, int64, uint64, error)
}

type contextReader interface {
	Read(context.Context, []byte) (int, error)
}

// Reader reads messages atomically from a CommitLog. Readers should not be
// used concurrently.
type Reader struct {
	ctxReader   contextReader
	offset      int64
	log         *commitLog
	uncommitted bool
}

// NewReader creates a new Reader starting at the given offset. If uncommitted
// is true, the Reader will read uncommitted messages from the log. Otherwise,
// it will only return committed messages.
func (l *commitLog) NewReader(offset int64, uncommitted bool) (*Reader, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ReadMessage reads a single message from the underlying CommitLog or blocks
// until one is available. It returns the SerializedMessage in addition to its
// offset, timestamp, and leader epoch. This may return uncommitted messages if
// the reader was created with the uncommitted flag set to true.
//
// ReadMessage should not be called concurrently, and the headersBuf slice
// should have a capacity of at least 28.
//
// TODO: Should this just return a MessageSet directly instead of a Message and
// the MessageSet header values?
func (r *Reader) ReadMessage(ctx context.Context, headersBuf []byte) (SerializedMessage, int64, int64, uint64, error) {
	_ = "STUB: not implemented"
	return *new(SerializedMessage), 0, 0, 0, nil
}

// The log was deleted while we were trying to read.

// The log was closed while we were trying to read.

// The log was set to readonly while we were trying to read.

// ErrSegmentReplaced indicates we attempted to read from a log
// segment that was replaced due to compaction, so reinitialize the
// contextReader and try again to read from the new segment.

type uncommittedReader struct {
	cl  *commitLog
	seg *segment
	mu  sync.Mutex
	pos int64
}

func (r *uncommittedReader) Read(ctx context.Context, p []byte) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// We hit the end of the segment.

// Check if there are more segments.

// Otherwise, wait for segment to be written to (or split).

// At this point, either the segment has more data or, if it was
// full, a new segment was rolled. Try to read from the segment
// again.

// We hit an EOF after waiting for data which means a new segment was
// rolled, so move to the next segment.

// If there are not enough segments to read, wait for new segment to be
// appended or the context to be canceled.

func (r *uncommittedReader) waitForData(ctx context.Context, seg *segment) bool {
	_ = "STUB: not implemented"
	return false
}

// newReaderUncommitted returns a contextReader which reads data from the log
// starting at the given offset.
func (l *commitLog) newReaderUncommitted(offset int64) (contextReader, error) {
	_ = "STUB: not implemented"
	return *new(contextReader), nil
}

type committedReader struct {
	cl    *commitLog
	seg   *segment
	hwSeg *segment
	mu    sync.Mutex
	pos   int64
	hwPos int64
	hw    int64
}

func (r *committedReader) Read(ctx context.Context, p []byte) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// If seg is nil then the reader offset exceeded the HW, i.e. the log is
// either empty or the offset overflows the HW. This means we need to wait
// for data.

// We want to read the next committed message.

// The HW has not changed, so wait for it to update.

// Sync the HW.

func (r *committedReader) readLoop(
	ctx context.Context, p []byte, segments []*segment) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// If we're reading from the HW segment, read up to the HW pos.

// We hit the end of the segment, so jump to the next one.

// QUESTION: Should this ever happen?

// We hit the HW, so sync the latest.

// The HW has not changed, so wait for it to update.

// Sync the HW.

func (r *committedReader) waitForHW(ctx context.Context, hw int64) error {
	_ = "STUB: not implemented"
	return nil
}

// newReaderCommitted returns a contextReader which reads only committed data
// from the log starting at the given offset.
func (l *commitLog) newReaderCommitted(offset int64) (contextReader, error) {
	_ = "STUB: not implemented"
	return *new(contextReader), nil
}

// If offset exceeds HW, wait for the next message. This also covers the
// case when the log is empty.

func getHWPos(segments []*segment, hw int64) (int, int64, error) {
	_ = "STUB: not implemented"
	return 0, 0, nil
}

func min(x, y int64) int64 { _ = "STUB: not implemented"; return 0 }

// ReverseReader reads messages in reverse order (newest to oldest) from a
// CommitLog. ReverseReaders should not be used concurrently.
type ReverseReader struct {
	log         *commitLog
	segments    []*segment
	segIdx      int // Current segment index (starts at last segment)
	scanner     *reverseSegmentScanner
	stopOffset  int64 // Stop reading at this offset (inclusive)
	uncommitted bool
}

// NewReverseReader creates a new ReverseReader starting at the given offset
// and reading backwards. If uncommitted is true, the Reader will read
// uncommitted messages from the log. Otherwise, it will only return committed
// messages (starting from HW).
func (l *commitLog) NewReverseReader(startOffset int64, uncommitted bool) (*ReverseReader, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// For committed reads, start from HW if startOffset exceeds it

// Log is empty

// Find the segment containing the start offset

// Read all the way to the beginning by default

// NewReverseReaderFromEnd creates a new ReverseReader starting at the end of
// the log (either LEO for uncommitted or HW for committed).
func (l *commitLog) NewReverseReaderFromEnd(uncommitted bool) (*ReverseReader, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SetStopOffset sets the offset at which to stop reading (inclusive).
// Messages with offsets less than stopOffset will not be returned.
func (r *ReverseReader) SetStopOffset(offset int64) { _ = "STUB: not implemented"; return }

// ReadMessage reads the next message in reverse order (from newest to oldest).
// Returns io.EOF when there are no more messages or the stop offset is reached.
func (r *ReverseReader) ReadMessage(ctx context.Context, headersBuf []byte) (
	SerializedMessage, int64, int64, uint64, error) {
	_ = "STUB: not implemented"
	return *new(SerializedMessage), 0, 0, 0, nil
}

// Try to read from current segment

// Move to previous segment

// No more segments

// Check stop offset

// Extract message from message set
