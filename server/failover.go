package server

import (
	"context"
	"sync"
	"time"

	"google.golang.org/grpc/status"
)

type failoverExpiredHandler func()

type failoverHandler func(context.Context) *status.Status

// failover represents a resource that is controlled by a leader that can be
// failed over to another broker, such as a stream partition or consumer group.
type failover interface {
	// Quorum returns the number of failure reports needed from distinct
	// witnesses in order to trigger a failover to a new leader.
	Quorum() int

	// Timeout returns the time elapsed before expiring a failover. Each time a
	// report is made, the failover's timeout is reset. Upon timing out, the
	// timer for the leader failover is removed.
	Timeout() time.Duration

	// OnExpired is invoked when the timeout is reached indicating the failover
	// has expired.
	OnExpired()

	// Failover selects a new leader. It returns a Status if selecting a new
	// leader fails.
	Failover(context.Context) *status.Status
}

// failoverStatus tracks witnesses for a leader failover. Witnesses are
// replicas which have reported the leader as unresponsive. If a quorum of
// witnesses report the leader within a bounded period of time, the controller
// will select a new leader.
type failoverStatus struct {
	mu        sync.Mutex
	failover  failover
	timer     *time.Timer
	witnesses map[string]struct{}
}

func newFailoverStatus(f failover) *failoverStatus { _ = "STUB: not implemented"; return nil }

// report adds the given witness to the failoverStatus witnesses. If a quorum
// of witnesses have reported the leader, a new leader will be selected.
// Otherwise, the expiration timer is reset. A Status is returned if selecting
// a new leader fails.
func (f *failoverStatus) report(ctx context.Context, witness string) *status.Status {
	_ = "STUB: not implemented"
	return nil
}

// cancel stops the expiration timer, if there is one.
func (f *failoverStatus) cancel() { _ = "STUB: not implemented"; return }

// partitionFailover implements the failover interface for a stream partition
// leader. When a majority of a partition's ISR report the leader as failed, a
// new leader is selected.
type partitionFailover struct {
	partition  *partition
	timeout    time.Duration
	onExpired  failoverExpiredHandler
	onFailover failoverHandler
}

func newPartitionFailoverStatus(partition *partition, timeout time.Duration,
	onExpired failoverExpiredHandler, onFailover failoverHandler) *failoverStatus {
	_ = "STUB: not implemented"
	return nil
}

// Quorum returns (partition ISR size - 1) / 2. One is subtracted from the ISR
// size to exclude the leader.
func (p *partitionFailover) Quorum() int { _ = "STUB: not implemented"; return 0 }

// Timeout returns the configured ReplicaMaxLeaderTimeout.
func (p *partitionFailover) Timeout() time.Duration {
	_ = "STUB: not implemented"

	// OnExpired expires the failover.
	return *new(time.Duration)
}

func (p *partitionFailover) OnExpired() {
	_ = "STUB: not implemented"

	// Failover selects a new leader.
	return
}

func (p *partitionFailover) Failover(ctx context.Context) *status.Status {
	_ = "STUB: not implemented"
	return nil

	// groupFailover implements the failover interface for a consumer group
	// coordinator. When a majority of a consumer group's members report the
	// coordinator as failed, a new coordinator is selected.
}

type groupFailover struct {
	group      *consumerGroup
	timeout    time.Duration
	onExpired  failoverExpiredHandler
	onFailover failoverHandler
}

func newGroupFailoverStatus(group *consumerGroup, timeout time.Duration,
	onExpired failoverExpiredHandler, onFailover failoverHandler) *failoverStatus {
	_ = "STUB: not implemented"
	return nil
}

// Quorum returns members / 2.
func (g *groupFailover) Quorum() int { _ = "STUB: not implemented"; return 0 }

// Timeout returns the configured GroupsCoordinatorTimeout.
func (g *groupFailover) Timeout() time.Duration {
	_ = "STUB: not implemented"

	// OnExpired expires the failover.
	return *new(time.Duration)
}

func (g *groupFailover) OnExpired() {
	_ = "STUB: not implemented"

	// Failover selects a new coordinator.
	return
}

func (g *groupFailover) Failover(ctx context.Context) *status.Status {
	_ = "STUB: not implemented"
	return nil
}
