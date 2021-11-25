package constants

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
)

// BIPUSH /*将byte转换为int，并压入操作数栈*/
type BIPUSH struct {
	val int8
}

// SIPUSH /*将short转换为int，并压入操作数栈*/
type SIPUSH struct {
	val int16
}

func (self *BIPUSH) FetchOperands(reader *base.BytecodeReader) {
	self.val = reader.ReadInt8()
}

func (self *BIPUSH) Execute(frame *rtda.Frame) {
	tem := int32(self.val)
	frame.OperandStack().PushInt(tem)
}

func (self *SIPUSH) FetchOperands(reader *base.BytecodeReader) {
	self.val = reader.ReadInt16()
}

func (self *SIPUSH) Execute(frame *rtda.Frame) {
	tem := int32(self.val)
	frame.OperandStack().PushInt(tem)
}
