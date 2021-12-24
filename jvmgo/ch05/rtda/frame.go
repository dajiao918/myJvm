package rtda

/*Frame 栈帧:包括操作数栈，和局部变量表*/
type Frame struct {
	operandStack *OperandStack
	localVars    LocalVars
	thread       *Thread
	lower        *Frame
	nextPC       int
}

func newFrame(thread *Thread, maxLocals, maxStack uint16) *Frame {
	return &Frame{
		thread:       thread,
		operandStack: newOperandStack(maxStack),
		localVars:    newLocalVars(maxLocals),
	}
}

func (self *Frame) Localvars() LocalVars {
	return self.localVars
}

func (self *Frame) OperandStack() *OperandStack {
	return self.operandStack
}

func (self *Frame) SetNextPC(pc int) {
	self.nextPC = pc
}

func (self *Frame) NextPC() int {
	return self.nextPC
}

func (self *Frame) Thread() *Thread {
	return self.thread
}
