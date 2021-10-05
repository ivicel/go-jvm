package stack

import (
	"go-jvm/ch05/instructions/base"
	"go-jvm/ch05/rtda"
)

// pop 指令, 弹出 int, float 等占用一个操作数栈位置的变量
// ...[c][b][a] -> ...[c][b]
type POP struct{ base.NoOperandsInstruction }

func (s *POP) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	stack.PopSlot()
}

// pop2 指令, 弹出 long, double 等占用两个操作数栈位置的变量
// ...[c][b][a] -> ...[c]
type POP2 struct{ base.NoOperandsInstruction }

func (s *POP2) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	stack.PopSlot()
	stack.PopSlot()
}
