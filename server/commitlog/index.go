package commitlog

import (
	"os"
	"sync"

	"github.com/pkg/errors"
	"github.com/tysonmote/gommap"
)

var errIndexCorrupt = errors.New("corrupt index file")

const (
	offsetWidth    = 4
	timestampWidth = 8
	positionWidth  = 4
	sizeWidth      = 4
	entryWidth     = offsetWidth + timestampWidth + positionWidth + sizeWidth
)

type index struct {
	options
	mmap     gommap.MMap
	file     *os.File
	size     int64
	mu       sync.RWMutex
	position int64
	closed   bool
}

type entry struct {
	Offset      int64
	Timestamp   int64
	LeaderEpoch uint64
	Position    int64
	Size        int32
}

// relEntry is an Entry relative to the base fileOffset
type relEntry struct {
	Offset    int32
	Timestamp int64
	Position  int32
	Size      int32
}

func newRelEntry(e *entry, baseOffset int64) relEntry {
	_ = "STUB: not implemented"
	return *new(relEntry)
}

func (rel relEntry) fill(e *entry, baseOffset int64) { _ = "STUB: not implemented"; return }

type options struct {
	path       string
	bytes      int64
	baseOffset int64
}

func newIndex(opts options) (idx *index, err error) { _ = "STUB: not implemented"; return nil, nil }

// Pre-allocate the index if we just created it.

// Get updated stats after resize.

// Position returns the current position in the index to write to next. This
// value also represents the total length of the index.
func (idx *index) Position() int64 { _ = "STUB: not implemented"; return 0 }

func (idx *index) CountEntries() int64 { _ = "STUB: not implemented"; return 0 }

func (idx *index) writeEntries(entries []*entry) (err error) { _ = "STUB: not implemented"; return nil }

// ReadEntryAtFileOffset is used to read an index entry at the given
// byte offset of the index file. ReadEntryAtLogOffset is generally
// more useful for higher level use.
func (idx *index) ReadEntryAtFileOffset(e *entry, fileOffset int64) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// ReadEntryAtLogOffset is used to read an index entry at the given
// log offset of the index file.
func (idx *index) ReadEntryAtLogOffset(e *entry, logOffset int64) error {
	_ = "STUB: not implemented"
	return nil
}

func (idx *index) ReadAt(p []byte, offset int64) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (idx *index) writeAt(p []byte, offset int64) error {
	_ = "STUB: not implemented"
	// Check if we need to expand the index file.
	return nil
}

// Expand the index file.

// Re-mmap the index.

// Unmap the old index.

func (idx *index) Sync() error { _ = "STUB: not implemented"; return nil }

func (idx *index) sync() error { _ = "STUB: not implemented"; return nil }

func (idx *index) Close() error { _ = "STUB: not implemented"; return nil }

// Shrink truncates the memory-mapped index file to the size of its contents.
func (idx *index) Shrink() error { _ = "STUB: not implemented"; return nil }

func (idx *index) shrink() error { _ = "STUB: not implemented"; return nil }

func (idx *index) Name() string { _ = "STUB: not implemented"; return "" }

func (idx *index) InitializePosition() (*entry, error) {
	_ = "STUB: not implemented"
	// Find the first empty entry.
	return nil, nil
}

// Initialize the position.

// Index is empty.

// Return the last entry in the index.

// Do some sanity checks.

type indexScanner struct {
	idx    *index
	entry  *entry
	offset int64
}

func newIndexScanner(idx *index) *indexScanner { _ = "STUB: not implemented"; return nil }

func (s *indexScanner) Scan() (*entry, error) { _ = "STUB: not implemented"; return nil, nil }

// reverseIndexScanner is used to iterate over entries in reverse order
// (newest to oldest).
type reverseIndexScanner struct {
	idx    *index
	entry  *entry
	offset int64 // Current entry index (starts at last entry)
}

// newReverseIndexScanner creates a scanner that iterates from the given
// starting offset backwards to the beginning of the index.
func newReverseIndexScanner(idx *index, startOffset int64) *reverseIndexScanner {
	_ = "STUB: not implemented"
	return nil
}

// newReverseIndexScannerFromEnd creates a scanner that starts at the last
// entry in the index and iterates backwards.
func newReverseIndexScannerFromEnd(idx *index) *reverseIndexScanner {
	_ = "STUB: not implemented"
	// Get the number of entries in the index
	return nil
}

// Will return EOF on first Scan()

// Scan reads the current entry and moves to the previous one.
// Returns io.EOF when there are no more entries.
func (s *reverseIndexScanner) Scan() (*entry, error) { _ = "STUB: not implemented"; return nil, nil }

// Move to previous entry
