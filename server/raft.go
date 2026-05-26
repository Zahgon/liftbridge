package server

import (
	"context"
	"io"
	"sync"
	"time"

	"github.com/hashicorp/raft"
	raftboltdb "github.com/hashicorp/raft-boltdb/v2"
	"github.com/nats-io/nats.go"

	proto "github.com/liftbridge-io/liftbridge/server/protocol"
)

const (
	defaultJoinRaftGroupTimeout       = time.Second
	defaultRaftJoinAttempts           = 30
	defaultBootstrapMisconfigInterval = 10 * time.Second
	defaultRaftApplyTimeout           = 5 * time.Second
)

var (
	raftJoinAttempts           = defaultRaftJoinAttempts
	bootstrapMisconfigInterval = defaultBootstrapMisconfigInterval
)

// timeoutFuture wraps a raft.Future with a timeout on the call to Error().
// This is used as a workaround to an issue in the Hashicorp Raft library in
// which Raft operations can deadlock if a quorum cannot be reached (which can
// be indefinitely, e.g. in the case of shutting down the cluster). See this
// issue: https://github.com/hashicorp/raft/issues/498.
type timeoutFuture struct {
	deadline  time.Time
	wrapped   raft.Future
	mu        sync.RWMutex
	responded bool
	err       error
}

// newTimeoutFuture returns a raft.ApplyFuture wrapping the provided
// raft.Future whose call to Error() will timeout at the given deadline.
func newTimeoutFuture(deadline time.Time, wrapped raft.Future) raft.ApplyFuture {
	_ = "STUB: not implemented"
	return *new(raft.ApplyFuture)
}

func (t *timeoutFuture) Error() error { _ = "STUB: not implemented"; return nil }

func (t *timeoutFuture) Response() interface{} { _ = "STUB: not implemented"; return nil }

func (t *timeoutFuture) Index() uint64 { _ = "STUB: not implemented"; return 0 }

// raftNode is a handle to a member in a Raft consensus group.
type raftNode struct {
	leader int64
	sync.Mutex
	closed bool
	*raft.Raft
	store     *raftboltdb.BoltStore
	transport *raft.NetworkTransport
	logInput  io.WriteCloser
	joinSub   *nats.Subscription
	notifyCh  <-chan bool
}

// isLeader indicates if the Raft node is currently the leader.
func (r *raftNode) isLeader() bool { _ = "STUB: not implemented"; return false }

// setLeader sets the Raft node as the current leader or as a follower.
func (r *raftNode) setLeader(leader bool) { _ = "STUB: not implemented"; return }

// applyOperation proposes the given operation to the Raft cluster. This should
// only be called when the server is metadata leader. However, if the server
// has lost leadership, the returned future will yield an error. This will use
// the deadline provided on the context and check for preconditions using the
// supplied function, if provided. This will only return an error if
// preconditions have failed, indicating the operation was not proposed to the
// Raft cluster.
func (r *raftNode) applyOperation(ctx context.Context, op *proto.RaftLog,
	checkPreconditions func(*proto.RaftLog) error) (raft.ApplyFuture, error) {
	_ = "STUB: not implemented"
	return *new(raft.ApplyFuture), nil
}

// We will acquire the mutex to prevent Raft operations from interleaving
// such that preconditions can be validated.

// Ensure the FSM is up to date by issuing a barrier.

// Check that the FSM preconditions are valid before performing the
// Raft operation.

// Apply the Raft Operation.

// getCommitIndex returns the latest committed Raft index.
func (r *raftNode) getCommitIndex() uint64 { _ = "STUB: not implemented"; return 0 }

// shutdown attempts to stop the Raft node.
func (r *raftNode) shutdown() error { _ = "STUB: not implemented"; return nil }

// raftLogger implements io.WriteCloser by piping data to the Server logger.
type raftLogger struct {
	*Server
}

