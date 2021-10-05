package conversion

import (
	"go-jvm/ch05/instructions/base"
	"go-jvm/ch05/rtda"
)

// double 强制转换为 float
type D2F struct{ base.NoOperandsInstruction }

func (s *D2F) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	d := stack.PopDobule()
	f := float32(d)
	stack.PushFloat(f)
}

// double 强制转换为 long
type D2L struct{ base.NoOperandsInstruction }

func (s *D2L) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	d := stack.PopDobule()
	l := int64(d)
	stack.PushLong(l)
}

// double 强制转换为 int
type D2I struct{ base.NoOperandsInstruction }

func (s *D2I) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	d := stack.PopDobule()
	i := int32(d)
	stack.PushInt(i)
}
