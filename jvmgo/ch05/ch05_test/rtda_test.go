package ch04_test

import (
	"fmt"
	"jvmgo/ch04/rtda"
	"jvmgo/ch05/instructions/base"
	"testing"
)

func TestLocalVars(t *testing.T) {
	frame := rtda.NewFrame(100, 100)
	localVars := frame.Localvars()
	localVars.SetInt(0, 100)
	localVars.SetInt(1, -10)
	localVars.SetFloat(2, 10.3)
	localVars.SetLong(3, 100022133166555)
	localVars.SetDouble(5, 100022133166555.230)
	localVars.SetRef(7, nil)
	fmt.Printf("%v\n", localVars.GetInt(0))
	fmt.Println(localVars.GetInt(1))
	fmt.Println(localVars.GetFloat(2))
	fmt.Println(localVars.GetLong(3))
	fmt.Println(localVars.GetDouble(5))
	fmt.Println(localVars.GetRef(7))
}

func TestOperandStack(t *testing.T) {
	frame := rtda.NewFrame(100, 100)
	operand := frame.OperandStack()
	operand.PushInt(10)
	operand.PushFloat(100.56)
	operand.PushLong(100155454645666)
	operand.PushDouble(100155454645666.235)
	operand.PushRef(nil)
	fmt.Println(operand.PopRef())
	fmt.Println(operand.PopDouble())
	fmt.Println(operand.PopLong())
	fmt.Println(operand.PopFloat())
	fmt.Println(operand.PopInt())
}

func TestShift(t *testing.T) {
	frame := rtda.NewFrame(100, 100)
	operand := frame.OperandStack()
	operand.PushInt(-4)
	operand.PushInt(1)
	is := &IUSHR{}
	is.Execute(frame)
	t.Log(frame.OperandStack().PopInt())
}

type IUSHR struct{ base.NoOperandsInstruction }

func (self *IUSHR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	s := uint32(v2) & 0x1f
	result := int32(v1 >> s)
	stack.PushInt(result)
}
