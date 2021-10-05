package comparisons

import (
	"go-jvm/ch05/instructions/base"
	"go-jvm/ch05/rtda"
)

// long 型比较指令
type LCMP struct{ base.NoOperandsInstruction }

func (s *LCMP) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	if v1 > v2 {
		stack.PushInt(1)
	} else if v1 == v2 {
		stack.PushInt(0)
	} else {
		stack.PushInt(-1)
	}
}
