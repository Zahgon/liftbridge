package server

import (
	"context"
	"sync"
	"time"

	"github.com/pkg/errors"
	"google.golang.org/grpc/status"

	client "github.com/liftbridge-io/liftbridge-api/v2/go"

	proto "github.com/liftbridge-io/liftbridge/server/protocol"
)

const (
	defaultPropagateTimeout             = 5 * time.Second
	defaultFetchBrokerInfoTimeout       = 3 * time.Second
	maxReplicationFactor          int32 = -1
)

var (
	// ErrStreamExists is returned by CreateStream when attempting to create a
	// stream that already exists.
	ErrStreamExists = errors.New("stream already exists")

	// ErrStreamNotFound is returned by DeleteStream/PauseStream when
	// attempting to delete/pause a stream that does not exist.
	ErrStreamNotFound = errors.New("stream does not exist")

	// ErrPartitionNotFound is returned by PauseStream when attempting to pause
	// a stream partition that does not exist.
	ErrPartitionNotFound = errors.New("partition does not exist")

	// ErrConsumerGroupExists is returned by createConsumerGroup when
	// attempting to create a group that already exists.
	ErrConsumerGroupExists = errors.New("consumer group already exists")

	// ErrConsumerGroupNotFound is returned by JoinConsumerGroup when
	// attempting to join a group that does not exist.
	ErrConsumerGroupNotFound = errors.New("consumer group does not exist")

	// ErrConsumerAlreadyMember is returned by JoinConsumerGroup when the
	// consumer is already a member of the group.
	ErrConsumerAlreadyMember = errors.New("consumer is already a member of the consumer group")

	// ErrConsumerNotMember is returned by LeaveConsumerGroup when the
	// consumer if not a member of the group.
	ErrConsumerNotMember = errors.New("consumer is not a member of the consumer group")

	// ErrBrokerNotCoordinator is returned by GetConsumerGroupAssignments when
	// this server is not the coordinator for the requested consumer group.
	ErrBrokerNotCoordinator = errors.New("broker is not the consumer group coordinator")

	// ErrGroupEpoch is returned by GetConsumerGroupAssignments when the
	// client-provided group epoch differs from the server-side group epoch.
	ErrGroupEpoch = errors.New("client-provided group epoch differs from broker group epoch")
)

// metadataAPI is the internal API for interacting with cluster data. All
// stream access should go through the exported methods of the metadataAPI.
type metadataAPI struct {
	*Server
	streams            map[string]*stream
	mu                 sync.RWMutex
	partitionFailovers map[*partition]*failoverStatus
	cachedBrokers      []*client.Broker
	cachedServerIDs    map[string]struct{}
	lastCached         time.Time
	consumerGroupsMu   sync.RWMutex
	consumerGroups     map[string]*consumerGroup
	groupFailovers     map[*consumerGroup]*failoverStatus
	stats              struct {
		sync.RWMutex
		brokerLeaderLoad      map[string]int
		brokerPartitionLoad   map[string]int
		brokerCoordinatorLoad map[string]int
	}
}

func newMetadataAPI(s *Server) *metadataAPI { _ = "STUB: not implemented"; return nil }

// BrokerPartitionCounts returns a map of broker IDs to the number of
// partitions they are hosting.
func (m *metadataAPI) BrokerPartitionCounts() map[string]int { _ = "STUB: not implemented"; return nil }

// BrokerLeaderCounts returns a map of broker IDs to the number of
// partitions they are leading.
func (m *metadataAPI) BrokerLeaderCounts() map[string]int { _ = "STUB: not implemented"; return nil }

// FetchMetadata retrieves the cluster metadata for the given request. If the
// request specifies streams, it will only return metadata for those particular
// streams. If not, it will return metadata for all streams.
func (m *metadataAPI) FetchMetadata(ctx context.Context, req *client.FetchMetadataRequest) (
	*client.FetchMetadataResponse, *status.Status) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Check if we can use cached broker info.

// Query broker info from peers.

// Update the cache.

// FetchPartitionMetadata retrieves the metadata for the partition leader. This
// mainly serves the purpose of returning high watermark and newest offset.
func (m *metadataAPI) FetchPartitionMetadata(ctx context.Context, req *client.FetchPartitionMetadataRequest) (
	*client.FetchPartitionMetadataResponse, *status.Status) {
	_ = "STUB: not implemented"
	return nil, nil
}

