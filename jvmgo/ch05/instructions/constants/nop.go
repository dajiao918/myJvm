package constants

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
)

type Nop struct {
	base.NoOperandsInstruction
}

// Execute /*nop 指令什么也不做*/
func (self *Nop) Execute(frame *rtda.Frame) {}
