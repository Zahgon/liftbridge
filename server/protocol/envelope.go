package protocol

import (
	"bytes"
	"encoding/binary"
	"hash/crc32"

	pb "github.com/golang/protobuf/proto"
	client "github.com/liftbridge-io/liftbridge-api/v2/go"
)

// msgType indicates the type of message contained by an envelope.
type msgType byte

const (
	msgTypePublish msgType = iota
	msgTypeAck

	msgTypeReplicationRequest
	msgTypeReplicationResponse

	msgTypeRaftJoinRequest
	msgTypeRaftJoinResponse

	msgTypeLeaderEpochOffsetRequest
	msgTypeLeaderEpochOffsetResponse

	msgTypePropagatedRequest
	msgTypePropagatedResponse

	msgTypeServerInfoRequest
	msgTypeServerInfoResponse

	msgTypePartitionStatusRequest
	msgTypePartitionStatusResponse

	msgTypePartitionNotification
)

const (
	// envelopeProtoV0 is version 0 of the envelope protocol.
	envelopeProtoV0 = 0x00

	// envelopeMinHeaderLen is the minimum length of the envelope header, i.e.
	// without CRC-32C set.
	envelopeMinHeaderLen = 8
)

var (
	// Encoding is the byte order to use for protocol serialization.
	Encoding = binary.BigEndian

	// envelopeMagicNumber is a value that indicates if a NATS message is a
	// structured message protobuf. This was chosen by random but deliberately
	// restricted to invalid UTF-8 to reduce the chance of a collision. This
	// was also verified to not match known file signatures.
	envelopeMagicNumber    = []byte{0xB9, 0x0E, 0x43, 0xB4}
	envelopeMagicNumberLen = len(envelopeMagicNumber)

	crc32cTable = crc32.MakeTable(crc32.Castagnoli)
)

// MarshalPublish serializes a protobuf publish message into the Liftbridge
// envelope wire format.
func MarshalPublish(msg *client.Message) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// MarshalAck serializes a protobuf ack message into the Liftbridge envelope
// wire format.
func MarshalAck(ack *client.Ack) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// MarshalServerInfoRequest serializes a ServerInfoRequest protobuf into the
// Liftbridge envelope wire format.
func MarshalServerInfoRequest(req *ServerInfoRequest) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// MarshalServerInfoResponse serializes a ServerInfoResponse protobuf into the
// Liftbridge envelope wire format.
func MarshalServerInfoResponse(req *ServerInfoResponse) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// MarshalPropagatedRequest serializes a PropagatedRequest protobuf into the
// Liftbridge envelope wire format.
func MarshalPropagatedRequest(req *PropagatedRequest) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// MarshalPropagatedResponse serializes a PropagatedResponse protobuf into the
// Liftbridge envelope wire format.
func MarshalPropagatedResponse(req *PropagatedResponse) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// MarshalPartitionStatusRequest serializes a PartitionStatusRequest protobuf
// into the Liftbridge envelope wire format.
func MarshalPartitionStatusRequest(req *PartitionStatusRequest) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// MarshalPartitionStatusResponse serializes a PartitionStatusResponse protobuf
// into the Liftbridge envelope wire format.
func MarshalPartitionStatusResponse(resp *PartitionStatusResponse) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// MarshalReplicationRequest serializes a ReplicationRequest protobuf into the
// Liftbridge envelope wire format.
func MarshalReplicationRequest(req *ReplicationRequest) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// MarshalLeaderEpochOffsetRequest serializes a LeaderEpochOffsetRequest
// protobuf into the Liftbridge envelope wire format.
func MarshalLeaderEpochOffsetRequest(req *LeaderEpochOffsetRequest) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// MarshalLeaderEpochOffsetResponse serializes a LeaderEpochOffsetResponse
// protobuf into the Liftbridge envelope wire format.
func MarshalLeaderEpochOffsetResponse(req *LeaderEpochOffsetResponse) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// MarshalPartitionNotification serializes a PartitionNotification protobuf
// into the Liftbridge envelope wire format.
func MarshalPartitionNotification(req *PartitionNotification) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// MarshalRaftJoinRequest serializes a RaftJoinRequest protobuf into the
// Liftbridge envelope wire format.
func MarshalRaftJoinRequest(req *RaftJoinRequest) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// MarshalRaftJoinResponse serializes a RaftJoinResponse protobuf into the
// Liftbridge envelope wire format.
func MarshalRaftJoinResponse(req *RaftJoinResponse) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// WriteReplicationResponseHeader writes the envelope protocol header for
// replication messages to the buffer and returns the number of bytes written.
func WriteReplicationResponseHeader(buf *bytes.Buffer) int { _ = "STUB: not implemented"; return 0 }

