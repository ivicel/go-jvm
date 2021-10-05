package math

import (
	"go-jvm/ch05/instructions/base"
	"go-jvm/ch05/rtda"
)

/**
 * 位移指令
 */

// int 数值左移
type ISHL struct{ base.NoOperandsInstruction }

func (s *ISHL) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	// int 型数据可位移的位数只有 0-31 位, 所以要把位移距离和 0x1f 做一个与, 在 jvm 规范指令有提到
	len := uint32(v2) & 0x1f
	result := v1 << len
	stack.PushInt(result)
}

// int 数值右移
type ISHR struct{ base.NoOperandsInstruction }

func (s *ISHR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	len := uint32(v2) & 0x1f
	result := v1 >> len
	stack.PushInt(result)
}

// int 逻辑左移
type IUSHL struct{ base.NoOperandsInstruction }

func (s *IUSHL) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	len := uint32(v2) & 0x1f
	result := int32(uint32(v1) << len)
	stack.PushInt(result)
}

// int 逻辑右移
type IUSHR struct{ base.NoOperandsInstruction }

func (s *IUSHR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	len := uint32(v2) & 0x1f
	result := int32(uint32(v1) >> len)
	stack.PushInt(result)
}

// long 数值左移
type LSHL struct{ base.NoOperandsInstruction }

func (s *LSHL) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopLong()
	// int 型数据可位移的位数只有 0-63 位, 所以要把位移距离和 0x3f 做一个与, 在 jvm 规范指令有提到
	len := uint32(v2) & 0x3f
	result := v1 << len
	stack.PushLong(result)
}

// long 数值右移
type LSHR struct{ base.NoOperandsInstruction }

func (s *LSHR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopLong()
	len := uint32(v2) & 0x3f
	result := v1 >> len
	stack.PushLong(result)
}

// long 逻辑左移
type LUSHL struct{ base.NoOperandsInstruction }

func (s *LUSHL) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopLong()
	len := uint32(v2) & 0x3f
	result := int64(uint64(v1) << len)
	stack.PushLong(result)
}

// long 逻辑右移
type LUSHR struct{ base.NoOperandsInstruction }

func (s *LUSHR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopLong()
	len := uint32(v2) & 0x3f
	result := int64(uint64(v1) >> len)
	stack.PushLong(result)
}
