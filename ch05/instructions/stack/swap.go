package stack

import (
	"go-jvm/ch05/instructions/base"
	"go-jvm/ch05/rtda"
)

// 交换指令栈顶的两个位置
// // ...[c][b][a] -> ...[c][a][b]
type SWAP struct{ base.NoOperandsInstruction }

// Execute 交换两个变量
func (s *SWAP) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopSlot()
	v2 := stack.PopSlot()
	stack.PushSlot(v2)
	stack.PushSlot(v1)
}
