package stack

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
)

type DUP struct{ base.NoOperandsInstruction }
type DUP_x1 struct{ base.NoOperandsInstruction }
type DUP_x2 struct{ base.NoOperandsInstruction }
type DUP2 struct{ base.NoOperandsInstruction }
type DUP2_x1 struct{ base.NoOperandsInstruction }
type DUP2_x2 struct{ base.NoOperandsInstruction }

// Execute /*复制操作数栈顶的变量*/
func (self *DUP) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	slot := stack.PopSlot()
	stack.PushSlot(slot)
	stack.PushSlot(slot)
}

// Execute /*
//  1-2-3
//      |
//     /
//    V
//  1-3-2-3      */
func (self *DUP_x1) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	stack.PushSlot(slot1)
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
}

// Execute /*
//  1-2-3-4
//        |
//       /
//      /
//     V
//   1-4-3-2-4      */
func (self *DUP_x2) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	slot3 := stack.PopSlot()
	stack.PushSlot(slot1)
	stack.PushSlot(slot3)
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
}

// Execute /*
//  1-2-3-4
//      | |
//       \ \
//        \ \
//         V V
// 1-2-3-4-3-4      */
func (self *DUP2) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
}

// Execute /*
// 1-1-2-3-4
//       | |
//      /  /
//     /  /
//     V V
// 1-1-3-4-2-3-4     */
func (self *DUP2_x1) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	slot3 := stack.PopSlot()
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
	stack.PushSlot(slot3)
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
}

// Execute /*
// 1-1-2-3-4
//      | |
//     / /
//    / /
//   V V
// 1-3-4-1-2-3-4     */
func (self *DUP2_x2) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	slot3 := stack.PopSlot()
	slot4 := stack.PopSlot()
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
	stack.PushSlot(slot4)
	stack.PushSlot(slot3)
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
}
