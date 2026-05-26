package server

import (
	"net"
	"sync"
	"sync/atomic"

	"github.com/casbin/casbin/v2"
	"github.com/hashicorp/raft"
	gnatsd "github.com/nats-io/nats-server/v2/server"
	"github.com/nats-io/nats.go"
	"google.golang.org/grpc"

	"github.com/liftbridge-io/liftbridge/server/logger"
	"github.com/liftbridge-io/liftbridge/server/telemetry"
)

const stateFile = "liftbridge"

const (
	streamsConnName     = "streams"
	raftConnName        = "raft"
	replicationConnName = "replication"
	acksConnName        = "acks"
	publishesConnName   = "publishes"
	activityStream      = "__activity"
	cursorsStream       = "__cursors"
)

// reservedStreams contains reserved internal stream names.
var reservedStreams = []string{activityStream, cursorsStream}

// RaftLog represents an entry into the Raft log.
type RaftLog struct {
	*raft.Log
}

// RaftLogListener is a listener for Raft logs.
type RaftLogListener interface {
	Receive(*RaftLog)
}

// authzEnforcer contains a casbin enforcer and a lock, which is used to reload permissions safely
type authzEnforcer struct {
	enforcer  *casbin.Enforcer
	authzLock sync.RWMutex
}

// Server is the main Liftbridge object. Create it by calling New or
// RunServerWithConfig.
type Server struct {
	config             *Config
	listener           net.Listener
	port               int
	embeddedNATS       *gnatsd.Server
	nc                 *nats.Conn
	ncRaft             *nats.Conn
	ncRepl             *nats.Conn
	ncAcks             *nats.Conn
	ncPublishes        *nats.Conn
	logger             logger.Logger
	grpcServer         *grpc.Server
	api                *apiServer
	metadata           *metadataAPI
	shutdownCh         chan struct{}
	raftInitialized    chan struct{}
	raft               atomic.Value
	leaderSub          *nats.Subscription
	recoveryStarted    bool
	latestRecoveredLog *raft.Log
	mu                 sync.RWMutex
	shutdown           bool
	running            bool
	goroutineWait      sync.WaitGroup
	activity           *activityManager
	cursors            *cursorManager
	raftLogListenersMu sync.RWMutex
	raftLogListeners   []RaftLogListener
	authzEnforcer      *authzEnforcer
	telemetry          *telemetry.Collector
}

