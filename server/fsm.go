package server

import (
	"io"

	"github.com/hashicorp/raft"

	proto "github.com/liftbridge-io/liftbridge/server/protocol"
)

// recoverLatestCommittedFSMLog returns the last committed Raft FSM log entry.
// It returns nil if there are no entries in the Raft log.
func (s *Server) recoverLatestCommittedFSMLog(applyIndex uint64) (*raft.Log, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// No entries.

// We are committing the first FSM log.

// Apply applies a Raft log entry to the controller FSM. This is invoked by
// Raft once a log entry is committed. It returns a value which will be made
// available on the ApplyFuture returned by Raft.Apply if that method was
// called on the same Raft node as the FSM.
//
// Note that, on restart, this can be called for entries that have already been
// committed to Raft as part of the recovery process. As such, this should be
// an idempotent call.
func (s *Server) Apply(l *raft.Log) interface{} {
	_ = "STUB: not implemented"
	// If recoveryStarted is false, the server was just started. We are going
	// to recover the last committed Raft FSM log entry, if any, to determine
	// the recovery high watermark. Once we apply all entries up to that point,
	// we know we've completed the recovery process and subsequent entries are
	// newly committed operations.
	//
	// During the recovery process, any recovered streams will not be started
	// until recovery is finished to avoid starting streams in an intermediate
	// state. Streams that are deleted will only be marked for deletion to
	// avoid deleting potentially valid stream data, e.g. in the case of a
	// stream being deleted, recreated, and then published to. When recovery
	// completes, we'll call finishedRecovery() to start the recovered streams
	// and delete any tombstoned streams.
	return nil
}

// If this returns an error, something is very wrong.

// Check if this is a "recovered" Raft entry, meaning we are still applying
// logs up to and including the latest recovered log.

// We've applied all entries up to the latest recovered log, so
// recovery is finished. Call finishedRecovery() to start any
// recovered streams and consumer groups and delete tombstoned
// streams.

// Unmarshal the log data and apply the operation to the FSM.

// Don't panic if the server is shutting down, just return the
// error.

// Send the Raft log entry to listeners.

// apply the given RaftLog to the FSM. This returns a value, if any, which
// should be made available on the ApplyFuture returned by Raft.Apply if that
// method was called on the same Raft node as the FSM. An error is returned if
// the operation could not be applied. The index parameter is the index of the
// entry in the Raft log. The recovered parameter indicates if this entry is
// being applied during the recovery process.
func (s *Server) apply(log *proto.RaftLog, index uint64, recovered bool) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Make sure to set the leader epoch on the partitions.

// startedRecovery should be called when the FSM has started replaying any
// unapplied log entries.
func (s *Server) startedRecovery() { _ = "STUB: not implemented"; return }

// If LogRecovery is enabled, prefix recovery logs with "-->" so they
// are visually distinct.

// If LogRecovery is disabled, we need to suppress logs while replaying
// the Raft log. Do this by discarding the log output.

// finishedRecovery should be called when the FSM has finished replaying any
// unapplied log entries. This will start any stream partitions and consumer
// groups recovered during the replay and delete any tombstoned streams. It
// returns the number of streams which had partitions that were recovered and
// the number of consumer groups that were recovered.
func (s *Server) finishedRecovery(epoch uint64) (int, int, error) {
	_ = "STUB: not implemented"
	return 0,

		// If LogRecovery is enabled, clear the logging prefix.
		0, nil
}

// If LogRecovery is disabled, we need to restore the previous log
// output.

// fsmSnapshot is returned by an FSM in response to a Snapshot. It must be safe
// to invoke fsmSnapshot methods with concurrent calls to Apply.
type fsmSnapshot struct {
	*proto.MetadataSnapshot
}

// Persist should dump all necessary state to the WriteCloser sink and call
// sink.Close() when finished or call sink.Cancel() on error.
func (f *fsmSnapshot) Persist(sink raft.SnapshotSink) error { _ = "STUB: not implemented"; return nil }

// Encode data.

// Write size and data to sink.

// Close the sink.

// Release is invoked when we are finished with the snapshot.
func (f *fsmSnapshot) Release() {
	_ = "STUB: not implemented"

	// Snapshot is used to support log compaction. This call should return an
	// FSMSnapshot which can be used to save a point-in-time snapshot of the FSM.
	// Apply and Snapshot are not called in multiple threads, but Apply will be
	// called concurrently with Persist. This means the FSM should be implemented
	// in a fashion that allows for concurrent updates while a snapshot is
	// happening.
	return
}

func (s *Server) Snapshot() (raft.FSMSnapshot, error) {
	_ = "STUB: not implemented"
	return *new(raft.FSMSnapshot), nil
}

// Restore is used to restore an FSM from a snapshot. It is not called
// concurrently with any other command. The FSM must discard all previous
// state.
func (s *Server) Restore(snapshot io.ReadCloser) error { _ = "STUB: not implemented"; return nil }

// Read snapshot size.

// Read snapshot.

// Drop state and restore.

// Mark streams and groups as recovered so they don't start leader/follower
// loops until finishedRecovery() is called after log replay completes.
// This is critical because s.api is not yet initialized during Restore().

// applyCreateStream adds the given stream and its partitions to the metadata
// store. If the stream is being recovered, its partitions will not be started
// until after the recovery process completes. If it is not being recovered,
// the partitions will be started as a leader or follower if applicable. An
// error is returned if the stream or any of its partitions already exist.
func (s *Server) applyCreateStream(protoStream *proto.Stream, recovered bool, epoch uint64) error {
	_ = "STUB: not implemented"
	// QUESTION: If this broker is not a replica for the stream, can we just
	// store a "lightweight" representation of the stream (i.e. the protobuf)
	// for recovery purposes? There is no need to initialize a commit log for
	// it.
	return nil
}

// applyShrinkISR removes the given replica from the partition and updates the
// partition epoch. If the partition epoch is greater than or equal to the
// specified epoch, this does nothing.
func (s *Server) applyShrinkISR(stream, replica string, partitionID int32, epoch uint64) error {
	_ = "STUB: not implemented"
	return nil
}

// applyExpandISR adds the given replica to the partition and updates the
// partition epoch. If the partition epoch is greater than or equal to the
// specified epoch, this does nothing.
func (s *Server) applyExpandISR(stream, replica string, partitionID int32, epoch uint64) error {
	_ = "STUB: not implemented"
	return nil
}

// applyChangePartitionLeader sets the partition's leader to the given replica
// and updates the partition epoch. If the partition epoch is greater than or
// equal to the specified epoch, this does nothing.
func (s *Server) applyChangePartitionLeader(stream, leader string, partitionID int32, epoch uint64) error {
	_ = "STUB: not implemented"
	return nil
}

// applyDeleteStream deletes the given stream partition. If this operation is
// being applied during recovery, this will only mark the stream with a
// tombstone. Tombstoned streams will be deleted after the recovery process
// completes.
func (s *Server) applyDeleteStream(streamName string, recovered bool, epoch uint64) error {
	_ = "STUB: not implemented"
	return nil
}

// applyPauseStream pauses the given stream partitions.
func (s *Server) applyPauseStream(stream string, partitions []int32, resumeAll bool) error {
	_ = "STUB: not implemented"
	return nil
}

// applyResumeStream unpauses the given stream partitions in the metadata
// store. If the partitions are being recovered, they will not be started until
// after the recovery process completes. If they are not being recovered, the
// partitions will be started as a leader or follower if applicable.
func (s *Server) applyResumeStream(streamName string, partitionIDs []int32, recovered bool) error {
	_ = "STUB: not implemented"
	return nil
}

// applySetStreamReadonly changes the stream partitions readonly flag in the
// metadata store.
func (s *Server) applySetStreamReadonly(streamName string, partitions []int32, readonly bool) error {
	_ = "STUB: not implemented"
	return nil
}

// applyCreateConsumerGroup adds the given consumer group to the metadata
// store. An error is returned if the consumer group already exists. If the
// group is being recovered, the member liveness checks won't be started until
// after the recovery process completes.
func (s *Server) applyCreateConsumerGroup(protoGroup *proto.ConsumerGroup, recovered bool) error {
	_ = "STUB: not implemented"
	return nil
}

// applyJoinConsumerGroup adds the given consumer to the consumer group. An
// error is returned if the group does not exist, the consumer is already a
// member of the group, or any of the provided streams do not exist. If the
// group is being recovered, the consumer liveness check won't be started until
// after the recovery process completes.
func (s *Server) applyJoinConsumerGroup(groupID, consumerID string, streams []string, epoch uint64) error {
	_ = "STUB: not implemented"
	return nil
}

// applyLeaveConsumerGroup removes the given consumer from the consumer group.
// An error is returned if the group does not exist or the consumer is not a
// member of the group.
func (s *Server) applyLeaveConsumerGroup(groupID, consumerID string, epoch uint64) error {
	_ = "STUB: not implemented"
	return nil
}

// applyChangeConsumerGroupCoordinator sets the group's coordinator to the
// given broker and updates the group epoch. If the group epoch is greater than
// or equal to the specified epoch, this does nothing.
func (s *Server) applyChangeConsumerGroupCoordinator(groupID, coordinator string, epoch uint64) error {
	_ = "STUB: not implemented"
	return nil
}
