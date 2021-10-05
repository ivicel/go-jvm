package stores

import (
	"go-jvm/ch05/instructions/base"
	"go-jvm/ch05/rtda"
)

/**
 * 存储栈顶数据 double 类型到局部变量表
 */

func _dstore(frame *rtda.Frame, index uint) {
	val := frame.OperandStack().PopDobule()
	frame.LocalVars().SetDouble(index, val)
}

type DSTORE struct{ base.Index8Instruction }

func (s *DSTORE) Execute(frame *rtda.Frame) {
	_dstore(frame, s.Index)
}

type DSTORE_0 struct{ base.NoOperandsInstruction }

func (s *DSTORE_0) Execute(frame *rtda.Frame) {
	_dstore(frame, 0)
}

type DSTORE_1 struct{ base.NoOperandsInstruction }

func (s *DSTORE_1) Execute(frame *rtda.Frame) {
	_dstore(frame, 1)
}

type DSTORE_2 struct{ base.NoOperandsInstruction }

func (s *DSTORE_2) Execute(frame *rtda.Frame) {
	_dstore(frame, 2)
}

type DSTORE_3 struct{ base.NoOperandsInstruction }

func (s *DSTORE_3) Execute(frame *rtda.Frame) {
	_dstore(frame, 3)
}
