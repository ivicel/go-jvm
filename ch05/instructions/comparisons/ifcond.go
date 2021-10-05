package comparisons

import (
	"go-jvm/ch05/instructions/base"
	"go-jvm/ch05/rtda"
)

/**
 * 把操作数栈顶 int 变量弹出, 然后跟 0 进行比较, 满足则跳转
 */

// == 0 则跳转
type IFEQ struct{ base.BranchInstruction }

func (s *IFEQ) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val == 0 {
		base.Branch(frame, s.Offset)
	}
}

// != 0 则跳转
type IFNE struct{ base.BranchInstruction }

func (s *IFNE) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val != 0 {
		base.Branch(frame, s.Offset)
	}
}

// < 0 则跳转
type IFLT struct{ base.BranchInstruction }

func (s *IFLT) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val < 0 {
		base.Branch(frame, s.Offset)
	}
}

// <= 0 则跳转
type IFLE struct{ base.BranchInstruction }

func (s *IFLE) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val <= 0 {
		base.Branch(frame, s.Offset)
	}
}

// > 0 则跳转
type IFGT struct{ base.BranchInstruction }

func (s *IFGT) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val > 0 {
		base.Branch(frame, s.Offset)
	}
}

// >= 0 则跳转
type IFGE struct{ base.BranchInstruction }

func (s *IFGE) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val >= 0 {
		base.Branch(frame, s.Offset)
	}
}
