package stores

import (
	"go-jvm/ch05/instructions/base"
	"go-jvm/ch05/rtda"
)

/**
 * 存储栈顶数据 float 类型到局部变量表
 */

func _fstore(frame *rtda.Frame, index uint) {
	val := frame.OperandStack().PopFloat()
	frame.LocalVars().SetFloat(index, val)
}

type FSTORE struct{ base.Index8Instruction }

func (s *FSTORE) Execute(frame *rtda.Frame) {
	_fstore(frame, s.Index)
}

type FSTORE_0 struct{ base.NoOperandsInstruction }

func (s *FSTORE_0) Execute(frame *rtda.Frame) {
	_fstore(frame, 0)
}

type FSTORE_1 struct{ base.NoOperandsInstruction }

func (s *FSTORE_1) Execute(frame *rtda.Frame) {
	_fstore(frame, 1)
}

type FSTORE_2 struct{ base.NoOperandsInstruction }

func (s *FSTORE_2) Execute(frame *rtda.Frame) {
	_fstore(frame, 2)
}

type FSTORE_3 struct{ base.NoOperandsInstruction }

func (s *FSTORE_3) Execute(frame *rtda.Frame) {
	_fstore(frame, 3)
}
