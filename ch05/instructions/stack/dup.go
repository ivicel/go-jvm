package stack

import (
	"go-jvm/ch05/instructions/base"
	"go-jvm/ch05/rtda"
)

// 复制栈顶, 新的数据再压入栈顶,
type DUP struct{ base.NoOperandsInstruction }

func (s *DUP) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	slot := stack.PopSlot()
	stack.PushSlot(slot)
	stack.PushSlot(slot)
}

// 复制栈顶, 新的数据再压入栈顶前两位之后的位置, 也就是取出两个槽位后,
// 复制第一个压入, 再压回这两个原数据
// 栈底->栈顶
// ....[c][b][a] -> ...[c][a][b][a]
type DUP_X1 struct{ base.NoOperandsInstruction }

func (s *DUP_X1) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	stack.PushSlot(slot1)
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
}

// 复制栈顶, 新的数据再压入栈顶前三位之后的位置 也就是取出三个槽位后,
// 复制第一个压入, 再压回这三个原数据
// 栈底->栈顶
// ....[d][c][b][a] -> ...[d][a][c][b][a]
type DUP_X2 struct{ base.NoOperandsInstruction }

func (s *DUP_X2) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	slot3 := stack.PopSlot()
	stack.PushSlot(slot1)
	stack.PushSlot(slot3)
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
}

// 复制栈顶的两个槽位, 再按这顺序入栈
type DUP2 struct{ base.NoOperandsInstruction }

func (s *DUP2) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
}

// 复制栈顶的两个槽位, 再把复制的新数据插入到第一位之后, 然后这三个弹出的位置依次压入
// 复制两个后压入, 再压回这三个原数据
// 栈底->栈顶
// ....[d][c][b][a] -> ...[d][b][a][c][b][a]
type DUP2_X1 struct{ base.NoOperandsInstruction }

func (s *DUP2_X1) Execute(frame *rtda.Frame) {
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

// 复制栈顶的两个槽位, 再把复制的新数据插入到第二位之后, 然后这四个弹出的位置依次压入
// 复制两个后压入, 再压回这四个原数据
// 栈底->栈顶
// ....[d][c][b][a] -> ...[b][a][d][c][b][a]
type DUP2_X2 struct{ base.NoOperandsInstruction }

func (s *DUP2_X2) Execute(frame *rtda.Frame) {
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
