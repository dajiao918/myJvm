package base

import (
	"testing"
)

func TestByteReader(t *testing.T) {
	bytes := make([]byte, 10)
	bytes[0] = 1
	bytes[1] = 1
	bytes[2] = 1
	bytes[3] = 1
	reader := newBytecodeReader(0, bytes)
	//println(reader.ReadInt8())
	//println(reader.ReadInt16())
	println(reader.ReadInt32())
	println((1 << 24) | (1 << 16) | (1 << 8) | 1)
}
