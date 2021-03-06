package loads

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
)

// ALOAD /*load 根据索引，将局部变量表的变量到操作数栈*/
type ALOAD struct{ base.Index8Instruction }
type ALOAD_0 struct{ base.NoOperandsInstruction }
type ALOAD_1 struct{ base.NoOperandsInstruction }
type ALOAD_2 struct{ base.NoOperandsInstruction }
type ALOAD_3 struct{ base.NoOperandsInstruction }

func _aload(frame *rtda.Frame, index uint) {
	val := frame.Localvars().GetRef(index)
	frame.OperandStack().PushRef(val)
}

func (self *ALOAD) Execute(frame *rtda.Frame) {
	_aload(frame, self.Index)
}

func (self *ALOAD_0) Execute(frame *rtda.Frame) {
	_aload(frame, 0)
}
func (self *ALOAD_1) Execute(frame *rtda.Frame) {
	_aload(frame, 1)
}
func (self *ALOAD_2) Execute(frame *rtda.Frame) {
	_aload(frame, 2)
}
func (self *ALOAD_3) Execute(frame *rtda.Frame) {
	_aload(frame, 3)
}
