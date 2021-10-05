package math

import (
	"go-jvm/ch05/instructions/base"
	"go-jvm/ch05/rtda"
)

/**
 * 加法指令
 */

// double 加法指令
type DADD struct{ base.NoOperandsInstruction }

func (s *DADD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopDobule()
	v2 := stack.PopDobule()
	result := v1 + v2
	stack.PushDobule(result)
}

// float 加法指令
type FADD struct{ base.NoOperandsInstruction }

func (s *FADD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopFloat()
	v2 := stack.PopFloat()
	result := v1 + v2
	stack.PushFloat(result)
}

// int 加法指令
type IADD struct{ base.NoOperandsInstruction }

func (s *IADD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopInt()
	v2 := stack.PopInt()
	result := v1 + v2
	stack.PushInt(result)
}

// long 加法指令
type LADD struct{ base.NoOperandsInstruction }

func (s *LADD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopLong()
	v2 := stack.PopLong()
	result := v1 + v2
	stack.PushLong(result)
}
