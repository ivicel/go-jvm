package stores

import (
	"go-jvm/ch05/instructions/base"
	"go-jvm/ch05/rtda"
)

/**
 * 存储栈顶数据引用类型到局部变量表
 */

func _astore(frame *rtda.Frame, index uint) {
	val := frame.OperandStack().PopRef()
	frame.LocalVars().SetRef(index, val)
}

type ASTORE struct{ base.Index8Instruction }

func (s *ASTORE) Execute(frame *rtda.Frame) {
	_astore(frame, s.Index)
}

type ASTORE_0 struct{ base.NoOperandsInstruction }

func (s *ASTORE_0) Execute(frame *rtda.Frame) {
	_astore(frame, 0)
}

type ASTORE_1 struct{ base.NoOperandsInstruction }

func (s *ASTORE_1) Execute(frame *rtda.Frame) {
	_astore(frame, 1)
}

type ASTORE_2 struct{ base.NoOperandsInstruction }

func (s *ASTORE_2) Execute(frame *rtda.Frame) {
	_astore(frame, 2)
}

type ASTORE_3 struct{ base.NoOperandsInstruction }

func (s *ASTORE_3) Execute(frame *rtda.Frame) {
	_astore(frame, 3)
}
