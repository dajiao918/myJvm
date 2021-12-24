package stack

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
)

type SWAP struct {
	base.NoOperandsInstruction
}

// Execute //
//交换栈顶的两个变量
func (self *SWAP) Execute(frame *rtda.Frame) {
	slot1 := frame.OperandStack().PopSlot()
	slot2 := frame.OperandStack().PopSlot()
	frame.OperandStack().PushSlot(slot1)
	frame.OperandStack().PushSlot(slot2)
}