// Write pipes the given data to the Server logger.
func (r *raftLogger) Write(b []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// [DEBUG]

// [INFO]

// [WARN]

// [ERR]

// Close is a no-op to implement io.WriteCloser.
func (r *raftLogger) Close() error {
	_ = "STUB: not implemented"

	// setupMetadataRaft creates and starts an embedded Raft node for replicating
	// cluster metadata. The node will load configuration from previous state, if
	// there is any. If there isn't previous state, depending on server
	// configuration, this will attempt to join an existing cluster, bootstrap as a
	// seed node, or bootstrap using a predefined cluster configuration. If joining
	// an existing cluster, this will attempt to join for up to 30 seconds before
	// giving up and returning an error.
	return nil
}

func (s *Server) setupMetadataRaft() (*raftNode, error) { _ = "STUB: not implemented"; return nil, nil }

// Bootstrap if there is no previous state and we are starting this node as
// a seed or a cluster configuration is provided.

// Attempt to join the cluster if we're not bootstrapping.

// NATS transport uses ID for addr.

// Attempt to join for up to 30 seconds before giving up.

// If node is started with bootstrap, regardless if state exists or
// not, try to detect (and report) other nodes in same cluster started
// with bootstrap=true.

// bootstrapCluster bootstraps the node for the provided Raft group either as a
// seed node or with the given peer configuration, depending on configuration
// and with the latter taking precedence.
func (s *Server) bootstrapCluster(node *raft.Raft) error {
	_ = "STUB: not implemented"
	// Include ourself in the cluster.
	return nil
}

// Bootstrap using provided cluster configuration.

// Don't add ourselves twice.

// NATS transport uses ID as addr.

// Bootstrap as a seed node.

// Enforce quorum size limit. Any servers beyond the limit are non-voters.
// However, the local server must always be a voter to bootstrap successfully
// (required by raft v1.7.3+).

// Sort servers but ensure local server is placed in the voter range.
// First, sort all servers by ID.

// Find where local server is in the sorted list.

// If local server would be a non-voter (beyond maxQuorum), swap it with the
// last voter to ensure it can bootstrap.

// Mark servers beyond maxQuorum as non-voters.

// detectBootstrapMisconfig attempts to detect if any other servers were
// started in bootstrap seed mode. If any are detected, the server will panic
// since this is a fatal state.
func (s *Server) detectBootstrapMisconfig() { _ = "STUB: not implemented"; return }

// Ignore message to ourself

// createRaftNode creates and starts an embedded Raft node for replicating
// cluster metadata. It returns a bool indicating if the Raft node had existing
// state that was loaded.
func (s *Server) createRaftNode() (*raftNode, bool, error) {
	_ = "STUB: not implemented"
	return nil, false, nil
}

// Configure Raft.

// Setup a channel for reliable leader notifications.

// Setup Raft communication.

// Create the snapshot store. This allows Raft to truncate the log.

// Create the log store and cache.

// Instantiate the Raft node.

// Check if there is existing state.

// Handle requests to join the cluster.

// newClusterJoinRequestHandler creates a NATS handler for handling requests
// to join the Raft cluster.
func (s *Server) newClusterJoinRequestHandler(node *raft.Raft) func(*nats.Msg) {
	_ = "STUB: not implemented"
	return nil
}

// Drop the request if we're not the leader. There's no race condition
// after this check because even if we proceed with the cluster add, it
// will fail if the node is not the leader as cluster changes go
// through the Raft log.

// No-op if the request came from ourselves.

// Add the node to the cluster with appropriate suffrage. This is
// idempotent.

// Send the response.

// addAsVoter returns a bool indicating if a new node to be added to the
// cluster should be added as a voter or not based on current configuration. If
// we are below the max quorum size or there is no quorum limit, the new node
// will be added as a voter.  Otherwise, it's added as a non-voter.
func (s *Server) addAsVoter(node *raft.Raft) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// If there is a quorum limit, count the number of voting members.

// baseMetadataRaftSubject returns the base NATS subject used for Raft-related
// operations.
func (s *Server) baseMetadataRaftSubject() string { _ = "STUB: not implemented"; return "" }

func computeDeadline(ctx context.Context) time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}
