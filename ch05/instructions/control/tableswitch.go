package control

import (
	"go-jvm/ch05/instructions/base"
	"go-jvm/ch05/rtda"
)

// switch-case 语句实现, 留两种实现方法
// 如果 case 值 可以编码成一个索引表, 则实现成 tableswitch 指令
// 比如下面就是可以实现 tableswitch
// int chooseNear(int i) {
//		switch (i) {
//			case 0: return 0;
//			case 1: return 1;
//			case 2: return 2;
//			default: return -1
//		}
//	}
type TABLE_SWITCH struct {
	defaultOffset int32
	low           int32
	high          int32
	jumpOffsets   []int32
}

func (s *TABLE_SWITCH) FetchOperands(reader *base.BytecodeReader) {
	reader.SkipPadding()
	s.defaultOffset = reader.ReadInt32()
	s.low = reader.ReadInt32()
	s.high = reader.ReadInt32()
	jumpOffsetsCount := s.high - s.low
	s.jumpOffsets = reader.ReadInt32s(jumpOffsetsCount)
}

func (s *TABLE_SWITCH) Execute(frame *rtda.Frame) {
	index := frame.OperandStack().PopInt()
	var offset int
	if index >= s.low && index <= s.high {
		offset = int(s.jumpOffsets[index-s.low])
	} else {
		offset = int(s.defaultOffset)
	}

	base.Branch(frame, offset)
}
