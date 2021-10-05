package comparisons

import (
	"go-jvm/ch05/instructions/base"
	"go-jvm/ch05/rtda"
)

/**
 * 把操作数栈顶两个 int 变量弹出, 然后进行比较, 满足则跳转
 */

// ==  则跳转
type IF_ICMPEQ struct{ base.BranchInstruction }

func (s *IF_ICMPEQ) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	if v1 == v2 {
		base.Branch(frame, s.Offset)
	}
}

// !=  则跳转
type IF_ICMPNE struct{ base.BranchInstruction }

func (s *IF_ICMPNE) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	if v1 != v2 {
		base.Branch(frame, s.Offset)
	}
}

// < 则跳转
type IF_ICMPLT struct{ base.BranchInstruction }

func (s *IF_ICMPLT) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	if v1 < v2 {
		base.Branch(frame, s.Offset)
	}
}

// <= 则跳转
type IF_ICMPLE struct{ base.BranchInstruction }

func (s *IF_ICMPLE) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	if v1 <= v2 {
		base.Branch(frame, s.Offset)
	}
}

// > 则跳转
type IF_ICMPGT struct{ base.BranchInstruction }

func (s *IF_ICMPGT) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	if v1 > v2 {
		base.Branch(frame, s.Offset)
	}
}

// >= 则跳转
type IF_ICMPGE struct{ base.BranchInstruction }

func (s *IF_ICMPGE) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	if v1 >= v2 {
		base.Branch(frame, s.Offset)
	}
}
