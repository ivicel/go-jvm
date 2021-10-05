package math

import (
	"go-jvm/ch05/instructions/base"
	"go-jvm/ch05/rtda"
)

// double 除 double
type DDIV struct{ base.NoOperandsInstruction }

func (s *DDIV) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopDobule()
	v1 := stack.PopDobule()
	result := v1 / v2
	stack.PushDobule(result)
}

// float 除 float
type FDIV struct{ base.NoOperandsInstruction }

func (s *FDIV) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	result := v1 / v2
	stack.PushFloat(result)
}

// int 除 int
type IDIV struct{ base.NoOperandsInstruction }

func (s *IDIV) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	if v2 == 0 {
		panic("java.lang.ArithmeticException: / by zero")
	}

	result := v1 / v2
	stack.PushInt(result)
}

// long 除 long
type LDIV struct{ base.NoOperandsInstruction }

func (s *LDIV) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	if v2 == 0 {
		panic("java.lang.ArithmeticException: / by zero")
	}

	result := v1 / v2
	stack.PushLong(result)
}
