package math

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
)

type DMUL struct {
	base.NoOperandsInstruction
}

// Execute pop two double number from operand_stack, do mul
func (self *DMUL) Execute(frame *rtda.Frame) {
	d1 := frame.OperandStack().PopDouble()
	d2 := frame.OperandStack().PopDouble()
	result := d1 * d2
	frame.OperandStack().PushDouble(result)
}

type FMUL struct {
	base.NoOperandsInstruction
}

// Execute pop two float number from operand_stack, do mul
func (self *FMUL) Execute(frame *rtda.Frame) {
	d1 := frame.OperandStack().PopFloat()
	d2 := frame.OperandStack().PopFloat()
	result := d1 * d2
	frame.OperandStack().PushFloat(result)
}

type IMUL struct {
	base.NoOperandsInstruction
}

// Execute pop two int number from operand_stack, do mul
func (self *IMUL) Execute(frame *rtda.Frame) {
	d1 := frame.OperandStack().PopInt()
	d2 := frame.OperandStack().PopInt()
	result := d1 * d2
	frame.OperandStack().PushInt(result)
}

type LMUL struct {
	base.NoOperandsInstruction
}

// Execute pop two long number from operand_stack, do mul
func (self *LMUL) Execute(frame *rtda.Frame) {
	d1 := frame.OperandStack().PopLong()
	d2 := frame.OperandStack().PopLong()
	result := d1 * d2
	frame.OperandStack().PushLong(result)
}
