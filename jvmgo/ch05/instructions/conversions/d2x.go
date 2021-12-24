package conversions

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
)

type D2F struct {
	base.NoOperandsInstruction
}

func (self *D2F) Execute(frame *rtda.Frame) {
	d := frame.OperandStack().PopDouble()
	f := float32(d)
	frame.OperandStack().PushFloat(f)
}

type D2I struct {
	base.NoOperandsInstruction
}

func (self *D2I) Execute(frame *rtda.Frame) {
	d := frame.OperandStack().PopDouble()
	i := int32(d)
	frame.OperandStack().PushInt(i)
}

type D2L struct {
	base.NoOperandsInstruction
}

func (self *D2L) Execute(frame *rtda.Frame) {
	d := frame.OperandStack().PopDouble()
	l := int64(d)
	frame.OperandStack().PushLong(l)
}
