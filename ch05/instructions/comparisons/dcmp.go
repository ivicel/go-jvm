package comparisons

import (
	"go-jvm/ch05/instructions/base"
	"go-jvm/ch05/rtda"
)

// double 型比较指令
type DCMPG struct{ base.NoOperandsInstruction }

func (s *DCMPG) Execute(frame *rtda.Frame) {
	_dcmp(frame, true)
}

// double 型比较指令
type DCMPL struct{ base.NoOperandsInstruction }

func (s *DCMPL) Execute(frame *rtda.Frame) {
	_dcmp(frame, false)
}

func _dcmp(frame *rtda.Frame, gFlag bool) {
	stack := frame.OperandStack()
	v2 := stack.PopDobule()
	v1 := stack.PopDobule()
	if v1 > v2 {
		stack.PushInt(1)
	} else if v1 == v2 {
		stack.PushInt(0)
	} else if v1 < v2 {
		stack.PushInt(-1)
	} else if gFlag {
		stack.PushInt(1)
	} else {
		stack.PushInt(-1)
	}
}
