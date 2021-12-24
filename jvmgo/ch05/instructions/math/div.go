package math

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
)

type DDIV struct {
	base.NoOperandsInstruction
}

// Execute pop two double number from operand_stack, do div
func (self *DDIV) Execute(frame *rtda.Frame) {
	d1 := frame.OperandStack().PopDouble()
	d2 := frame.OperandStack().PopDouble()
	if d1 == 0 {
		panic("java.lang.ArithmeticException: / by zero")
	}
	result := d2 / d1
	frame.OperandStack().PushDouble(result)
}

type FDIV struct {
	base.NoOperandsInstruction
}

// Execute pop two float number from operand_stack, do div
func (self *FDIV) Execute(frame *rtda.Frame) {
	d1 := frame.OperandStack().PopFloat()
	d2 := frame.OperandStack().PopFloat()
	if d1 == 0 {
		panic("java.lang.ArithmeticException: / by zero")
	}
	result := d2 / d1
	frame.OperandStack().PushFloat(result)
}

type IDIV struct {
	base.NoOperandsInstruction
}

// Execute pop two double number from operand_stack, do div
func (self *IDIV) Execute(frame *rtda.Frame) {
	d1 := frame.OperandStack().PopInt()
	d2 := frame.OperandStack().PopInt()
	if d1 == 0 {
		panic("java.lang.ArithmeticException: / by zero")
	}
	result := d2 / d1
	frame.OperandStack().PushInt(result)
}

type LDIV struct {
	base.NoOperandsInstruction
}

// Execute pop two long number from operand_stack, do div
func (self *LDIV) Execute(frame *rtda.Frame) {
	d1 := frame.OperandStack().PopLong()
	d2 := frame.OperandStack().PopLong()
	if d1 == 0 {
		panic("java.lang.ArithmeticException: / by zero")
	}
	result := d2 / d1
	frame.OperandStack().PushLong(result)
}
