package server

import (
	"github.com/nats-io/nats.go"

	proto "github.com/liftbridge-io/liftbridge/server/protocol"
)

// getPropagateInbox returns the NATS subject used for handling propagated Raft
// operations. The server subscribes to this when it is the metadata leader.
// Followers can then forward operations for the leader to apply.
func (s *Server) getPropagateInbox() string { _ = "STUB: not implemented"; return "" }

// handlePropagatedRequest is a NATS handler used to process propagated Raft
// operations from followers in the Raft cluster. This is activated when the
// server becomes the metadata leader. If, for some reason, the server receives
// a forwarded operation and loses leadership at the same time, the operation
// will fail when it's proposed to the Raft cluster.
func (s *Server) handlePropagatedRequest(m *nats.Msg) { _ = "STUB: not implemented"; return }

func (s *Server) handleCreateStream(req *proto.PropagatedRequest) *proto.PropagatedResponse {
	_ = "STUB: not implemented"
	return nil
}

func (s *Server) handleShrinkISR(req *proto.PropagatedRequest) *proto.PropagatedResponse {
	_ = "STUB: not implemented"
	return nil
}

func (s *Server) handleExpandISR(req *proto.PropagatedRequest) *proto.PropagatedResponse {
	_ = "STUB: not implemented"
	return nil
}

func (s *Server) handleReportLeader(req *proto.PropagatedRequest) *proto.PropagatedResponse {
	_ = "STUB: not implemented"
	return nil
}

func (s *Server) handleDeleteStream(req *proto.PropagatedRequest) *proto.PropagatedResponse {
	_ = "STUB: not implemented"
	return nil
}

func (s *Server) handlePauseStream(req *proto.PropagatedRequest) *proto.PropagatedResponse {
	_ = "STUB: not implemented"
	return nil
}

func (s *Server) handleResumeStream(req *proto.PropagatedRequest) *proto.PropagatedResponse {
	_ = "STUB: not implemented"
	return nil
}

func (s *Server) handleSetStreamReadonly(req *proto.PropagatedRequest) *proto.PropagatedResponse {
	_ = "STUB: not implemented"
	return nil
}

func (s *Server) handleJoinConsumerGroup(req *proto.PropagatedRequest) *proto.PropagatedResponse {
	_ = "STUB: not implemented"
	return nil
}

func (s *Server) handleLeaveConsumerGroup(req *proto.PropagatedRequest) *proto.PropagatedResponse {
	_ = "STUB: not implemented"
	return nil
}

func (s *Server) handleReportConsumerGroupCoordinator(req *proto.PropagatedRequest) *proto.PropagatedResponse {
	_ = "STUB: not implemented"
	return nil
}
