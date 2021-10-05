package constants

import (
	"go-jvm/ch05/instructions/base"
	"go-jvm/ch05/rtda"
)

// nop 指令, 什么也不做
type NOP struct {
	base.NoOperandsInstruction
}

func (s *NOP) Execute(frame *rtda.Frame) {}
