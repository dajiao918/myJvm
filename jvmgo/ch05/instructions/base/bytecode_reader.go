package base

type BytecodeReader struct {
	code []byte
	pc   int
}

func newBytecodeReader(pc int, code []byte) *BytecodeReader {
	return &BytecodeReader{pc: pc, code: code}
}

func (r *BytecodeReader) ReadUint8() uint8 {
	res := r.code[r.pc]
	r.pc++
	return res
}

func (r *BytecodeReader) ReadInt8() int8 {
	return int8(r.ReadUint8())
}

func (r *BytecodeReader) ReadUint16() uint16 {
	low := uint16(r.ReadUint8())
	high := uint16(r.ReadUint8())
	return high<<8 | low
}

func (r *BytecodeReader) ReadInt16() int16 {
	return int16(r.ReadUint16())
}

func (r *BytecodeReader) ReadInt32() int32 {
	low := int32(r.ReadInt16())
	high := int32(r.ReadInt16())
	return high<<16 | low
}
