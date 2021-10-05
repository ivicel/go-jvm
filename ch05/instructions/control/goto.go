package control

import (
	"go-jvm/ch05/instructions/base"
	"go-jvm/ch05/rtda"
)

// goto 指令进行无条件跳转
type GOTO struct{ base.BranchInstruction }

func (s *GOTO) Execute(frame *rtda.Frame) {
	base.Branch(frame, s.Offset)
}
