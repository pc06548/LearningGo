package conv

// This package includes functions to convert go data types to & from byte array.
// All the functions will panic on unsuccessful conversion.
// This is only intended to use when source of byte array is confirmed.
// Because of this, no error handling is implemented.

import (
	"math"
	"encoding/binary"
)

func StringToByte(v string) []byte {
	return []byte(v)
}

func ByteToString(b []byte) string {
	return string(b)
}

func Int32ToByte(v int32) []byte {
	buf := make([]byte, 4)
	binary.LittleEndian.PutUint32(buf, uint32(v))
	return buf
}

func ByteToInt32(v []byte) int32 {
	return int32(binary.LittleEndian.Uint32(v))
}

func Int64ToByte(v int64) []byte {
	buf := make([]byte, 8)
	binary.LittleEndian.PutUint64(buf, uint64(v))
	return buf
}

func ByteToInt64(v []byte) int64 {
	return int64(binary.LittleEndian.Uint64(v))
}

func Uint32ToByte(v uint32) []byte {
	buf := make([]byte, 4)
	binary.LittleEndian.PutUint32(buf, v)
	return buf
}

func ByteToUint32(v []byte) uint32 {
	return binary.LittleEndian.Uint32(v)
}

func Uint64ToByte(v uint64) []byte {
	buf := make([]byte, 8)
	binary.LittleEndian.PutUint64(buf, v)
	return buf
}

func ByteToUint64(v []byte) uint64 {
	return binary.LittleEndian.Uint64(v)
}

func Float64ToByte(v float64) []byte {
	buf := make([]byte, binary.MaxVarintLen64)
	binary.LittleEndian.PutUint64(buf[:], math.Float64bits(v))
	return buf
}

func ByteToFloat64(v []byte) float64 {
	return math.Float64frombits(binary.LittleEndian.Uint64(v))
}
