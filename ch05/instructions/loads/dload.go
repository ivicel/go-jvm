package loads

import (
	"go-jvm/ch05/instructions/base"
	"go-jvm/ch05/rtda"
)

/**
 * 加载指令从局部变量表获取变量, 然后从推入操作数栈顶
 */

func _dload(frame *rtda.Frame, index uint) {
	val := frame.LocalVars().GetDouble(index)
	frame.OperandStack().PushDobule(val)
}

type DLOAD struct{ base.Index8Instruction }

// 将局部变量表 double 值压入操作数栈顶
func (s *DLOAD) Execute(frame *rtda.Frame) {
	_dload(frame, s.Index)
}

type DLOAD_0 struct{ base.NoOperandsInstruction }

// 将局部变量表 double 值压入操作数栈顶
func (s *DLOAD_0) Execute(frame *rtda.Frame) {
	_dload(frame, 0)
}

type DLOAD_1 struct{ base.NoOperandsInstruction }

// 将局部变量表 double 值压入操作数栈顶
func (s *DLOAD_1) Execute(frame *rtda.Frame) {
	_dload(frame, 1)
}

type DLOAD_2 struct{ base.NoOperandsInstruction }

// 将局部变量表 double 值压入操作数栈顶
func (s *DLOAD_2) Execute(frame *rtda.Frame) {
	_dload(frame, 2)
}

type DLOAD_3 struct{ base.NoOperandsInstruction }

// 将局部变量表 double 值压入操作数栈顶
func (s *DLOAD_3) Execute(frame *rtda.Frame) {
	_dload(frame, 3)
}
