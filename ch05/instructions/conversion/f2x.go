package conversion

import (
	"go-jvm/ch05/instructions/base"
	"go-jvm/ch05/rtda"
)

// float 强制转换为 double
type F2D struct{ base.NoOperandsInstruction }

func (s *F2D) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	f := stack.PopFloat()
	d := float64(f)
	stack.PushDobule(d)
}

// float 强制转换为 long
type F2L struct{ base.NoOperandsInstruction }

func (s *F2L) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	f := stack.PopFloat()
	l := int64(f)
	stack.PushLong(l)
}

// float 强制转换为 int
type F2I struct{ base.NoOperandsInstruction }

func (s *F2I) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	f := stack.PopFloat()
	i := int32(f)
	stack.PushInt(i)
}
