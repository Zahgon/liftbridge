package server

import (
	"bytes"
	"context"
	"sync"
	"time"

	"github.com/nats-io/nats.go"

	"github.com/liftbridge-io/liftbridge/server/commitlog"
	proto "github.com/liftbridge-io/liftbridge/server/protocol"
)

// replicationOverhead is the non-data size overhead of replication messages: 8
// bytes for the leader epoch and 8 bytes for the HW.
const replicationOverhead = 16

// replicationRequest wraps a ReplicationRequest protobuf and a NATS subject
// where responses should be sent.
type replicationRequest struct {
	*proto.ReplicationRequest
	request  *nats.Msg
	received time.Time
}

// replicator handles replication requests from a particular replica and tracks
// its health. Requests are received on the requests channel and a long-running
// loop processes them and sends responses. If the replica does not catch up to
// the leader's log in maxLagTime, it's removed from the ISR until it catches
// back up.
type replicator struct {
	partition    *partition
	replica      string
	maxLagTime   time.Duration
	lastCaughtUp time.Time
	lastSeen     time.Time
	requests     chan replicationRequest
	mu           sync.RWMutex
	leader       string
	epoch        uint64
	headersBuf   [28]byte // scratch buffer for reading message headers
	writer       replicationProtocolWriter
	waiter       <-chan struct{}
}

func newReplicator(epoch uint64, replica string, p *partition) *replicator {
	_ = "STUB: not implemented"
	return nil
}

// start a long-running replication loop for the given leader epoch until the
// stop channel is closed. This loop will receive messages from the requests
// channel, update the replica last-seen timestamp and latest offset, and send
// a batch of messages starting at the requested offset, if there are any
// available. The response will also include the leader epoch and HW. If the
// replica doesn't send a request or catch up to the leader's log in
// maxLagTime, it will be removed from the ISR until it catches back up.
func (r *replicator) start(stop <-chan struct{}) { _ = "STUB: not implemented"; return }

// Start a goroutine to track the replica's health.

// Update the ISR replica's latest offset for the partition. This is
// used by the leader to know when to commit messages.

// Check if we're caught up.

// Create a log reader starting at the requested offset. Wrap this in
// an anonymous function to avoid leaking the deferred context cancel.

// Send a response to short-circuit request timeout.

// Send a batch of messages to the replica.

// Send a response to short-circuit request timeout.

func (r *replicator) request(req replicationRequest) { _ = "STUB: not implemented"; return }

// tick is a long-running call that checks to see if the follower hasn't sent
// any replication requests or hasn't consumed up to the leader's log end
// offset for the lag-time duration. If this is the case, the follower is
// removed from the ISR until it catches back up.
func (r *replicator) tick(stop <-chan struct{}) { _ = "STUB: not implemented"; return }

// Follower has not sent a request or has not caught up in
// maxLagTime, so remove it from the ISR.

// Add replica back into ISR.

// shrinkISR sends a ShrinkISR request to the controller to remove the replica
// from the ISR.
func (r *replicator) shrinkISR() { _ = "STUB: not implemented"; return }

// expandISR sends an ExpandISR request to the controller to add the replica to
// the ISR.
func (r *replicator) expandISR() { _ = "STUB: not implemented"; return }

// replicate sends a batch of messages to the given NATS inbox along with the
// leader epoch and HW.
func (r *replicator) replicate(
	ctx context.Context, reader *commitlog.Reader, request *nats.Msg, offset int64) error {
	_ = "STUB: not implemented"
	return nil
}

// Check if this message will put us over the batch size limit. If it
// does, flush the batch now.

// Write the message to the buffer.

// Flush the batch.

// caughtUp is called when the follower has caught up with the leader's log.
// This will register a data waiter on the log so that the leader can notify
// the follower when new data is available to replicate.
func (r *replicator) caughtUp(stop <-chan struct{}, leo int64, req replicationRequest) {
	_ = "STUB: not implemented"
	return
}

// Register a waiter to be notified when new messages are written after
// the current log end offset to preempt an idle follower.

// sendHW sends the leader epoch and HW to the given NATS inbox.
func (r *replicator) sendHW(request *nats.Msg) error { _ = "STUB: not implemented"; return nil }

type replicationProtocolWriter interface {
	Write(offset int64, headers, message []byte) error
	Flush(func(data []byte) error) error
	Len() int
	Reset()
}

type protocolWriter struct {
	*replicator
	buf        *bytes.Buffer
	log        commitlog.CommitLog
	lastOffset int64
	dataPos    int
	stop       <-chan struct{}
}

func newReplicationProtocolWriter(r *replicator, stop <-chan struct{}) replicationProtocolWriter {
	_ = "STUB: not implemented"
	return *new(replicationProtocolWriter)
}

func (w *protocolWriter) Write(offset int64, headers, message []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (w *protocolWriter) Flush(write func([]byte) error) error {
	_ = "STUB: not implemented"
	return nil

	// Replace the HW.
}

func (w *protocolWriter) Len() int { _ = "STUB: not implemented"; return 0 }

func (w *protocolWriter) Reset() { _ = "STUB: not implemented"; return }

// Write envelope header.

// Write the leader epoch.

// Reserve space for the HW. This will be replaced with the HW at the time
// of flush.
