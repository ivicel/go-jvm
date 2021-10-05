package stores

import (
	"go-jvm/ch05/instructions/base"
	"go-jvm/ch05/rtda"
)

/**
 * 存储栈顶数据 int 类型到局部变量表
 */

func _istore(frame *rtda.Frame, index uint) {
	val := frame.OperandStack().PopInt()
	frame.LocalVars().SetInt(index, val)
}

type ISTORE struct{ base.Index8Instruction }

func (s *ISTORE) Execute(frame *rtda.Frame) {
	_istore(frame, s.Index)
}

type ISTORE_0 struct{ base.NoOperandsInstruction }

func (s *ISTORE_0) Execute(frame *rtda.Frame) {
	_istore(frame, 0)
}

type ISTORE_1 struct{ base.NoOperandsInstruction }

func (s *ISTORE_1) Execute(frame *rtda.Frame) {
	_istore(frame, 1)
}

type ISTORE_2 struct{ base.NoOperandsInstruction }

func (s *ISTORE_2) Execute(frame *rtda.Frame) {
	_istore(frame, 2)
}

type ISTORE_3 struct{ base.NoOperandsInstruction }

func (s *ISTORE_3) Execute(frame *rtda.Frame) {
	_istore(frame, 3)
}