// marshalEnvelope serializes a protobuf message into the Liftbridge envelope
// wire format.
func marshalEnvelope(msg pb.Message, msgType msgType) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Version

// HeaderLen

// Flags

// MsgType

// UnmarshalPublish deserializes a Liftbridge publish envelope into a protobuf
// message.
func UnmarshalPublish(data []byte) (*client.Message, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// UnmarshalAck deserializes a Liftbridge ack envelope into a protobuf message.
func UnmarshalAck(data []byte) (*client.Ack, error) { _ = "STUB: not implemented"; return nil, nil }

// UnmarshalPropagatedRequest deserializes a Liftbridge PropagatedRequest
// envelope into a protobuf message.
func UnmarshalPropagatedRequest(data []byte) (*PropagatedRequest, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// UnmarshalPropagatedResponse deserializes a Liftbridge PropagatedResponse
// envelope into a protobuf message.
func UnmarshalPropagatedResponse(data []byte) (*PropagatedResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// UnmarshalServerInfoRequest deserializes a Liftbridge ServerInfoRequest
// envelope into a protobuf message.
func UnmarshalServerInfoRequest(data []byte) (*ServerInfoRequest, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// UnmarshalServerInfoResponse deserializes a Liftbridge ServerInfoResponse
// envelope into a protobuf message.
func UnmarshalServerInfoResponse(data []byte) (*ServerInfoResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// UnmarshalPartitionStatusRequest deserializes a Liftbridge
// PartitionStatusRequest envelope into a protobuf message.
func UnmarshalPartitionStatusRequest(data []byte) (*PartitionStatusRequest, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// UnmarshalPartitionStatusResponse deserializes a Liftbridge
// PartitionStatusResponse envelope into a protobuf message.
func UnmarshalPartitionStatusResponse(data []byte) (*PartitionStatusResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// UnmarshalRaftJoinRequest deserializes a Liftbridge RaftJoinRequest envelope
// into a protobuf message.
func UnmarshalRaftJoinRequest(data []byte) (*RaftJoinRequest, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// UnmarshalRaftJoinResponse deserializes a Liftbridge RaftJoinResponse
// envelope into a protobuf message.
func UnmarshalRaftJoinResponse(data []byte) (*RaftJoinResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// UnmarshalPartitionNotification deserializes a Liftbridge
// PartitionNotification envelope into a protobuf message.
func UnmarshalPartitionNotification(data []byte) (*PartitionNotification, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// UnmarshalLeaderEpochOffsetRequest deserializes a Liftbridge
// LeaderEpochOffsetRequest envelope into a protobuf message.
func UnmarshalLeaderEpochOffsetRequest(data []byte) (*LeaderEpochOffsetRequest, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// UnmarshalLeaderEpochOffsetResponse deserializes a Liftbridge
// LeaderEpochOffsetResponse envelope into a protobuf message.
func UnmarshalLeaderEpochOffsetResponse(data []byte) (*LeaderEpochOffsetResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// UnmarshalReplicationRequest deserializes a Liftbridge ReplicationRequest
// envelope into a protobuf message.
func UnmarshalReplicationRequest(data []byte) (*ReplicationRequest, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// UnmarshalReplicationResponse deserializes a Liftbridge replication response
// envelope and returns the leader epoch, HW, and message data.
func UnmarshalReplicationResponse(data []byte) (uint64, int64, []byte, error) {
	_ = "STUB: not implemented"
	return 0, 0, nil, nil
}

// We should have at least 16 bytes, 8 for leader epoch and 8 for HW.

// unmarshalEnvelope deserializes a Liftbridge envelope into a protobuf
// message.
func unmarshalEnvelope(data []byte, msg pb.Message, msgType msgType) error {
	_ = "STUB: not implemented"
	return nil
}

func checkEnvelope(data []byte, expectedType msgType) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Check CRC.

// Make sure there is a CRC present.

// hasBit checks if the given bit position is set on the provided byte.
func hasBit(n byte, pos uint8) bool { _ = "STUB: not implemented"; return false }
