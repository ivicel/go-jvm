package math

import (
	"go-jvm/ch05/instructions/base"
	"go-jvm/ch05/rtda"
	"math"
)

type DREM struct{ base.NoOperandsInstruction }

func (s *DREM) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopDobule()
	v1 := stack.PopDobule()
	result := math.Mod(v1, v2)
	stack.PushDobule(result)
}

type FREM struct{ base.NoOperandsInstruction }

func (s *FREM) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	result := math.Mod(float64(v1), float64(v2))
	stack.PushFloat(float32(result))
}

type IREM struct{ base.NoOperandsInstruction }

func (s *IREM) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()

	if v2 == 0 {
		panic("java.lang.ArithmeticException: / by zero")
	}

	result := v1 % v2
	stack.PushInt(result)
}

type LREM struct{ base.NoOperandsInstruction }

func (s *LREM) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()

	if v2 == 0 {
		panic("java.lang.ArithmeticException: / by zero")
	}

	result := v1 % v2
	stack.PushLong(result)
}
