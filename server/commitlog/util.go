package commitlog

// findSegment returns the first segment whose next assignable offset is
// greater than the given offset. Returns nil and the index where the segment
// would be if there is no such segment.
func findSegment(segments []*segment, offset int64) (*segment, int) {
	_ = "STUB: not implemented"
	return nil, 0
}

// findSegmentContains returns the first segment whose next assignable offset
// is greater than the given offset and a bool indicating if the returned
// segment contains the offset, meaning the offset is between the segment's
// base offset and next assignable offset. Note that because the segment could
// be compacted, "contains" does not guarantee the offset is actually present,
// only that it's within the bounds.
func findSegmentContains(segments []*segment, offset int64) (*segment, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// findSegmentIndexByTimestamp returns the index of the first segment whose
// base timestamp is greater than the given timestamp. Returns the index where
// the segment would be if there is no segment whose base timestamp is greater,
// i.e. the length of the slice.
func findSegmentIndexByTimestamp(segments []*segment, timestamp int64) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Read the first entry in the segment to determine the base timestamp.

// findSegmentByBaseOffset returns the first segment whose base offset is
// greater than or equal to the given offset. Returns nil if there is no such
// segment.
func findSegmentByBaseOffset(segments []*segment, offset int64) *segment {
	_ = "STUB: not implemented"
	return nil
}

func roundDown(total, factor int64) int64 { _ = "STUB: not implemented"; return 0 }

func exists(path string) bool { _ = "STUB: not implemented"; return false }
