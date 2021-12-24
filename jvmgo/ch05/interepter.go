package main

import (
	"fmt"
	"jvmgo/ch03/classFile"
	"jvmgo/ch05/instructions"
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
)

func interpret(method *classFile.MemberInfo) {
	codeAttr := method.CodeAttribute()
	maxLocals := codeAttr.MaxLocals()
	maxStack := codeAttr.MaxStack()
	byteCode := codeAttr.Code()
	fmt.Printf(" ===========\n %v \n============\n", byteCode)
	thread := rtda.NewThread()
	frame := thread.NewFrame(maxLocals, maxStack)
	thread.PushFrame(frame)

	defer catchErr(frame)
	loop(thread, byteCode)
}

func loop(thread *rtda.Thread, code []byte) {
	frame := thread.PopFrame()
	reader := &base.BytecodeReader{}
	for {
		pc := frame.NextPC()
		thread.SetPC(pc)

		reader.Reset(pc, code)
		opcode := reader.ReadUint8()
		inst := instructions.NewInstruction(opcode)
		inst.FetchOperands(reader)
		frame.SetNextPC(reader.PC())

		fmt.Printf("pc %2d inst %T %v\n", pc, inst, inst)
		inst.Execute(frame)
	}
}

func catchErr(frame *rtda.Frame) {
	if r := recover(); r != nil {
		fmt.Printf("LocalVals %v \n", frame.Localvars())
		fmt.Printf("OperandStack %v \n", frame.OperandStack())
		panic(r)
	}
}
