package loads

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
)

// ILOAD /*load 根据索引，将局部变量表的变量到操作数栈*/
type LLOAD struct{ base.Index8Instruction }
type LLOAD_0 struct{ base.NoOperandsInstruction }
type LLOAD_1 struct{ base.NoOperandsInstruction }
type LLOAD_2 struct{ base.NoOperandsInstruction }
type LLOAD_3 struct{ base.NoOperandsInstruction }

func _lload(frame *rtda.Frame, index uint) {
	val := frame.Localvars().GetLong(index)
	frame.OperandStack().PushLong(val)
}

func (self *LLOAD) Execute(frame *rtda.Frame) {
	_lload(frame, self.Index)
}

func (self *LLOAD_0) Execute(frame *rtda.Frame) {
	_lload(frame, 0)
}
func (self *LLOAD_1) Execute(frame *rtda.Frame) {
	_lload(frame, 1)
}
func (self *LLOAD_2) Execute(frame *rtda.Frame) {
	_lload(frame, 2)
}
func (self *LLOAD_3) Execute(frame *rtda.Frame) {
	_lload(frame, 3)
}
