package comparisons

import (
	"go-jvm/ch05/instructions/base"
	"go-jvm/ch05/rtda"
)

/**
 * 把操作数栈顶两个引用变量弹出, 然后进行比较, 满足则跳转
 */

// ==  则跳转
type IF_ACMPEQ struct{ base.BranchInstruction }

func (s *IF_ACMPEQ) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopRef()
	v1 := stack.PopRef()
	if v1 == v2 {
		base.Branch(frame, s.Offset)
	}
}

// !=  则跳转
type IF_ACMPNE struct{ base.BranchInstruction }

func (s *IF_ACMPNE) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopRef()
	v1 := stack.PopRef()
	if v1 != v2 {
		base.Branch(frame, s.Offset)
	}
}
