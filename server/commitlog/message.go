package commitlog

import (
	"hash/crc32"

	client "github.com/liftbridge-io/liftbridge-api/v2/go"
)

var crc32cTable = crc32.MakeTable(crc32.Castagnoli)

// Message is the object that gets serialized and written to the log.
type Message struct {
	Crc        int32
	MagicByte  int8
	Attributes int8
	Key        []byte
	Value      []byte
	Headers    map[string][]byte

	// Transient fields
	Timestamp     int64
	LeaderEpoch   uint64
	AckInbox      string
	CorrelationID string
	AckPolicy     client.AckPolicy
	Offset        int64
}

// Encode the Message into the packetEncoder.
func (m *Message) Encode(e packetEncoder) error { _ = "STUB: not implemented"; return nil }

// crcField is used to perform a CRC32 check on a message.
type crcField struct {
	StartOffset int
}

// SaveOffset sets the position to fill the CRC digest.
func (f *crcField) SaveOffset(in int) {
	_ = "STUB: not implemented"

	// ReserveSize sets the number of bytes to reserve for the CRC digest.
	return
}

func (f *crcField) ReserveSize() int {
	_ = "STUB: not implemented"

	// Fill sets the CRC digest.
	return 0
}

func (f *crcField) Fill(curOffset int, buf []byte) error { _ = "STUB: not implemented"; return nil }

// Check the CRC digest.
func (f *crcField) Check(curOffset int, buf []byte) error { _ = "STUB: not implemented"; return nil }

// SerializedMessage is a serialized message read from the log.
type SerializedMessage []byte

// Crc returns the CRC32 digest of the message.
func (m SerializedMessage) Crc() uint32 { _ = "STUB: not implemented"; return 0 }

// MagicByte returns the byte used for encoding protocol version detection.
func (m SerializedMessage) MagicByte() int8 {
	_ = "STUB: not implemented"

	// Attributes returns the byte used for message flags.
	return 0
}

func (m SerializedMessage) Attributes() int8 {
	_ = "STUB: not implemented"

	// Key returns the message key.
	return 0
}

func (m SerializedMessage) Key() []byte { _ = "STUB: not implemented"; return nil }

// Value returns the message value.
func (m SerializedMessage) Value() []byte { _ = "STUB: not implemented"; return nil }

// Headers returns the message headers map.
func (m SerializedMessage) Headers() map[string][]byte { _ = "STUB: not implemented"; return nil }

func (m SerializedMessage) keyOffsets() (start, end, size int32) {
	_ = "STUB: not implemented"
	return 0, 0, 0
}

func (m SerializedMessage) valueOffsets() (start, end, size int32) {
	_ = "STUB: not implemented"
	return 0, 0, 0
}