// brokerCache checks if the cache of broker metadata is clean and, if it is
// and it's not past the metadata cache max age, returns the cached broker
// list. The bool returned indicates if the cached data is returned or not.
func (m *metadataAPI) brokerCache(serverIDs map[string]struct{}) ([]*client.Broker, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// fetchBrokerInfo retrieves the broker metadata for the cluster. The numPeers
// argument is the expected number of peers to get a response from.
func (m *metadataAPI) fetchBrokerInfo(ctx context.Context, numPeers int) ([]*client.Broker, *status.Status) {
	_ = "STUB: not implemented"
	// Brokers load data
	return nil, nil
}

// Add ourselves.

// Make sure there is a deadline on the request.

// Create subscription to receive responses on.

// Survey the cluster.

// Gather responses.

// createMetadataResponse creates a FetchMetadataResponse and populates it with
// stream and group metadata. If the provided list of stream names is empty, it
// will populate metadata for all streams. Otherwise, it populates only the
// specified streams.
func (m *metadataAPI) createMetadataResponse(streams, groups []string) *client.FetchMetadataResponse {
	_ = "STUB: not implemented"
	// If no stream names were provided, fetch metadata for all streams.
	return nil
}

// Stream does not exist.

// Group does not exist.

// CreateStream creates a new stream if this server is the metadata leader. If
// it is not, it will forward the request to the leader and return the
// response. This operation is replicated by Raft. The metadata leader will
// select replicationFactor nodes to participate and a leader for each
// partition.  If successful, this will return once the partitions have been
// replicated to the cluster and the partition leaders have started.
func (m *metadataAPI) CreateStream(ctx context.Context, req *proto.CreateStreamOp) *status.Status {
	_ = "STUB: not implemented"
	// Forward the request if we're not the leader.
	return nil
}

// If we have since become leader, continue on with the request.

// Select replicationFactor nodes to participate in the partition.

// Select a leader for the partition.

// Replicate stream create through Raft.

// Wait on result of replication.

// Wait for leaders to create partitions (best effort).

// DeleteStream deletes a stream if this server is the metadata leader. If it is
// not, it will forward the request to the leader and return the response. This
// operation is replicated by Raft. If successful, this will return once the
// stream has been deleted from the cluster.
func (m *metadataAPI) DeleteStream(ctx context.Context, req *proto.DeleteStreamOp) *status.Status {
	_ = "STUB: not implemented"
	// Forward the request if we're not the leader.
	return nil
}

// If we have since become leader, continue on with the request.

// Replicate partition deletion through Raft.

// Wait on result of deletion.

// PauseStream pauses a stream if this server is the metadata leader. If it is
// not, it will forward the request to the leader and return the response. This
// operation is replicated by Raft. If successful, this will return once the
// stream has been paused.
func (m *metadataAPI) PauseStream(ctx context.Context, req *proto.PauseStreamOp) *status.Status {
	_ = "STUB: not implemented"
	// Forward the request if we're not the leader.
	return nil
}

// If we have since become leader, continue on with the request.

// Replicate stream pausing through Raft.

// Wait on result of pausing.

// ResumeStream unpauses a stream partition(s) if this server is the metadata
// leader. If it is not, it will forward the request to the leader and return
// the response. This operation is replicated by Raft. Resume is intended to
// be idempotent. If the partition is already resumed when this is called, this
// will return nil. If the partition to resume is not specified on the request,
// this will resume all paused partitions in the stream.
func (m *metadataAPI) ResumeStream(ctx context.Context, req *proto.ResumeStreamOp) *status.Status {
	_ = "STUB: not implemented"
	// Forward the request if we're not the leader.
	return nil
}

// If we have since become leader, continue on with the request.

// Replicate stream resume through Raft.

// Wait on result of replication.

// Wait for leader to resume partition(s) (best effort).

// ShrinkISR removes the specified replica from the partition's in-sync
// replicas set if this server is the metadata leader. If it is not, it will
// forward the request to the leader and return the response. This operation is
// replicated by Raft.
func (m *metadataAPI) ShrinkISR(ctx context.Context, req *proto.ShrinkISROp) *status.Status {
	_ = "STUB: not implemented"
	// Forward the request if we're not the leader.
	return nil
}

// If we have since become leader, continue on with the request.

// Verify the partition exists.

// Check the leader epoch.

// Replicate ISR shrink through Raft.

// Wait on result of replication.

// ExpandISR adds the specified replica to the partition's in-sync replicas set
// if this server is the metadata leader. If it is not, it will forward the
// request to the leader and return the response. This operation is replicated
// by Raft.
func (m *metadataAPI) ExpandISR(ctx context.Context, req *proto.ExpandISROp) *status.Status {
	_ = "STUB: not implemented"
	// Forward the request if we're not the leader.
	return nil
}

// If we have since become leader, continue on with the request.

// Verify the partition exists.

// Check the leader epoch.

// Replicate ISR expand through Raft.

// Wait on result of replication.

// ReportLeader marks the partition leader as unresponsive with respect to the
// specified replica if this server is the metadata leader. If it is not, it
// will forward the request to the leader and return the response. If a quorum
// of replicas report the partition leader within a bounded period, the
// metadata leader will select a new partition leader.
func (m *metadataAPI) ReportLeader(ctx context.Context, req *proto.ReportLeaderOp) *status.Status {
	_ = "STUB: not implemented"
	// Forward the request if we're not the leader.
	return nil
}

// If we have since become leader, continue on with the request.

// Verify the partition exists.

// Check the leader epoch.

func (m *metadataAPI) newPartitionFailoverExpiredHandler(p *partition) failoverExpiredHandler {
	_ = "STUB: not implemented"
	return *new(failoverExpiredHandler)
}

func (m *metadataAPI) newPartitionFailoverHandler(p *partition) failoverHandler {
	_ = "STUB: not implemented"
	return *new(failoverHandler)
}

// SetStreamReadonly sets a stream's readonly flag if this server is the
// metadata leader. If it is not, it will forward the request to the leader and
// return the response. This operation is replicated by Raft. If successful,
// this will return once the readonly flag has been set.
func (m *metadataAPI) SetStreamReadonly(ctx context.Context, req *proto.SetStreamReadonlyOp) *status.Status {
	_ = "STUB: not implemented"
	// Forward the request if we're not the leader.
	return nil
}

// If we have since become leader, continue on with the request.

// Replicate the stream readonly flag through Raft.

// Wait on result of setting the readonly flag.

// JoinConsumerGroup adds a consumer to a consumer group if this server is the
// metadata leader. The group is created first if it does not yet exist. If
// this server is not the metadata leader, it will forward the request to the
// leader and return the response. This operation is replicated by Raft. If
// successful, this will return once the consumer has been added to the group.
// Returns the group coordinator ID and coordinator epoch on success.
func (m *metadataAPI) JoinConsumerGroup(ctx context.Context, req *proto.JoinConsumerGroupOp) (
	string, uint64, *status.Status) {
	_ = "STUB: not implemented"

	// Forward the request if we're not the leader.
	return "", 0, nil
}

// If we have since become leader, continue on with the request.

// Check if group exists. If it doesn't, create it with the member.

// The coordinator epoch for a new group is always 0.

// If the group already existed, replicate the join request through Raft.

// Wait on result of replication.

// LeaveConsumerGroup removes a consumer from a consumer group. If this is the
// last member of the group, the group will be deleted. This operation is
// replicated by Raft. If successful, this will return once the consumer has
// been removed from the group.
func (m *metadataAPI) LeaveConsumerGroup(ctx context.Context, req *proto.LeaveConsumerGroupOp) *status.Status {
	_ = "STUB: not implemented"
	// Forward the request if we're not the leader.
	return nil
}

// If we have since become leader, continue on with the request.

// Replicate the leave request through Raft.

// Wait on result of replication.

// ReportGroupCoordinator marks the consumer group coordinator as unresponsive
// with respect to the specified member if this server is the metadata leader.
// If it is not, it will forward the request to the leader and return the
// response. If a quorum of members report the coordinator within a bounded
// period, the metadata leader will select a new group coordinator.
func (m *metadataAPI) ReportGroupCoordinator(ctx context.Context, req *proto.ReportConsumerGroupCoordinatorOp) *status.Status {
	_ = "STUB: not implemented"
	// Forward the request if we're not the leader.
	return nil
}

// If we have since become leader, continue on with the request.

// Verify the group exists.

// Check the group epoch.

// Ensure the consumer is actually a member.

func (m *metadataAPI) newGroupFailoverExpiredHandler(g *consumerGroup) failoverExpiredHandler {
	_ = "STUB: not implemented"
	return *new(failoverExpiredHandler)
}

func (m *metadataAPI) newGroupFailoverHandler(g *consumerGroup) failoverHandler {
	_ = "STUB: not implemented"
	return *new(failoverHandler)
}

// createConsumerGroup creates a new consumer group bootstrapped with the
// provided consumer by replicating the operation via Raft. This should be
// called after checking that this server is the metadata leader, but if the
// server is not the leader or the leadership has since changed, the operation
// will return an error. This will also return an error if the group being
// created already exists. This returns the group coordinator on success.
func (m *metadataAPI) createConsumerGroup(ctx context.Context, req *proto.JoinConsumerGroupOp) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Replicate the group create through Raft.

// Wait on result of replication.

// AddConsumerGroup adds the given consumer group to the metadata store. It
// returns an error if a consumer group with the same ID already exists.
func (m *metadataAPI) AddConsumerGroup(protoGroup *proto.ConsumerGroup, recovered bool) (*consumerGroup, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Update broker load counts.

// removeConsumerGroupMember sends a LeaveConsumerGroup request to the
// controller to remove the expired consumer from the group.
func (m *metadataAPI) removeConsumerGroupMember(groupID, consumerID string) error {
	_ = "STUB: not implemented"
	return nil
}

// countStreamPartitions returns the number of partitions for the stream or 0
// if the stream does not exist.
func (m *metadataAPI) countStreamPartitions(streamName string) int32 {
	_ = "STUB: not implemented"
	return 0
}

// removeConsumerGroup removes the given consumer group from the metadata
// store. This should only be called within the scope of the consumerGroupsMu.
func (m *metadataAPI) removeConsumerGroup(groupID string) { _ = "STUB: not implemented"; return }

// Update broker load counts.

// AddConsumerToGroup adds the given consumer to the consumer group. It returns
// an error if the group does not exist, the consumer is already a member of
// the group, or any of the provided streams do not exist.
func (m *metadataAPI) AddConsumerToGroup(groupID, consumerID string, streams []string, epoch uint64) error {
	_ = "STUB: not implemented"
	return nil
}

// RemoveConsumerFromGroup removes the given consumer from the consumer group.
// It returns an error if the group does not exist or the consumer is not a
// member of the group. If this is the last member of the group, the group will
// be deleted. Returns a bool indicating if this was the last member of the
// group and the group has been deleted.
func (m *metadataAPI) RemoveConsumerFromGroup(groupID, consumerID string, epoch uint64) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// If the last member was removed, delete the group.

// GetConsumerGroupAssignments returns the group's partition assignments for
// the given consumer and the group epoch.
func (m *metadataAPI) GetConsumerGroupAssignments(groupID, consumerID string, epoch uint64) (
	partitionAssignments, uint64, error) {
	_ = "STUB: not implemented"
	return *new(partitionAssignments), 0, nil
}

// AddStream adds the given stream and its partitions to the metadata store. It
// returns an error if a stream with the same name or any partitions with the
// same ID for the stream already exist. If the stream is recovered, this will
// not start the partitions until recovery completes. Partitions will also not
// be started if they are currently paused.
func (m *metadataAPI) AddStream(protoStream *proto.Stream, recovered bool, epoch uint64) (*stream, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// If this operation is being applied during recovery, check if this
// stream is tombstoned, i.e. was marked for deletion previously. If it
// is, un-tombstone it by closing the existing stream and then
// recreating it, leaving the existing data intact.

// This is an invalid state because it means the stream already
// exists.

// Un-tombstone by closing the existing stream and removing it from
// the streams store so that it can be recreated.

// Update broker load counts.

func (m *metadataAPI) addPartition(stream *stream, protoPartition *proto.Partition, recovered bool, config *proto.StreamConfig) error {
	_ = "STUB: not implemented"
	return nil
}

// Partition already exists for stream.

// This will initialize/recover the durable commit log.

// If we're loading a partition that was paused, we need to re-pause it.

// Start leader/follower loop if necessary.

// ResumePartition unpauses the given stream partition in the metadata store.
// It returns ErrPartitionNotFound if there is no partition with the ID for the
// stream. If the partition is recovered, this will not start the partition
// until recovery completes.
func (m *metadataAPI) ResumePartition(streamName string, id int32, recovered bool) (*partition, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// If it's not paused, do nothing.

// Resume the partition by replacing it.

// Update latest pause status change timestamp.

// Start leader/follower loop if necessary.

// Update broker load counts.

// RemoveFromISR removes the given replica from the partition's ISR if the
// given epoch is greater than the current epoch.
func (m *metadataAPI) RemoveFromISR(streamName, replica string, partitionID int32, epoch uint64) error {
	_ = "STUB: not implemented"
	return nil
}

// Idempotency check.

// AddToISR adds the given replica to the partition's ISR if the given epoch is
// greater than the current epoch.
func (m *metadataAPI) AddToISR(streamName, replica string, partitionID int32, epoch uint64) error {
	_ = "STUB: not implemented"
	return nil
}

// Idempotency check.

// ChangeLeader changes the partition's leader to the given replica if the
// given epoch is greater than the current epoch.
func (m *metadataAPI) ChangeLeader(streamName, leader string, partitionID int32, epoch uint64) error {
	_ = "STUB: not implemented"
	return nil
}

// Idempotency check.

// Update broker load counts.

// ChangeGroupCoordinator changes the consumer group's coordinator to the given
// broker if the given epoch is greater than the current epoch.
func (m *metadataAPI) ChangeGroupCoordinator(groupID, coordinator string, newEpoch uint64) error {
	_ = "STUB: not implemented"
	return nil
}

// Idempotency check.

// Update broker load counts.

// PausePartitions pauses the given partitions for the stream. If the list of
// partitions is empty, this pauses all partitions.
func (m *metadataAPI) PausePartitions(streamName string, partitions []int32, resumeAll bool) error {
	_ = "STUB: not implemented"
	return nil
}

// Update broker load counts.

// SetReadonly changes the stream partitions' readonly flag in the metadata
// store.
func (m *metadataAPI) SetReadonly(streamName string, partitions []int32, readonly bool) error {
	_ = "STUB: not implemented"
	return nil
}

// GetStreams returns all streams from the metadata store.
func (m *metadataAPI) GetStreams() []*stream { _ = "STUB: not implemented"; return nil }

// GetStream returns the stream with the given name or nil if no such stream
// exists.
func (m *metadataAPI) GetStream(name string) *stream { _ = "STUB: not implemented"; return nil }

// GetPartition returns the stream partition for the given stream and partition
// ID. It returns nil if no such partition exists.
func (m *metadataAPI) GetPartition(streamName string, id int32) *partition {
	_ = "STUB: not implemented"
	return nil
}

// GetConsumerGroups returns all consumer groups from the metadata store.
func (m *metadataAPI) GetConsumerGroups() []*consumerGroup { _ = "STUB: not implemented"; return nil }

func (m *metadataAPI) getConsumerGroups() []*consumerGroup { _ = "STUB: not implemented"; return nil }

// GetConsumerGroup returns the consumer group with the given id. It returns
// nil if no such group exists.
func (m *metadataAPI) GetConsumerGroup(id string) *consumerGroup {
	_ = "STUB: not implemented"
	return nil
}

// Reset closes all streams and consumer groups and clears all existing state
// in the metadata store.
func (m *metadataAPI) Reset() error { _ = "STUB: not implemented"; return nil }

// resetFailovers cancels all in-flight failovers and clears the failover state
// in the metadata store. Both the metadata API and consumer groups mutexes
// must be held when calling this.
func (m *metadataAPI) resetFailovers() { _ = "STUB: not implemented"; return }

// RemoveStream closes the stream, removes it from the metadata store, and
// deletes the associated on-disk data for it. However, if this operation is
// being applied during Raft recovery, this will only mark the stream with a
// tombstone. Tombstoned streams will be deleted after the recovery process
// completes.
func (m *metadataAPI) RemoveStream(stream *stream, recovered bool, epoch uint64) error {
	_ = "STUB: not implemented"
	return nil
}

// If this operation is being applied during recovery, only tombstone the
// stream. We don't want to delete streams until recovery finishes to avoid
// deleting potentially valid data, e.g. in the case of a stream being
// deleted, recreated, and then published to. In this scenario, the
// recreate will un-tombstone the stream.

// Update broker load counts.

// RemoveTombstonedStream closes the tombstoned stream, removes it from the
// metadata store, and deletes the associated on-disk data for it.
func (m *metadataAPI) RemoveTombstonedStream(stream *stream, epoch uint64) error {
	_ = "STUB: not implemented"
	return nil
}

// LostLeadership should be called when the server loses metadata leadership.
// This will cancel in-flight failovers.
func (m *metadataAPI) LostLeadership() { _ = "STUB: not implemented"; return }

// deleteStream deletes the stream and the associated on-disk data for it.
func (m *metadataAPI) deleteStream(stream *stream, epoch uint64) error {
	_ = "STUB: not implemented"
	return nil
}

// Remove the stream data directory

// removeStream removes the stream from the stream store, cancels any
// in-flight failovers for its partitions, and triggers a rebalance of consumer
// group assignments.
func (m *metadataAPI) removeStream(stream *stream, epoch uint64) { _ = "STUB: not implemented"; return }

func (m *metadataAPI) getStreams() []*stream { _ = "STUB: not implemented"; return nil }

// getPartitionReplicas selects replicationFactor replicas to participate in
// the stream partition. Replicas are selected based on the amount of partition
// load they have.
func (m *metadataAPI) getPartitionReplicas(replicationFactor int32) ([]string, *status.Status) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Order servers by partition load.

// getClusterServerIDs returns a list of all the broker IDs in the cluster.
func (m *metadataAPI) getClusterServerIDs() ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// electNewPartitionLeader selects a new leader for the given partition,
// applies this update to the Raft group, and notifies the replica set. This
// will fail if the current broker is not the metadata leader.
func (m *metadataAPI) electNewPartitionLeader(ctx context.Context, partition *partition) *status.Status {
	_ = "STUB: not implemented"
	return nil

	// TODO: add support for "unclean" leader elections.
}

// Select a new leader.

// Replicate leader change through Raft.

// Wait on result of replication.

// electNewGroupCoordinator selects a new coordinator for the given consumer
// group and applies this update to the Raft group. This will fail if the
// current broker is not the metadata leader.
func (m *metadataAPI) electNewGroupCoordinator(ctx context.Context, group *consumerGroup) *status.Status {
	_ = "STUB: not implemented"
	return nil
}

// Select a new coordinator.

// Replicate coordinator change through Raft.

// Wait on result of replication.

// propagateCreateStream forwards a CreateStream request to the metadata
// leader. The bool indicates if this server has since become leader and the
// request should be performed locally. A Status is returned if the propagated
// request failed.
func (m *metadataAPI) propagateCreateStream(ctx context.Context, req *proto.CreateStreamOp) (bool, *status.Status) {
	_ = "STUB: not implemented"
	return false, nil
}

// propagateDeleteStream forwards a DeleteStream request to the metadata
// leader. The bool indicates if this server has since become leader and the
// request should be performed locally. A Status is returned if the propagated
// request failed.
func (m *metadataAPI) propagateDeleteStream(ctx context.Context, req *proto.DeleteStreamOp) (bool, *status.Status) {
	_ = "STUB: not implemented"
	return false, nil
}

// propagatePauseStream forwards a PauseStream request to the metadata
// leader. The bool indicates if this server has since become leader and the
// request should be performed locally. A Status is returned if the propagated
// request failed.
func (m *metadataAPI) propagatePauseStream(ctx context.Context, req *proto.PauseStreamOp) (bool, *status.Status) {
	_ = "STUB: not implemented"
	return false, nil
}

// propagateResumeStream forwards a ResumeStream request to the metadata
// leader. The bool indicates if this server has since become leader and the
// request should be performed locally. A Status is returned if the propagated
// request failed.
func (m *metadataAPI) propagateResumeStream(ctx context.Context, req *proto.ResumeStreamOp) (bool, *status.Status) {
	_ = "STUB: not implemented"
	return false, nil
}

// propagateShrinkISR forwards a ShrinkISR request to the metadata leader. The
// bool indicates if this server has since become leader and the request should
// be performed locally. A Status is returned if the propagated request failed.
func (m *metadataAPI) propagateShrinkISR(ctx context.Context, req *proto.ShrinkISROp) (bool, *status.Status) {
	_ = "STUB: not implemented"
	return false, nil
}

// propagateExpandISR forwards a ExpandISR request to the metadata leader. The
// bool indicates if this server has since become leader and the request should
// be performed locally. A Status is returned if the propagated request failed.
func (m *metadataAPI) propagateExpandISR(ctx context.Context, req *proto.ExpandISROp) (bool, *status.Status) {
	_ = "STUB: not implemented"
	return false, nil
}

// propagateReportLeader forwards a ReportLeader request to the metadata
// leader. The bool indicates if this server has since become leader and the
// request should be performed locally. A Status is returned if the propagated
// request failed.
func (m *metadataAPI) propagateReportLeader(ctx context.Context, req *proto.ReportLeaderOp) (bool, *status.Status) {
	_ = "STUB: not implemented"
	return false, nil
}

// propagateSetStreamReadonly forwards a SetStreamReadonly request to the
// metadata leader. The bool indicates if this server has since become leader
// and the request should be performed locally. A Status is returned if the
// propagated request failed.
func (m *metadataAPI) propagateSetStreamReadonly(ctx context.Context, req *proto.SetStreamReadonlyOp) (bool, *status.Status) {
	_ = "STUB: not implemented"
	return false, nil
}

// propagateJoinConsumerGroup forwards a JoinConsumerGroup request to the
// metadata leader. The bool indicates if this server has since become leader
// and the request should be performed locally. A Status is returned if the
// propagated request failed.
func (m *metadataAPI) propagateJoinConsumerGroup(ctx context.Context, req *proto.JoinConsumerGroupOp) (
	*proto.PropagatedResponse_JoinConsumerGroupResponse, bool, *status.Status) {
	_ = "STUB: not implemented"
	return nil, false, nil
}

// propagateLeaveConsumerGroup forwards a LeaveConsumerGroup request to the
// metadata leader. The bool indicates if this server has since become leader
// and the request should be performed locally. A Status is returned if the
// propagated request failed.
func (m *metadataAPI) propagateLeaveConsumerGroup(ctx context.Context, req *proto.LeaveConsumerGroupOp) (bool, *status.Status) {
	_ = "STUB: not implemented"
	return false, nil
}

// propagateReportGroupCoordinator forwards a ReportGroupCoordinator request to
// the metadata leader. The bool indicates if this server has since become
// leader and the request should be performed locally. A Status is returned if
// the propagated request failed.
func (m *metadataAPI) propagateReportGroupCoordinator(ctx context.Context, req *proto.ReportConsumerGroupCoordinatorOp) (
	bool, *status.Status) {
	_ = "STUB: not implemented"
	return false, nil
}

// propagateRequest forwards a metadata request to the metadata leader. The
// bool indicates if this server has since become leader and the request should
// be performed locally. A Status is returned if the propagated request failed.
func (m *metadataAPI) propagateRequest(ctx context.Context, req *proto.PropagatedRequest) (*proto.PropagatedResponse, bool, *status.Status) {
	_ = "STUB: not implemented"
	// Check if there is currently a metadata leader.
	return nil, false, nil
}

// This server has since become metadata leader, so the request should be
// performed locally.

// waitForMetadataLeader waits up to the deadline specified on the Context
// until a metadata leader is established. If no leader is established in time,
// an error is returned. The bool indicates if this server has become the
// leader. False and a nil error indicates another server has become leader.
func (m *metadataAPI) waitForMetadataLeader(ctx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// Wait up to deadline for a metadata leader to be established.

// waitForPartitionLeader does a best-effort wait for the leader of the given
// partition to create and start the partition.
func (m *metadataAPI) waitForPartitionLeader(ctx context.Context, partition *proto.Partition) {
	_ = "STUB: not implemented"
	return
}

// If we're the partition leader, there's no need to make a status
// request. We can just apply a Raft barrier since the FSM is local.

// The leader hasn't finished creating the partition, so wait a bit
// and retry.

// checkCreateStreamPreconditions checks if the stream to be created already
// exists. If it does, it returns ErrStreamExists. Otherwise, it returns nil.
func (m *metadataAPI) checkCreateStreamPreconditions(op *proto.RaftLog) error {
	_ = "STUB: not implemented"
	return nil
}

// checkDeleteStreamPreconditions checks if the stream being deleted exists. If
// it doesn't, it returns ErrStreamNotFound. Otherwise, it returns nil.
func (m *metadataAPI) checkDeleteStreamPreconditions(op *proto.RaftLog) error {
	_ = "STUB: not implemented"
	return nil
}

// checkPauseStreamPreconditions checks if the stream and partitions being
// paused exist. If the stream doesn't exist, it returns ErrStreamNotFound. If
// one or more specified partitions don't exist, it returns
// ErrPartitionNotFound. Otherwise, it returns nil.
func (m *metadataAPI) checkPauseStreamPreconditions(op *proto.RaftLog) error {
	_ = "STUB: not implemented"
	return nil
}

// checkSetStreamReadonlyPreconditions checks if the stream and partitions being
// set readonly exist. If the stream doesn't exist, it returns
// ErrStreamNotFound. If one or more specified partitions don't exist, it
// returns ErrPartitionNotFound. Otherwise, it returns nil.
func (m *metadataAPI) checkSetStreamReadonlyPreconditions(op *proto.RaftLog) error {
	_ = "STUB: not implemented"
	return nil
}

// checkResumeStreamPreconditions checks if the stream and partitions to be
// resumed exist. If the stream does not exist, it returns ErrStreamNotFound.
// If any partitions do not exist, it returns ErrPartitionNotFound. Otherwise,
// it returns nil.
func (m *metadataAPI) checkResumeStreamPreconditions(op *proto.RaftLog) error {
	_ = "STUB: not implemented"
	return nil
}

// checkShrinkISRPreconditions checks if the partition whose ISR is being
// shrunk exists. If the stream doesn't exist, it returns ErrStreamNotFound. If
// the partition doesn't exist, it returns ErrPartitionNotFound. Otherwise, it
// returns nil.
func (m *metadataAPI) checkShrinkISRPreconditions(op *proto.RaftLog) error {
	_ = "STUB: not implemented"
	return nil
}

// checkExpandISRPreconditions checks if the partition whose ISR is being
// expanded exists. If the stream doesn't exist, it returns ErrStreamNotFound.
// If the partition doesn't exist, it returns ErrPartitionNotFound. Otherwise,
// it returns nil.
func (m *metadataAPI) checkExpandISRPreconditions(op *proto.RaftLog) error {
	_ = "STUB: not implemented"
	return nil
}

// checkChangeLeaderPreconditions checks if the partition whose leader is being
// changed exists. If the stream doesn't exist, it returns ErrStreamNotFound.
// If the partition doesn't exist, it returns ErrPartitionNotFound. Otherwise,
// it returns nil.
func (m *metadataAPI) checkChangeLeaderPreconditions(op *proto.RaftLog) error {
	_ = "STUB: not implemented"
	return nil
}

// checkCreateConsumerGroupPreconditions checks if the group to be created
// already exists. If it does, it returns ErrConsumerGroupExists. If any of the
// initial members' requested streams do not exist, returns ErrStreamNotFound.
// Otherwise, it returns nil.
func (m *metadataAPI) checkCreateConsumerGroupPreconditions(op *proto.RaftLog) error {
	_ = "STUB: not implemented"
	return nil
}

// checkJoinConsumerGroupPreconditions checks if the group to be joined exists.
// If it does not, it returns ErrConsumerGroupNotFound. If the consumer is
// already a member of the group, returns ErrConsumerAlreadyMember. If any of
// the requested streams do not exist, returns ErrStreamNotFound. Otherwise, it
// returns nil.
func (m *metadataAPI) checkJoinConsumerGroupPreconditions(op *proto.RaftLog) error {
	_ = "STUB: not implemented"
	return nil
}

// checkLeaveConsumerGroupPreconditions checks if the group to be joined
// exists. If it does not, it returns ErrConsumerGroupNotFound. If the consumer
// is not a member of the group, returns ErrConsumerNotMember. Otherwise, it
// returns nil.
func (m *metadataAPI) checkLeaveConsumerGroupPreconditions(op *proto.RaftLog) error {
	_ = "STUB: not implemented"
	return nil
}

// checkChangeGroupCoordinatorPreconditions checks if the consumer group whose
// coordinator is being changed exists. If the group doesn't exist, it returns
// ErrConsumerGroupNotFound. Otherwise, it returns nil.
func (m *metadataAPI) checkChangeGroupCoordinatorPreconditions(op *proto.RaftLog) error {
	_ = "STUB: not implemented"
	return nil
}

// partitionExists indicates if the given partition exists in the stream. If
// the stream doesn't exist, it returns ErrStreamNotFound. If the partition
// doesn't exist, it returns ErrPartitionNotFound.
func (m *metadataAPI) partitionExists(streamName string, partitionID int32) error {
	_ = "STUB: not implemented"
	return nil
}

// selectPartitionLeader selects a replica from the list of replicas to act as
// leader by attempting to select the replica with the least partition
// leadership load.
func (m *metadataAPI) selectPartitionLeader(replicas []string) string {
	_ = "STUB: not implemented"
	// Order servers by leader load.
	return ""
}

// selectGroupCoordinator selects a broker to act as a consumer group
// coordinator by attempting to select the broker with the least coordinator
// load.
func (m *metadataAPI) selectGroupCoordinator(candidates []string) string {
	_ = "STUB: not implemented"
	// Order servers by coordinator load.
	return ""
}

// ensureTimeout ensures there is a timeout on the Context. If there is, it
// returns the Context. If there isn't it returns a new Context wrapping the
// provided one with the default timeout applied. It also returns a cancel
// function which must be invoked by the caller in all cases to avoid a Context
// leak.
func ensureTimeout(ctx context.Context, defaultTimeout time.Duration) (context.Context, context.CancelFunc) {
	_ = "STUB: not implemented"
	return *new(context.Context), *new(context.CancelFunc)
}

// eventTimestampsToProto returns a client proto's partition event timestamps
// from a partition's event timestamps struct.
func eventTimestampsToProto(timestamps EventTimestamps) *client.PartitionEventTimestamps {
	_ = "STUB: not implemented"
	return nil
}

// Calling UnixNano() on a zero time is undefined, so we need to make these
// checks.

// getPartitionMetadata returns a partition's metadata.
func getPartitionMetadata(partitionID int32, partition *partition) *client.PartitionMetadata {
	_ = "STUB: not implemented"
	return nil
}
