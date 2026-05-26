package server

import (
	"sync"
	"time"

	"github.com/hashicorp/raft"
	client "github.com/liftbridge-io/liftbridge-api/v2/go"
)

const maxActivityPublishBackoff = 10 * time.Second

// activityManager ensures that activity events get published to the activity
// stream. This ensures that events are published at least once and in the
// order in which they occur with respect to the Raft log.
type activityManager struct {
	*Server
	lastPublishedRaftIndex uint64
	commitCh               chan struct{}
	leadershipLostCh       chan struct{}
	mu                     sync.RWMutex
}

func newActivityManager(s *Server) *activityManager { _ = "STUB: not implemented"; return nil }

// SetLastPublishedRaftIndex sets the Raft index of the latest event published
// to the activity stream. This is used to determine where to begin publishing
// events from in the log in the case of failovers or restarts.
func (a *activityManager) SetLastPublishedRaftIndex(index uint64) {
	_ = "STUB: not implemented"
	return
}

// LastPublishedRaftIndex returns the Raft index of the latest event published
// to the activity stream. This is used to determine where to begin publishing
// events from in the log in the case of failovers or restarts.
func (a *activityManager) LastPublishedRaftIndex() uint64 { _ = "STUB: not implemented"; return 0 }

// SignalCommit indicates a new event was committed to the Raft log.
func (a *activityManager) SignalCommit() { _ = "STUB: not implemented"; return }

// BecomeLeader should be called when this node has been elected as the
// metadata leader. This will set up the activity stream if it's enabled. It
// will then reconcile the last published event with the Raft log and begin
// publishing any un-published events. This should be called on the same
// goroutine as BecomeFollower.
func (a *activityManager) BecomeLeader() error { _ = "STUB: not implemented"; return nil }

// BecomeFollower should be called when this node has lost metadata leadership.
// This should be called on the same goroutine as BecomeLeader.
func (a *activityManager) BecomeFollower() error { _ = "STUB: not implemented"; return nil }

// dispatch is a long-running goroutine that runs while the server is the
// metadata leader. It handles publishing events to the activity stream as they
// are committed to the Raft log. Events are always published in the order in
// which they were committed to the log.
func (a *activityManager) dispatch() { _ = "STUB: not implemented"; return }

// TODO: Should we instead pass the commit index from FSM apply?

// We are caught up with the Raft log, so wait for new commits.

// handleRaftLog unmarshals the Raft log into an operation and, if applicable,
// publishes an event to the activity stream.
func (a *activityManager) handleRaftLog(l *raft.Log) error { _ = "STUB: not implemented"; return nil }

// Members on create should always contain a single consumer.

// Treat this as a join since it bootstraps the group.

// createActivityStream creates the activity stream and connects a local client
// that will be subscribed to it.
func (a *activityManager) createActivityStream() error { _ = "STUB: not implemented"; return nil }

// publishActivityEvent publishes an event on the activity stream.
func (a *activityManager) publishActivityEvent(event *client.ActivityStreamEvent) error {
	_ = "STUB: not implemented"
	return nil
}

// Update last published index in Raft.

func computeActivityPublishBackoff(previousBackoff time.Duration) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}
