package loads

import (
	"go-jvm/ch05/instructions/base"
	"go-jvm/ch05/rtda"
)

/**
 * 加载指令从局部变量表获取变量, 然后从推入操作数栈顶
 */

func _iload(frame *rtda.Frame, index uint) {
	val := frame.LocalVars().GetInt(index)
	frame.OperandStack().PushInt(val)
}

type ILOAD struct{ base.Index8Instruction }

// 将局部变量表 int 值压入操作数栈顶
func (s *ILOAD) Execute(frame *rtda.Frame) {
	_iload(frame, s.Index)
}

type ILOAD_0 struct{ base.NoOperandsInstruction }

// 将局部变量表 int 值压入操作数栈顶
func (s *ILOAD_0) Execute(frame *rtda.Frame) {
	_iload(frame, 0)
}

type ILOAD_1 struct{ base.NoOperandsInstruction }

// 将局部变量表 int 值压入操作数栈顶
func (s *ILOAD_1) Execute(frame *rtda.Frame) {
	_iload(frame, 1)
}

type ILOAD_2 struct{ base.NoOperandsInstruction }

// 将局部变量表 int 值压入操作数栈顶
func (s *ILOAD_2) Execute(frame *rtda.Frame) {
	_iload(frame, 2)
}

type ILOAD_3 struct{ base.NoOperandsInstruction }

// 将局部变量表 int 值压入操作数栈顶
func (s *ILOAD_3) Execute(frame *rtda.Frame) {
	_iload(frame, 3)
}
