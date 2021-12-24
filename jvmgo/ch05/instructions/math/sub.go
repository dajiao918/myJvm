package math

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
)

type DSUB struct {
	base.NoOperandsInstruction
}

// Execute pop two double number from operand_stack, do sub
func (self *DSUB) Execute(frame *rtda.Frame) {
	d1 := frame.OperandStack().PopDouble()
	d2 := frame.OperandStack().PopDouble()
	result := d2 - d1
	frame.OperandStack().PushDouble(result)
}

type FSUB struct {
	base.NoOperandsInstruction
}

// Execute pop two float number from operand_stack, do sub
func (self *FSUB) Execute(frame *rtda.Frame) {
	d1 := frame.OperandStack().PopFloat()
	d2 := frame.OperandStack().PopFloat()
	result := d2 - d1
	frame.OperandStack().PushFloat(result)
}

type ISUB struct {
	base.NoOperandsInstruction
}

// Execute pop two double number from operand_stack, do sub
func (self *ISUB) Execute(frame *rtda.Frame) {
	d1 := frame.OperandStack().PopInt()
	d2 := frame.OperandStack().PopInt()
	result := d2 - d1
	frame.OperandStack().PushInt(result)
}

type LSUB struct {
	base.NoOperandsInstruction
}

// Execute pop two long number from operand_stack, do sub
func (self *LSUB) Execute(frame *rtda.Frame) {
	d1 := frame.OperandStack().PopLong()
	d2 := frame.OperandStack().PopLong()
	result := d2 - d1
	frame.OperandStack().PushLong(result)
}
