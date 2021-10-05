package math

import (
	"go-jvm/ch05/instructions/base"
	"go-jvm/ch05/rtda"
)

// double 乘 double
type DMUL struct{ base.NoOperandsInstruction }

func (s *DMUL) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopDobule()
	v1 := stack.PopDobule()
	result := v1 * v2
	stack.PushDobule(result)
}

// float 乘 float
type FMUL struct{ base.NoOperandsInstruction }

func (s *FMUL) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	result := v1 * v2
	stack.PushFloat(result)
}

// int 乘 int
type IMUL struct{ base.NoOperandsInstruction }

func (s *IMUL) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	result := v1 * v2
	stack.PushInt(result)
}

// long 乘 long
type LMUL struct{ base.NoOperandsInstruction }

func (s *LMUL) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	result := v1 * v2
	stack.PushLong(result)
}
