package constants

import (
	"go-jvm/ch05/instructions/base"
	"go-jvm/ch05/rtda"
)

// 从操作数中获取一个 byte 型整数, 扩展成 int 型, 然后推入栈顶
type BIPUSH struct{ val int8 }

func (s *BIPUSH) FetchOperands(reader *base.BytecodeReader) {
	s.val = reader.ReadInt8()
}
func (s *BIPUSH) Execute(frame *rtda.Frame) {
	i := int32(s.val)
	frame.OperandStack().PushInt(i)
}

// 从操作数中获取一个 byte 型整数, 扩展成 short 型, 然后推入栈顶
type SIPUSH struct{ val int16 }

func (s *SIPUSH) FetchOperands(reader *base.BytecodeReader) {
	s.val = reader.ReadInt16()
}
func (s *SIPUSH) Execute(frame *rtda.Frame) {
	i := int32(s.val)
	frame.OperandStack().PushInt(i)
}
