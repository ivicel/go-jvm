package constants

import (
	"go-jvm/ch05/instructions/base"
	"go-jvm/ch05/rtda"
)

/**
 * 加载常量值到栈顶
 */

type ACONST_NULL struct{ base.NoOperandsInstruction }

// 将常量 null 压入操作数栈顶
func (s *ACONST_NULL) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushRef(nil)
}

type DCONST_0 struct{ base.NoOperandsInstruction }

// 将常量 double 值 0.0 压入操作数栈顶
func (s *DCONST_0) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushDobule(0.0)
}

type DCONST_1 struct{ base.NoOperandsInstruction }

// 将常量 double 值 1.0 压入操作数栈顶
func (s *DCONST_1) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushDobule(1.0)
}

type FCONST_0 struct{ base.NoOperandsInstruction }

// 将常量 float 值 0.0 压入操作数栈顶
func (s *FCONST_0) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushFloat(0.0)
}

type FCONST_1 struct{ base.NoOperandsInstruction }

// 将常量 float 值 1.0 压入操作数栈顶
func (s *FCONST_1) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushFloat(1.0)
}

type FCONST_2 struct{ base.NoOperandsInstruction }

// 将常量 float 值 2.0 压入操作数栈顶
func (s *FCONST_2) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushFloat(2.0)
}

type ICONST_M1 struct{ base.NoOperandsInstruction }

// 将常量 int 值 -1 压入操作数栈顶
func (s *ICONST_M1) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(-1)
}

type ICONST_0 struct{ base.NoOperandsInstruction }

// 将常量 int 值 0 压入操作数栈顶
func (s *ICONST_0) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(0)
}

type ICONST_1 struct{ base.NoOperandsInstruction }

// 将常量 int 值 1 压入操作数栈顶
func (s *ICONST_1) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(1)
}

type ICONST_2 struct{ base.NoOperandsInstruction }

// 将常量 int 值 2 压入操作数栈顶
func (s *ICONST_2) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(2)
}

type ICONST_3 struct{ base.NoOperandsInstruction }

// 将常量 int 值 3 压入操作数栈顶
func (s *ICONST_3) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(3)
}

type ICONST_4 struct{ base.NoOperandsInstruction }

// 将常量 int 值 4 压入操作数栈顶
func (s *ICONST_4) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(4)
}

type ICONST_5 struct{ base.NoOperandsInstruction }

// 将常量 int 值 5 压入操作数栈顶
func (s *ICONST_5) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(5)
}

type LCONST_0 struct{ base.NoOperandsInstruction }

// 将常量 long 值 0 压入操作数栈顶
func (s *LCONST_0) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(0)
}

type LCONST_1 struct{ base.NoOperandsInstruction }

// 将常量 long 值 1 压入操作数栈顶
func (s *LCONST_1) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(1)
}
