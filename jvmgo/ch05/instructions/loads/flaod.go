package loads

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
)

// FLOAD /*load 根据索引，将局部变量表的变量到操作数栈*/
type FLOAD struct{ base.Index8Instruction }
type FLOAD_0 struct{ base.NoOperandsInstruction }
type FLOAD_1 struct{ base.NoOperandsInstruction }
type FLOAD_2 struct{ base.NoOperandsInstruction }
type FLOAD_3 struct{ base.NoOperandsInstruction }

func _fload(frame *rtda.Frame, index uint) {
	val := frame.Localvars().GetFloat(index)
	frame.OperandStack().PushFloat(val)
}

func (self *FLOAD) Execute(frame *rtda.Frame) {
	_fload(frame, self.Index)
}

func (self *FLOAD_0) Execute(frame *rtda.Frame) {
	_fload(frame, 0)
}
func (self *FLOAD_1) Execute(frame *rtda.Frame) {
	_fload(frame, 1)
}
func (self *FLOAD_2) Execute(frame *rtda.Frame) {
	_fload(frame, 2)
}
func (self *FLOAD_3) Execute(frame *rtda.Frame) {
	_fload(frame, 3)
}
