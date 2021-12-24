package control

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
)

type GOTO struct {
	base.BranchInstruction
	offset int
}

func (self *GOTO) Execute(frame *rtda.Frame) {
	base.Branch(frame, self.Offset)
}
