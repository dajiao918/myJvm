package control

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
)

// LookUpSwitch
// 类似于switch
// case 10: case 2: case 5  这种乱序的case
type LookUpSwitch struct {
	defaultOffsets int32
	npairs         int32
	matchOffsets   []int32
}

func (self *LookUpSwitch) FetchOperands(reader *base.BytecodeReader) {
	reader.SkipPadding()
	self.defaultOffsets = reader.ReadInt32()
	self.npairs = reader.ReadInt32()
	self.matchOffsets = reader.ReadTable32(self.npairs * 2)
}

func (self *LookUpSwitch) Execute(frame *rtda.Frame) {
	index := frame.OperandStack().PopInt()

	var offset int
	for i := 0; i < len(self.matchOffsets); i += 2 {
		//前一个是case的值索引
		if index == self.matchOffsets[i] {
			//后一个是case执行的指令索引
			offset = int(self.matchOffsets[i+1])
			base.Branch(frame, offset)
			return
		}
	}
	//都不满足，跳转default
	offset = int(self.defaultOffsets)
	base.Branch(frame, offset)
}
