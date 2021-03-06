package stores

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
)

// FSTORE /*将操作数栈顶的数字存入到局部变量表的指定索引中*/
type FSTORE struct{ base.Index8Instruction }
type FSTORE_0 struct{ base.NoOperandsInstruction }
type FSTORE_1 struct{ base.NoOperandsInstruction }
type FSTORE_2 struct{ base.NoOperandsInstruction }
type FSTORE_3 struct{ base.NoOperandsInstruction }

func _fstore(frame *rtda.Frame, index uint) {
	val := frame.OperandStack().PopFloat()
	frame.Localvars().SetFloat(index, val)
}

func (self *FSTORE) Execute(frame *rtda.Frame) {
	_fstore(frame, self.Index)
}

func (self *FSTORE_0) Execute(frame *rtda.Frame) {
	_fstore(frame, 0)
}

func (self *FSTORE_1) Execute(frame *rtda.Frame) {
	_fstore(frame, 1)
}

func (self *FSTORE_2) Execute(frame *rtda.Frame) {
	_fstore(frame, 2)
}

func (self *FSTORE_3) Execute(frame *rtda.Frame) {
	_fstore(frame, 3)
}
