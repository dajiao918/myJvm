package base

import "jvmgo/ch05/rtda"

type Instruction interface {
	FetchOperands(reader *BytecodeReader)
	Execute(frame *rtda.Frame)
}

/*NoOperandsInstruction 无操作码指令*/
type NoOperandsInstruction struct{}

/*FetchOperands 无操作码指令不用读取操作码*/
func (self *NoOperandsInstruction) FetchOperands(reader *BytecodeReader) {}

// BranchInstruction /*跳转指令*/
type BranchInstruction struct {
	offset int
}

// FetchOperands /*读取16字节的pc偏移*/
func (self *BranchInstruction) FetchOperands(reader *BytecodeReader) {
	self.offset = int(reader.ReadInt16())
}

// Index8Instruction /*加载类指令和存储类指令需要局部变量表的int8索引*/
type Index8Instruction struct {
	Index uint
}

func (self *Index8Instruction) FetchOperands(reader *BytecodeReader) {
	self.Index = uint(reader.ReadUint8())
}

type Index16Instruction struct {
	Index uint
}

func (self *Index16Instruction) FetchOperands(reader *BytecodeReader) {
	self.Index = uint(reader.ReadUint16())
}
