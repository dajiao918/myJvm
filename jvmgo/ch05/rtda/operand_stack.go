package rtda

import "math"

type OperandStack struct {
	size  uint
	slots []Slot
}

func newOperandStack(maxStack uint) *OperandStack {
	if maxStack > 0 {
		return &OperandStack{
			slots: make([]Slot, maxStack),
		}
	}
	return nil
}

func (self *OperandStack) PushInt(val int32) {
	self.slots[self.size].num = val
	self.size++
}

func (self *OperandStack) PopInt() int32 {
	self.size--
	return self.slots[self.size].num
}

func (self *OperandStack) PushFloat(val float32) {
	bits := math.Float32bits(val)
	self.PushInt(int32(bits))
}

func (self *OperandStack) PopFloat() float32 {
	val := self.PopInt()
	return math.Float32frombits(uint32(val))
}

func (self *OperandStack) PushLong(val int64) {
	self.PushInt(int32(val))
	self.PushInt(int32(val >> 32))
}

func (self *OperandStack) PopLong() int64 {
	high := uint32(self.PopInt())
	low := uint32(self.PopInt())
	return int64(high)<<32 | int64(low)
}

func (self *OperandStack) PushDouble(val float64) {
	bits := math.Float64bits(val)
	self.PushLong(int64(bits))
}

func (self *OperandStack) PopDouble() float64 {
	val := self.PopLong()
	return math.Float64frombits(uint64(val))
}

func (self *OperandStack) PushRef(ref *Object) {
	self.slots[self.size].ref = ref
	self.size++
}

func (self *OperandStack) PopRef() *Object {
	self.size--
	return self.slots[self.size].ref
}

func (self *OperandStack) PushSlot(slot Slot) {
	self.slots[self.size] = slot
	self.size++
}

func (self *OperandStack) PopSlot() Slot {
	self.size--
	return self.slots[self.size]
}
