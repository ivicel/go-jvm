package extended

import (
	"go-jvm/ch05/instructions/base"
	"go-jvm/ch05/rtda"
)

type GOTO_W struct {
	offset int
}

func (s *GOTO_W) FetchOperands(reader *base.BytecodeReader) {
	s.offset = int(reader.ReadInt32())
}

func (s *GOTO_W) Execute(frame *rtda.Frame) {
	base.Branch(frame, s.offset)
}
