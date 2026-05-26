package server

import (
	"context"
	"hash/crc32"
	"sync"
	"time"

	"github.com/nats-io/nats.go"
	"google.golang.org/grpc/status"

	client "github.com/liftbridge-io/liftbridge-api/v2/go"
	proto "github.com/liftbridge-io/liftbridge/server/protocol"
)

const (
	waitForNewMessages int64 = -1
	asyncAckTimeout          = 5 * time.Second
)

var hasher = crc32.ChecksumIEEE

// apiServer implements the gRPC server interface clients interact with.
type apiServer struct {
	client.UnimplementedAPIServer
	*Server
}

// enforce authorization policy per action/subject/object
func (a *apiServer) enforcePolicy(subject, object, action string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// CreateStream creates a new stream attached to a NATS subject. It returns an
// AlreadyExists status code if a stream with the given subject and name
// already exists.
func (a *apiServer) CreateStream(ctx context.Context, req *client.CreateStreamRequest) (
	*client.CreateStreamResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DeleteStream deletes a stream attached to a NATS subject.
func (a *apiServer) DeleteStream(ctx context.Context, req *client.DeleteStreamRequest) (
	*client.DeleteStreamResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// PauseStream pauses a stream's partitions. If no partitions are specified,
// all of the stream's partitions will be paused. Partitions are resumed when
// they are published to via the Liftbridge Publish API.
func (a *apiServer) PauseStream(ctx context.Context, req *client.PauseStreamRequest) (
	*client.PauseStreamResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SetStreamReadonly sets the readonly status on a stream's partitions. If no
// partitions are specified, all of the stream's partitions will have their
// readonly status set.
func (a *apiServer) SetStreamReadonly(ctx context.Context, req *client.SetStreamReadonlyRequest) (
	*client.SetStreamReadonlyResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Subscribe creates an ephemeral subscription for the given stream partition.
// It begins to receive messages starting at the given offset and waits for new
// messages when it reaches the end of the partition. If the subscriber is part
// of a consumer group, this will ensure only one member of the group is
// subscribed to a given partition at a time. Use the request context to close
// the subscription.
func (a *apiServer) Subscribe(req *client.SubscribeRequest, out client.API_SubscribeServer) error {
	_ = "STUB: not implemented"
	return nil
}

// Send an empty message which signals the subscription was successfully
// created.

// SubscribeInternal creates an ephemeral subscription for the given stream
// partition. It begins to receive messages starting at the given offset and
// waits for new messages when it reaches the end of the partition. If the
// subscriber is part of a consumer group, this will ensure only one member of
// the group is subscribed to a given partition at a time. Use the request
// context to close the subscription. This is a non-gRPC API for internal use.
func (a *apiServer) SubscribeInternal(ctx context.Context, req *client.SubscribeRequest) (
	*subscription, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Consumer groups are not compatible with ReadISRReplica.

// FetchMetadata retrieves the latest cluster metadata, including stream broker
// information.
func (a *apiServer) FetchMetadata(ctx context.Context, req *client.FetchMetadataRequest) (
	*client.FetchMetadataResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// FetchPartitionMetadata retrieves metatadata from the partition leader. This
// is mainly useful when client would like to know the high watermark and
// newest offset for a partition.
func (a *apiServer) FetchPartitionMetadata(ctx context.Context, req *client.FetchPartitionMetadataRequest) (
	*client.FetchPartitionMetadataResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Publish a new message to a stream. If the AckPolicy is not NONE and a
// deadline is provided, this will synchronously block until the ack is
// received. If the ack is not received in time, a DeadlineExceeded status code
// is returned. A FailedPrecondition status code is returned if the partition is
// readonly.
func (a *apiServer) Publish(ctx context.Context, req *client.PublishRequest) (
	*client.PublishResponse, error) {
	_ = "STUB: not implemented"

	// TODO: Deprecate in favor of PublishAsync and log a warning.
	return nil, nil
}

// Asynchronously publish messages to a stream. This returns a stream which
// will yield PublishResponses for messages whose AckPolicy is not NONE.
func (a *apiServer) PublishAsync(stream client.API_PublishAsyncServer) error {
	_ = "STUB: not implemented"
	return nil
}

// Publish a Liftbridge message to a NATS subject. If the AckPolicy is not NONE
// and a deadline is provided, this will synchronously block until the first
// ack is received. If an ack is not received in time, a DeadlineExceeded
// status code is returned.
func (a *apiServer) PublishToSubject(ctx context.Context, req *client.PublishToSubjectRequest) (
	*client.PublishToSubjectResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SetCursor stores a cursor position for a particular stream partition which
// is uniquely identified by an opaque string.
//
// NOTE: This is a beta endpoint and is subject to change. It is not included
// as part of Liftbridge's semantic versioning scheme.
func (a *apiServer) SetCursor(ctx context.Context, req *client.SetCursorRequest) (
	*client.SetCursorResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// FetchCursor retrieves a partition cursor position.
//
// NOTE: This is a beta endpoint and is subject to change. It is not included
// as part of Liftbridge's semantic versioning scheme.
func (a *apiServer) FetchCursor(ctx context.Context, req *client.FetchCursorRequest) (
	*client.FetchCursorResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// JoinConsumerGroup adds a consumer to a consumer group. If the group does not
// exist, it will create it first.
//
// NOTE: This is a beta endpoint and is subject to change. It is not included
// as part of Liftbridge's semantic versioning scheme.
func (a *apiServer) JoinConsumerGroup(ctx context.Context, req *client.JoinConsumerGroupRequest) (
	*client.JoinConsumerGroupResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// LeaveConsumerGroup removes a consumer from a consumer group.
//
// NOTE: This is a beta endpoint and is subject to change. It is not included
// as part of Liftbridge's semantic versioning scheme.
func (a *apiServer) LeaveConsumerGroup(ctx context.Context, req *client.LeaveConsumerGroupRequest) (
	*client.LeaveConsumerGroupResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// FetchConsumerGroupAssignments retrieves the partition assignments for a
// consumer. This also acts as a heartbeat for the consumer so that the
// coordinator keeps the consumer active in the group.
//
// NOTE: This is a beta endpoint and is subject to change. It is not included
// as part of Liftbridge's semantic versioning scheme.
func (a *apiServer) FetchConsumerGroupAssignments(ctx context.Context, req *client.FetchConsumerGroupAssignmentsRequest) (
	*client.FetchConsumerGroupAssignmentsResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ReportConsumerGroupCoordinator reports a consumer group coordinator as
// failed. If a majority of the group's members report the coordinator within a
// bounded period, the cluster will select a new coordinator.
//
// NOTE: This is a beta endpoint and is subject to change. It is not included
// as part of Liftbridge's semantic versioning scheme.
func (a *apiServer) ReportConsumerGroupCoordinator(ctx context.Context, req *client.ReportConsumerGroupCoordinatorRequest) (
	*client.ReportConsumerGroupCoordinatorResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// isValidSubject indicates if the string is a valid NATS subject.
func isValidSubject(subj string) bool { _ = "STUB: not implemented"; return false }

func (a *apiServer) ensureAuthorizationPermission(ctx context.Context, stream, apiMethod string) error {
	_ = "STUB: not implemented"
	// Verify authorization permissions
	return nil
}

func (a *apiServer) ensureCreateStreamPrecondition(req *client.CreateStreamRequest) *status.Status {
	_ = "STUB: not implemented"
	// Verify if an encrypted stream is requested, the
	// LIFTBRIDGE_ENCRYPTION_KEY is correctly set.
	return nil
}

func (a *apiServer) ensurePublishPreconditions(req *client.PublishRequest) *client.PublishAsyncError {
	_ = "STUB: not implemented"
	return nil
}

// Verify stream exists

// Verify partition exists

// Verify stream is not read only

// Verify AckPolicy is set for streams with Optimistic Concurrency Control

func (a *apiServer) resumeStream(ctx context.Context, streamName string, partitionID int32) error {
	_ = "STUB: not implemented"
	return nil
}

// If ResumeAll is enabled, resume any paused partitions in the stream.

// Otherwise just resume the partition being published to if it's
// paused.

// Reset the ResumeAll flag on the stream.

func (a *apiServer) getPublishSubject(req *client.PublishRequest) (string, *client.PublishAsyncError) {
	_ = "STUB: not implemented"
	return "", nil
}

func (a *apiServer) publish(ctx context.Context, subject, ackInbox string,
	ackPolicy client.AckPolicy, msg *client.Message) (*client.Ack, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// If AckPolicy is NONE or a timeout isn't specified, then we will fire and
// forget.

// Otherwise we need to publish and wait for the ack.

func (a *apiServer) publishSync(ctx context.Context, subject,
	ackInbox string, msg []byte) (*client.Ack, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// subscribe sets up a subscription on the given partition and begins sending
// messages on the returned channel. The subscription will run until the cancel
// channel is closed, the context is canceled, or an error is returned
// asynchronously on the status channel.
func (a *apiServer) subscribe(ctx context.Context, partition *partition,
	req *client.SubscribeRequest) (*subscription, *status.Status) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Resuming a partition creates a new one, so we have to get a pointer
// to it.

func getStreamConfig(req *client.CreateStreamRequest) *proto.StreamConfig {
	_ = "STUB: not implemented"
	return nil
}

func convertPublishAsyncError(err *client.PublishAsyncError) error {
	_ = "STUB: not implemented"
	return nil
}

func convertAckError(ackError client.Ack_Error) *client.PublishAsyncError {
	_ = "STUB: not implemented"
	return nil
}

// publishAsyncSession maintains state for long-lived PublishAsync RPCs.
type publishAsyncSession struct {
	*apiServer
	mu       sync.Mutex
	inflight int32
	stream   client.API_PublishAsyncServer
	ackInbox string
	sub      *nats.Subscription
}

func (a *apiServer) newPublishAsyncSession(stream client.API_PublishAsyncServer) *publishAsyncSession {
	_ = "STUB: not implemented"
	return nil
}

// dispatchAcks sets up a subscription on the ack inbox to dispatch acks for
// published messages back to the client.
func (p *publishAsyncSession) dispatchAcks() error { _ = "STUB: not implemented"; return nil }

// publishLoop is a long-lived loop that receives messages from the client and
// publishes them. It returns nil on completion or an error which is terminal.
// If the client closes the stream, this will attempt to wait for remaining
// acks for any in-flight messages before ending the session.
func (p *publishAsyncSession) publishLoop() error { _ = "STUB: not implemented"; return nil }

// Increment in-flight count if we're expecting an ack.

// sendPublishAsyncError sends a PublishResponse containing an error back to
// the client.
func (p *publishAsyncSession) sendPublishAsyncError(correlationID string, err *client.PublishAsyncError) {
	_ = "STUB: not implemented"
	return
}

// Set an Ack with an empty correlation id so we don't break older
// clients that are unaware of AsyncError. TODO (2.0.0): Remove when
// clients are expected to check for AsyncError.

// waitForInflight attempts to wait for remaining acks for any in-flight
// messages.
func (p *publishAsyncSession) waitForInflight() { _ = "STUB: not implemented"; return }

func (p *publishAsyncSession) close() { _ = "STUB: not implemented"; return }

// isReservedStream indicates if the provided stream name is a reserved stream.
func isReservedStream(stream string) bool { _ = "STUB: not implemented"; return false }
