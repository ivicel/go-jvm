package extended

import (
	"go-jvm/ch05/instructions/base"
	"go-jvm/ch05/rtda"
)

type IFNULL struct{ base.BranchInstruction }

func (s *IFNULL) Execute(frame *rtda.Frame) {
	ref := frame.OperandStack().PopRef()
	if ref == nil {
		base.Branch(frame, s.Offset)
	}
}

type IFNONNULL struct{ base.BranchInstruction }

func (s *IFNONNULL) Execute(frame *rtda.Frame) {
	ref := frame.OperandStack().PopRef()
	if ref != nil {
		base.Branch(frame, s.Offset)
	}
}
