package rtda

/*Frame 栈帧:包括操作数栈，和局部变量表*/
type Frame struct {
	operandStack *OperandStack
	localVars    LocalVars
	lower        *Frame
}

func newFrame(maxLocals, maxStack uint) *Frame {
	return &Frame{
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
