package control

import (
	"go-jvm/ch05/instructions/base"
	"go-jvm/ch05/rtda"
)

type LOOKUP_SWITCH struct {
	defaultOffset int32
	npair         int32
	matchOffsets  []int32
}

func (s *LOOKUP_SWITCH) FetchOperands(reader *base.BytecodeReader) {
	reader.SkipPadding()
	s.defaultOffset = reader.ReadInt32()
	s.npair = reader.ReadInt32()
	s.matchOffsets = reader.ReadInt32s(s.npair * 2)
}

func (s *LOOKUP_SWITCH) Execute(frame *rtda.Frame) {
	key := frame.OperandStack().PopInt()
	for i := int32(0); i < s.npair*2; i += 2 {
		if s.matchOffsets[i] == key {
			offset := s.matchOffsets[i+1]
			base.Branch(frame, int(offset))
			return
		}
	}

	base.Branch(frame, int(s.defaultOffset))
}
