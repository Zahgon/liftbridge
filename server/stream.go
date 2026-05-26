package server

import (
	"sync"
	"time"

	proto "github.com/liftbridge-io/liftbridge/server/protocol"
)

// stream is a message stream consisting of one or more partitions. Each
// partition maps to a NATS subject and is the unit of replication.
type stream struct {
	name         string
	subject      string
	config       *proto.StreamConfig
	partitions   map[int32]*partition
	resumeAll    bool // When partition(s) are paused, this indicates if all should be resumed
	tombstone    bool // Indicates if the stream is marked for deletion during Raft recovery
	creationTime time.Time
	mu           sync.RWMutex
}

// newStream creates a stream for the given NATS subject. All stream
// interactions should only go through the exported functions.
func newStream(name, subject string, config *proto.StreamConfig, creationTime time.Time,
	srvConfig *Config) *stream {
	_ = "STUB: not implemented"
	return nil
}

// applyReservedStreamOverrides lets us apply special handling to internal
// streams. This is used to allow changing configuration of internal streams on
// server restart without having to apply the changes through Raft.
// TODO: This is a bit of a hack and should probably be replaced if/when a
// stream patch Raft operation is supported.
func applyReservedStreamOverrides(s *stream, config *Config) { _ = "STUB: not implemented"; return }

// String returns a human-readable representation of the stream.
func (s *stream) String() string { _ = "STUB: not implemented"; return "" }

// GetName returns the stream's globally unique name.
func (s *stream) GetName() string { _ = "STUB: not implemented"; return "" }

// GetSubject returns the stream's NATS subject.
func (s *stream) GetSubject() string { _ = "STUB: not implemented"; return "" }

// GetConfig returns the stream's custom configuration.
func (s *stream) GetConfig() *proto.StreamConfig { _ = "STUB: not implemented"; return nil }

// GetResumeAll returns a bool indicating if the stream was paused with
// ResumeAll enabled. This means a message published to any of the stream's
// partitions will resume any paused partitions.
func (s *stream) GetResumeAll() bool { _ = "STUB: not implemented"; return false }

// SetResumeAll sets the bool used to indicate if the stream was paused with
// ResumeAll enabled. This means a message published to any of the stream's
// partitions will resume any paused partitions.
func (s *stream) SetResumeAll(resumeAll bool) { _ = "STUB: not implemented"; return }

// GetPartitions returns a map of partition ID to partition.
func (s *stream) GetPartitions() map[int32]*partition { _ = "STUB: not implemented"; return nil }

// GetPartition returns the partition with the given ID or nil if there is no
// such partition.
func (s *stream) GetPartition(id int32) *partition { _ = "STUB: not implemented"; return nil }

// SetPartition sets the partition with the given id on the stream.
func (s *stream) SetPartition(id int32, p *partition) { _ = "STUB: not implemented"; return }

// GetCreationTime returns the steam's creation time.
func (s *stream) GetCreationTime() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

// Close the stream by closing each of its partitions.
func (s *stream) Close() error { _ = "STUB: not implemented"; return nil }

// Pause some or all the partitions of this stream. Returns a list of the
// partitions that were paused.
func (s *stream) Pause(partitions []int32, resumeAll bool) ([]*partition, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Delete the stream by closing and deleting each of its partitions.
func (s *stream) Delete() error { _ = "STUB: not implemented"; return nil }

// SetReadonly sets the readonly flag on some or all the partitions of this
// stream.
func (s *stream) SetReadonly(partitions []int32, readonly bool) error {
	_ = "STUB: not implemented"
	return nil
}

// Tombstone sets the tombstone marker on the stream which determines if the
// stream is marked for deletion during the Raft recovery process.
func (s *stream) Tombstone() { _ = "STUB: not implemented"; return }

// IsTombstoned indicates if the stream is marked for deletion during the Raft
// recovery process.
func (s *stream) IsTombstoned() bool { _ = "STUB: not implemented"; return false }
