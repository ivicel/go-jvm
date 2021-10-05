package loads

import (
	"go-jvm/ch05/instructions/base"
	"go-jvm/ch05/rtda"
)

func _alod(frame *rtda.Frame, index uint) {
	ref := frame.LocalVars().GetRef(index)
	frame.OperandStack().PushRef(ref)
}

// 加载引用类型到操作数栈
type ALOAD struct{ base.Index8Instruction }

func (s *ALOAD) Execute(frame *rtda.Frame) {
	_alod(frame, s.Index)
}

// 加载引用类型, 位置 0 到操作数栈
type ALOAD_0 struct{ base.NoOperandsInstruction }

func (s *ALOAD_0) Execute(frame *rtda.Frame) {
	_alod(frame, 0)
}

// 加载引用类型, 位置 1 到操作数栈
type ALOAD_1 struct{ base.NoOperandsInstruction }

func (s *ALOAD_1) Execute(frame *rtda.Frame) {
	_alod(frame, 1)
}

// 加载引用类型, 位置 2 到操作数栈
type ALOAD_2 struct{ base.NoOperandsInstruction }

func (s *ALOAD_2) Execute(frame *rtda.Frame) {
	_alod(frame, 2)
}

// 加载引用类型, 位置 3 到操作数栈
type ALOAD_3 struct{ base.NoOperandsInstruction }

func (s *ALOAD_3) Execute(frame *rtda.Frame) {
	_alod(frame, 3)
}
