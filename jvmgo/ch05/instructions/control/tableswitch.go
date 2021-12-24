package control

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
)

// TABLESWITCH
//类似于 switch
//		case 1: case 2: case 3:
//这样顺序排序的，
type TABLESWITCH struct {
	defaultOffset int32
	low           int32
	high          int32
	jumpOffsets   []int32
}

func (self *TABLESWITCH) FetchOperands(reader *base.BytecodeReader) {
	reader.SkipPadding()
	//从code属性中解析switch语句
	self.defaultOffset = reader.ReadInt32()
	self.low = reader.ReadInt32()
	self.high = reader.ReadInt32()
	jumpOffsets := self.high - self.low + 1
	self.jumpOffsets = reader.ReadTable32(jumpOffsets)
}

func (self *TABLESWITCH) Execute(frame *rtda.Frame) {
	//注意，从操作数栈中弹出switch语句的索引
	index := frame.OperandStack().PopInt()
	var offset int
	if index >= self.low && index <= self.high {
		offset = int(self.jumpOffsets[index-self.low])
	} else {
		offset = int(self.defaultOffset)
	}
	base.Branch(frame, offset)
}