// RunServerWithConfig creates and starts a new Server with the given
// configuration. It returns an error if the Server failed to start.
func RunServerWithConfig(config *Config) (*Server, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// New creates a new Server with the given configuration. Call Start to run the
// Server.
func New(config *Config) *Server {
	_ = "STUB: not implemented"
	// Default data path to /tmp/liftbridge/<namespace> if not set.
	return nil
}

// Start the Server. This is not a blocking call. It will return an error if
// the Server cannot start properly.
func (s *Server) Start() (err error) { _ = "STUB: not implemented"; return nil }

// Create the data directory if it doesn't exist.

// Recover and persist metadata state.

// Initialize telemetry collector.

// Start embedded NATS server if configured.

// Set a lower bound of one second for SegmentMaxAge to avoid frequent log
// rolls which will cause performance problems. This is mainly here because
// SegmentMaxAge defaults to RetentionMaxAge if it's not set explicitly,
// so users could otherwise unknowingly cause frequent log rolls.

// Start telemetry collector.

// Stop will attempt to gracefully shut the Server down by signaling the stop
// and waiting for all goroutines to return.
func (s *Server) Stop() error { _ = "STUB: not implemented"; return nil }

// Stop telemetry collector.

// Close the raftInitialized channel in case the Raft node was never
// initialized to prevent a deadlock.

// Wait for goroutines to stop.

// IsLeader indicates if the server is currently the metadata leader or not. If
// consistency is required for an operation, it should be threaded through the
// Raft cluster since that is the single source of truth. If a server thinks
// it's leader when it's not, the operation it proposes to the Raft cluster
// will fail.
func (s *Server) IsLeader() bool { _ = "STUB: not implemented"; return false }

// IsRunning indicates if the server is currently running or has been stopped.
func (s *Server) IsRunning() bool { _ = "STUB: not implemented"; return false }

// GetListenPort returns the port the server is listening to. Returns 0 if the
// server is not listening.
func (s *Server) GetListenPort() int { _ = "STUB: not implemented"; return 0 }

// AddRaftLogListener adds a Raft log listener.
func (s *Server) AddRaftLogListener(listener RaftLogListener) { _ = "STUB: not implemented"; return }

// getConnectionAddress returns the connection address that should be used by
// the server. It uses the port the server is currently listening to if the
// connection port is 0, so that an OS-assigned port can be used as a connection
// port.
func (s *Server) getConnectionAddress() HostPort { _ = "STUB: not implemented"; return *new(HostPort) }

// recoverAndPersistState recovers any existing server metadata state from disk
// to initialize the server then writes the metadata back to disk.
func (s *Server) recoverAndPersistState() error {
	_ = "STUB: not implemented"
	// Attempt to recover state.
	return nil
}

// Recovered previous state.

// Persist server state.

// startEmbeddedNATS starts a NATS server embedded in this process. It returns
// once the server is ready to accept connections.
func (s *Server) startEmbeddedNATS() error { _ = "STUB: not implemented"; return nil }

// Disable NATS signal handling to prevent race with Liftbridge's signal
// handler. Without this, both servers register handlers for SIGINT/SIGTERM
// and whichever runs first wins - if NATS wins, it calls os.Exit()
// immediately, preventing Liftbridge from performing graceful shutdown.
// See: https://github.com/liftbridge-io/liftbridge/issues/373

// createNATSConns creates various NATS connections used by the server,
// including connections for stream data, Raft, replication, acks, and
// publishes.
func (s *Server) createNATSConns() error {
	_ = "STUB: not implemented"
	// NATS connection used for stream data.
	return nil
}

// NATS connection used for Raft metadata replication.

// NATS connection used for stream replication.

// NATS connection used for sending acks.

// NATS connection used for publishing messages.

// closeNATSConns closes the various NATS connections used by the server,
// including connections for stream data, Raft, replication, acks, and
// publishes.
func (s *Server) closeNATSConns() { _ = "STUB: not implemented"; return }

// startAPIServer configures and starts the gRPC API server.
func (s *Server) startAPIServer() error { _ = "STUB: not implemented"; return nil }

// Setup TLS if key/cert is set.

// Configure Authentication

// Configure authorization

// createNATSConn creates a new NATS connection with the given name.
func (s *Server) createNATSConn(name string) (*nats.Conn, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Shorten the time we wait to reconnect. Don't make it too short because
// it may exhaust the number of available FDs.

// Try to reconnect indefinitely.

// Disable buffering in the NATS client to avoid possible duplicate
// deliveries.

// Set connection handlers.

// startRaftLeadershipLoop start a goroutine for automatically responding to
// Raft leadership changes.
func (s *Server) startRaftLeadershipLoop(node *raftNode) { _ = "STUB: not implemented"; return }

// Node shutdown, just return.

// Node lost leadership, continue loop.

// Step down as leader.

// setRaft sets the Raft node for the server. This should only be called once
// on server start.
func (s *Server) setRaft(r *raftNode) { _ = "STUB: not implemented"; return }

// getRaft returns the Raft node for the server.
func (s *Server) getRaft() *raftNode { _ = "STUB: not implemented"; return nil }

// leadershipAcquired should be called when this node is elected leader.
func (s *Server) leadershipAcquired(raft *raftNode) error { _ = "STUB: not implemented"; return nil }

// Use a barrier to ensure all preceding operations are applied to the FSM.

// Subscribe to leader NATS subject for propagated requests.

// leadershipLost should be called when this node loses leadership.
func (s *Server) leadershipLost(raft *raftNode) error { _ = "STUB: not implemented"; return nil }

// Unsubscribe from leader NATS subject for propagated requests.

func (s *Server) isShutdown() bool { _ = "STUB: not implemented"; return false }

// natsDisconnectedHandler fires when the given NATS connection has been
// disconnected. This may indicate a temporary disconnect, in which case the
// client will automatically attempt to reconnect.
func (s *Server) natsDisconnectedHandler(nc *nats.Conn) {
	_ = "STUB: not implemented"
	// If the server was shut down, do nothing since this is an expected
	// disconnect.
	return
}

// natsReconnectedHandler fires when the given NATS connection has successfully
// reconnected.
func (s *Server) natsReconnectedHandler(nc *nats.Conn) { _ = "STUB: not implemented"; return }

// natsClosedHandler fires when the given NATS connection has been closed, i.e.
// permanently disconnected. At this point, the client will not attempt to
// reconnect to NATS.
func (s *Server) natsClosedHandler(nc *nats.Conn) {
	_ = "STUB: not implemented"
	// If the server was shut down, do nothing since this is an expected close.
	return
}

// natsErrorHandler fires when there is an asynchronous error on the NATS
// connection.
func (s *Server) natsErrorHandler(nc *nats.Conn, sub *nats.Subscription, err error) {
	_ = "STUB: not implemented"
	return
}

// handleServerInfoRequest is a NATS handler used to process requests for
// server information used in the metadata API.
func (s *Server) handleServerInfoRequest(m *nats.Msg) { _ = "STUB: not implemented"; return }

// Ignore requests from ourself.

// handlePartitionStatusRequest is a NATS handler used to process requests
// querying the status of a partition. This is used as a readiness check to
// determine if a created partition has actually started.
func (s *Server) handlePartitionStatusRequest(m *nats.Msg) { _ = "STUB: not implemented"; return }

// handlePartitionNotification is a NATS handler used to process notifications
// from a leader that new data is available on a partition for the follower to
// replicate if the follower is idle.
//
// When a follower reaches the end of the log, it starts to sleep in between
// replication requests to avoid overloading the leader. However, this causes
// added commit latency when new messages are published to the log since the
// follower is idle. As a result, the leader will note when a follower is
// caught up and send a notification in order to wake an idle follower back up
// when new data is written to the log.
func (s *Server) handlePartitionNotification(m *nats.Msg) { _ = "STUB: not implemented"; return }

// Wake the follower up.

// getServerInfoInbox returns the NATS subject used for handling server
// information requests.
func (s *Server) getServerInfoInbox() string { _ = "STUB: not implemented"; return "" }

// getPartitionStatusInbox returns the NATS subject used for handling stream
// status requests.
func (s *Server) getPartitionStatusInbox(id string) string { _ = "STUB: not implemented"; return "" }

// getMetadataReplyInbox returns a random NATS subject to use for metadata
// responses scoped to the cluster namespace.
func (s *Server) getMetadataReplyInbox() string { _ = "STUB: not implemented"; return "" }

// getPartitionNotificationInbox returns the NATS subject used for leaders to
// indicate new data is available on a partition for a follower to replicate if
// the follower is idle.
func (s *Server) getPartitionNotificationInbox(id string) string {
	_ = "STUB: not implemented"
	return ""
}

// getAckInbox returns a random NATS subject to use for publish acks scoped to
// the cluster namespace.
func (s *Server) getAckInbox() string { _ = "STUB: not implemented"; return "" }

// getActivityStreamSubject returns the NATS subject used for publishing
// activity stream events.
func (s *Server) getActivityStreamSubject() string { _ = "STUB: not implemented"; return "" }

// getCursorStreamSubject returns the NATS subject used for storing consumer
// partition cursors.
func (s *Server) getCursorStreamSubject() string { _ = "STUB: not implemented"; return "" }

// startGoroutine starts a goroutine which is managed by the server. This adds
// the goroutine to a WaitGroup so that the server can wait for all running
// goroutines to stop on shutdown. This should be used instead of a "naked"
// goroutine.
func (s *Server) startGoroutine(f func()) { _ = "STUB: not implemented"; return }

// startGoroutineWG starts a goroutine which is managed by the server and calls
// Done() on the provided WaitGroup upon completion. This adds the goroutine to
// a WaitGroup so that the server can wait for all running goroutines to stop
// on shutdown. This should be used instead of a "naked" goroutine.
func (s *Server) startGoroutineWG(f func(), wg sync.WaitGroup) { _ = "STUB: not implemented"; return }

// startGoroutineWithArgs starts a goroutine which is managed by the server and
// is passed the provided arguments. This adds the goroutine to a WaitGroup so
// that the server can wait for all running goroutines to stop on shutdown.
// This should be used instead of a "naked" goroutine.
func (s *Server) startGoroutineWithArgs(f func(...interface{}), args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

// startGoroutineWithArgsWG starts a goroutine which is managed by the server
// and is passed the provided arguments and calls Done() on the provided
// WaitGroup upon completion. This adds the goroutine to a WaitGroup so that
// the server can wait for all running goroutines to stop on shutdown. This
// should be used instead of a "naked" goroutine.
func (s *Server) startGoroutineWithArgsWG(f func(...interface{}), wg sync.WaitGroup, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}
