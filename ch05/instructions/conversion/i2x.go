package conversion

import (
	"go-jvm/ch05/instructions/base"
	"go-jvm/ch05/rtda"
)

// int 强制转换为 float
type I2F struct{ base.NoOperandsInstruction }

func (s *I2F) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	i := stack.PopInt()
	f := float32(i)
	stack.PushFloat(f)
}

// int 强制转换为 long
type I2L struct{ base.NoOperandsInstruction }

func (s *I2L) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	i := stack.PopInt()
	l := int64(i)
	stack.PushLong(l)
}

// int 强制转换为 double
type I2D struct{ base.NoOperandsInstruction }

func (s *I2D) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	i := stack.PopInt()
	d := float64(i)
	stack.PushDobule(d)
}

// int 强制转换为 byte
type I2B struct{ base.NoOperandsInstruction }

func (s *I2B) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	i := stack.PopInt()
	d := int32(int8(i))
	stack.PushInt(d)
}


// int 强制转换为 char
type I2C struct{ base.NoOperandsInstruction }

func (s *I2C) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	i := stack.PopInt()
	d := int32(uint16(i))
	stack.PushInt(d)
}

// int 强制转换为 char
type I2S struct{ base.NoOperandsInstruction }

func (s *I2S) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	i := stack.PopInt()
	d := int32(int16(i))
	stack.PushInt(d)
}