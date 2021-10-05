package math

import (
	"go-jvm/ch05/instructions/base"
	"go-jvm/ch05/rtda"
)

/**
 * 减法指令
 */

// double 减法指令
type DSUB struct{ base.NoOperandsInstruction }

func (s *DSUB) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopDobule()
	v2 := stack.PopDobule()
	result := v1 - v2
	stack.PushDobule(result)
}

// float 减法指令
type FSUB struct{ base.NoOperandsInstruction }

func (s *FSUB) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopFloat()
	v2 := stack.PopFloat()
	result := v1 - v2
	stack.PushFloat(result)
}

// int 减法指令
type ISUB struct{ base.NoOperandsInstruction }

func (s *ISUB) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopInt()
	v2 := stack.PopInt()
	result := v1 - v2
	stack.PushInt(result)
}

// long 减法指令
type LSUB struct{ base.NoOperandsInstruction }

func (s *LSUB) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopLong()
	v2 := stack.PopLong()
	result := v1 - v2
	stack.PushLong(result)
}
