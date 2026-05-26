package commitlog

import (
	"context"
)

const (
	offsetPos       = 0
	timestampPos    = 8
	leaderEpochPos  = 16
	sizePos         = 24
	msgSetHeaderLen = 28
)

type messageSet []byte

func entriesForMessageSet(basePos int64, ms []byte) []*entry { _ = "STUB: not implemented"; return nil }

func newMessageSetFromProto(baseOffset, basePos int64, msgs []*Message, concurrencyControl bool) (
	messageSet, []*entry, error) {
	_ = "STUB: not implemented"

	// When concurrency control is enabled, messages shall be processed on by one
	return *new(messageSet), nil, nil
}

// Check expected offset for concurrency in case of Optimistic Concurrency Control

// readMessage reads a single message from the reader or blocks until one is
// available. It returns the Message in addition to its offset, timestamp, and
// leader epoch. This may return uncommitted messages if the reader was created
// with the uncommitted flag set to true.
func readMessage(ctx context.Context, reader contextReader, headersBuf []byte) (SerializedMessage, int64, int64, uint64, error) {
	_ = "STUB: not implemented"
	return *new(SerializedMessage), 0, 0, 0, nil
}

// Check the CRC on the message.

// If the CRC doesn't match, data on disk is corrupted which means the
// server is in an unrecoverable state.

func (ms messageSet) Offset() int64 { _ = "STUB: not implemented"; return 0 }

func (ms messageSet) Timestamp() int64 { _ = "STUB: not implemented"; return 0 }

func (ms messageSet) LeaderEpoch() uint64 { _ = "STUB: not implemented"; return 0 }

func (ms messageSet) Size() int32 { _ = "STUB: not implemented"; return 0 }

func (ms messageSet) Message() SerializedMessage {
	_ = "STUB: not implemented"
	return *new(SerializedMessage)
}
