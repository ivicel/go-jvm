package math

import (
	"go-jvm/ch05/instructions/base"
	"go-jvm/ch05/rtda"
)

// double 取负指令
type DNEG struct{ base.NoOperandsInstruction }

func (s *DNEG) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopDobule()
	stack.PushDobule(-v1)
}

// float 取负指令
type FNEG struct{ base.NoOperandsInstruction }

func (s *FNEG) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopFloat()
	stack.PushFloat(-v1)
}

// int 取负指令
type INEG struct{ base.NoOperandsInstruction }

func (s *INEG) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopInt()
	stack.PushInt(-v1)
}

// long 取负指令
type LNEG struct{ base.NoOperandsInstruction }

func (s *LNEG) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopLong()
	stack.PushLong(-v1)
}
