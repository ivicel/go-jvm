package loads

import (
	"go-jvm/ch05/instructions/base"
	"go-jvm/ch05/rtda"
)

/**
 * 加载指令从局部变量表获取变量, 然后从推入操作数栈顶
 */

func _lload(frame *rtda.Frame, index uint) {
	val := frame.LocalVars().GetLong(index)
	frame.OperandStack().PushLong(val)
}

type LLOAD struct{ base.Index8Instruction }

// 将局部变量表 int 值压入操作数栈顶
func (s *LLOAD) Execute(frame *rtda.Frame) {
	_lload(frame, s.Index)
}

type LLOAD_0 struct{ base.NoOperandsInstruction }

// 将局部变量表 int 值压入操作数栈顶
func (s *LLOAD_0) Execute(frame *rtda.Frame) {
	_lload(frame, 0)
}

type LLOAD_1 struct{ base.NoOperandsInstruction }

// 将局部变量表 int 值压入操作数栈顶
func (s *LLOAD_1) Execute(frame *rtda.Frame) {
	_lload(frame, 1)
}

type LLOAD_2 struct{ base.NoOperandsInstruction }

// 将局部变量表 int 值压入操作数栈顶
func (s *LLOAD_2) Execute(frame *rtda.Frame) {
	_lload(frame, 2)
}

type LLOAD_3 struct{ base.NoOperandsInstruction }

// 将局部变量表 int 值压入操作数栈顶
func (s *LLOAD_3) Execute(frame *rtda.Frame) {
	_lload(frame, 3)
}
