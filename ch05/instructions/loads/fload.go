package loads

import (
	"go-jvm/ch05/instructions/base"
	"go-jvm/ch05/rtda"
)

/**
 * 加载指令从局部变量表获取变量, 然后从推入操作数栈顶
 */

func _fload(frame *rtda.Frame, index uint) {
	val := frame.LocalVars().GetFloat(index)
	frame.OperandStack().PushFloat(val)
}

type FLOAD struct{ base.Index8Instruction }

// 将局部变量表 float 值压入操作数栈顶
func (s *FLOAD) Execute(frame *rtda.Frame) {
	_fload(frame, s.Index)
}

type FLOAD_0 struct{ base.NoOperandsInstruction }

// 将局部变量表 float 值压入操作数栈顶
func (s *FLOAD_0) Execute(frame *rtda.Frame) {
	_fload(frame, 0)
}

type FLOAD_1 struct{ base.NoOperandsInstruction }

// 将局部变量表 float 值压入操作数栈顶
func (s *FLOAD_1) Execute(frame *rtda.Frame) {
	_fload(frame, 1)
}

type FLOAD_2 struct{ base.NoOperandsInstruction }

// 将局部变量表 float 值压入操作数栈顶
func (s *FLOAD_2) Execute(frame *rtda.Frame) {
	_fload(frame, 2)
}

type FLOAD_3 struct{ base.NoOperandsInstruction }

// 将局部变量表 float 值压入操作数栈顶
func (s *FLOAD_3) Execute(frame *rtda.Frame) {
	_fload(frame, 3)
}
