package conversion

import (
	"go-jvm/ch05/instructions/base"
	"go-jvm/ch05/rtda"
)

// long 强制转换为 float
type L2F struct{ base.NoOperandsInstruction }

func (s *L2F) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	l := stack.PopLong()
	f := float32(l)
	stack.PushFloat(f)
}

// long 强制转换为 int
type L2L struct{ base.NoOperandsInstruction }

func (s *L2L) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	l := stack.PopLong()
	i := int32(l)
	stack.PushInt(i)
}

// long 强制转换为 double
type L2D struct{ base.NoOperandsInstruction }

func (s *L2D) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	l := stack.PopLong()
	d := float64(l)
	stack.PushDobule(d)
}

// long 强制转换为 int
type L2I struct{ base.NoOperandsInstruction }

func (s *L2I) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	l := stack.PopLong()
	i := int32(l)
	stack.PushInt(i)
}
