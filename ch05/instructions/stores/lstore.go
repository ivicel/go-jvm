package stores

import (
	"go-jvm/ch05/instructions/base"
	"go-jvm/ch05/rtda"
)

/**
 * 存储栈顶数据 long 类型到局部变量表
 */

func _lstore(frame *rtda.Frame, index uint) {
	val := frame.OperandStack().PopLong()
	frame.LocalVars().SetLong(index, val)
}

type LSTORE struct{ base.Index8Instruction }

func (s *LSTORE) Execute(frame *rtda.Frame) {
	_lstore(frame, s.Index)
}

type LSTORE_0 struct{ base.NoOperandsInstruction }

func (s *LSTORE_0) Execute(frame *rtda.Frame) {
	_lstore(frame, 0)
}

type LSTORE_1 struct{ base.NoOperandsInstruction }

func (s *LSTORE_1) Execute(frame *rtda.Frame) {
	_lstore(frame, 1)
}

type LSTORE_2 struct{ base.NoOperandsInstruction }

func (s *LSTORE_2) Execute(frame *rtda.Frame) {
	_lstore(frame, 2)
}

type LSTORE_3 struct{ base.NoOperandsInstruction }

func (s *LSTORE_3) Execute(frame *rtda.Frame) {
	_lstore(frame, 3)
}
