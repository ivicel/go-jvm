package math

import (
	"go-jvm/ch05/instructions/base"
	"go-jvm/ch05/rtda"
)

// int 型异或 ^ 操作
type IXOR struct{ base.NoOperandsInstruction }

func (s *IXOR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	result := v1 ^ v2
	stack.PushInt(result)
}

// long 型异或 ^ 操作
type LXOR struct{ base.NoOperandsInstruction }

func (s *LXOR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	result := v1 ^ v2
	stack.PushLong(result)
}
