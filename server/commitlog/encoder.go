package commitlog

import (
	"encoding/binary"
	"errors"
)

var (
	// encoding is the byte order to use for internal disk serialization.
	encoding = binary.BigEndian

	errInvalidStringLength    = errors.New("invalid string length")
	errInvalidArrayLength     = errors.New("invalid array length")
	errInvalidByteSliceLength = errors.New("invalid byteslice length")
)

// packetEncoder is used to serialize an object.
type packetEncoder interface {
	PutBool(in bool)
	PutInt8(in int8)
	PutInt16(in int16)
	PutInt32(in int32)
	PutInt64(in int64)
	PutArrayLength(in int) error
	PutRawBytes(in []byte) error
	PutBytes(in []byte) error
	PutString(in string) error
	PutNullableString(in *string) error
	PutStringArray(in []string) error
	PutInt32Array(in []int32) error
	PutInt64Array(in []int64) error
	Push(pe pushEncoder)
	Pop()
}

// pushEncoder is used to push an operation onto the stack to perform later
// once serialized bytes are filled.
type pushEncoder interface {
	SaveOffset(in int)
	ReserveSize() int
	Fill(curOffset int, buf []byte) error
}

// encoder is a struct that can be serialized.
type encoder interface {
	Encode(e packetEncoder) error
}

// encode serializes the struct to bytes.
func encode(e encoder) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// lenEncoder is a packetEncoder that tracks the running length of serialized
// bytes.
type lenEncoder struct {
	Length int
}

// PutBool increments length for a bool.
func (e *lenEncoder) PutBool(in bool) {
	_ = "STUB: not implemented"

	// PutInt8 increments length for an int8.
	return
}

func (e *lenEncoder) PutInt8(in int8) {
	_ = "STUB: not implemented"

	// PutInt16 increments length for an int16.
	return
}

func (e *lenEncoder) PutInt16(in int16) {
	_ = "STUB: not implemented"

	// PutInt32 increments length for an int32.
	return
}

func (e *lenEncoder) PutInt32(in int32) {
	_ = "STUB: not implemented"

	// PutInt64 increments length for an int64.
	return
}

func (e *lenEncoder) PutInt64(in int64) {
	_ = "STUB: not implemented"

	// PutArrayLength increments length for an array size.
	return
}

func (e *lenEncoder) PutArrayLength(in int) error { _ = "STUB: not implemented"; return nil }

// arrays

// PutBytes increments length for a size-prefixed byte array.
func (e *lenEncoder) PutBytes(in []byte) error { _ = "STUB: not implemented"; return nil }

// PutRawBytes increments length for a raw byte array.
func (e *lenEncoder) PutRawBytes(in []byte) error { _ = "STUB: not implemented"; return nil }

// PutString increments length for a string.
func (e *lenEncoder) PutString(in string) error { _ = "STUB: not implemented"; return nil }

// PutNullableString increments length for a nullable string.
func (e *lenEncoder) PutNullableString(in *string) error { _ = "STUB: not implemented"; return nil }

// PutStringArray increments length for a string array.
func (e *lenEncoder) PutStringArray(in []string) error { _ = "STUB: not implemented"; return nil }

// PutInt32Array increments length for an int32 array.
func (e *lenEncoder) PutInt32Array(in []int32) error { _ = "STUB: not implemented"; return nil }

// PutInt64Array increments length for an int64 array.
func (e *lenEncoder) PutInt64Array(in []int64) error { _ = "STUB: not implemented"; return nil }

// Push increments length based on the pushEncoder's reserved size.
func (e *lenEncoder) Push(pe pushEncoder) { _ = "STUB: not implemented"; return }

// Pop is a no-op.
func (e *lenEncoder) Pop() {
	_ = "STUB: not implemented"

	// byteEncoder is a packetEncoder that serializes data into a byte slice.
	return
}

type byteEncoder struct {
	b     []byte
	off   int
	stack []pushEncoder
}

// Bytes returns the underlying byte slice.
func (e *byteEncoder) Bytes() []byte {
	_ = "STUB: not implemented"

	// NewByteEncoder creates a new ByteEncoder with the given backing
	// pre-allocated byte slice.
	return nil
}

func newByteEncoder(b []byte) *byteEncoder { _ = "STUB: not implemented"; return nil }

// PutBool serializes a bool.
func (e *byteEncoder) PutBool(in bool) { _ = "STUB: not implemented"; return }

// PutInt8 serializes an int8.
func (e *byteEncoder) PutInt8(in int8) { _ = "STUB: not implemented"; return }

// PutInt16 serializes an int16.
func (e *byteEncoder) PutInt16(in int16) { _ = "STUB: not implemented"; return }

// PutInt32 serializes an int32.
func (e *byteEncoder) PutInt32(in int32) { _ = "STUB: not implemented"; return }

// PutInt64 serializes an int64.
func (e *byteEncoder) PutInt64(in int64) { _ = "STUB: not implemented"; return }

// PutArrayLength serializes an array length as an int32.
func (e *byteEncoder) PutArrayLength(in int) error { _ = "STUB: not implemented"; return nil }

// PutRawBytes serializes a byte slice.
func (e *byteEncoder) PutRawBytes(in []byte) error { _ = "STUB: not implemented"; return nil }

// PutBytes serializes a size-prefixed byte slice.
func (e *byteEncoder) PutBytes(in []byte) error { _ = "STUB: not implemented"; return nil }

// PutString serializes a size-prefixed string.
func (e *byteEncoder) PutString(in string) error { _ = "STUB: not implemented"; return nil }

// PutNullableString serializes a nullable string.
func (e *byteEncoder) PutNullableString(in *string) error { _ = "STUB: not implemented"; return nil }

// PutStringArray serializes a string array.
func (e *byteEncoder) PutStringArray(in []string) error { _ = "STUB: not implemented"; return nil }

// PutInt32Array serializes an int32 array.
func (e *byteEncoder) PutInt32Array(in []int32) error { _ = "STUB: not implemented"; return nil }

// PutInt64Array serializes an int64 array.
func (e *byteEncoder) PutInt64Array(in []int64) error { _ = "STUB: not implemented"; return nil }

// Push adds the given pushEncoder to the stack and saves the current offset
// position.
func (e *byteEncoder) Push(pe pushEncoder) { _ = "STUB: not implemented"; return }

// Pop the stack and run the popped pushEncoder on the serialized data.
func (e *byteEncoder) Pop() {
	_ = "STUB: not implemented"
	// this is go's ugly pop pattern (the inverse of append)
	return
}
