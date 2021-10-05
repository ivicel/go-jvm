package extended

import (
	"go-jvm/ch05/instructions/base"
	"go-jvm/ch05/instructions/loads"
	"go-jvm/ch05/instructions/math"
	"go-jvm/ch05/instructions/stores"
	"go-jvm/ch05/rtda"
)

type WIDE struct {
	modifiedInstruction base.Instruction
}

func (s *WIDE) FetchOperands(reader *base.BytecodeReader) {
	opcode := reader.ReadUint8()
	switch opcode {
	case 0x15: // iload
		inst := &loads.ILOAD{}
		inst.Index = uint(reader.ReadUint16())
		s.modifiedInstruction = inst
	case 0x16: // lload
		inst := &loads.LLOAD{}
		inst.Index = uint(reader.ReadUint16())
		s.modifiedInstruction = inst
	case 0x17: // fload
		inst := &loads.FLOAD{}
		inst.Index = uint(reader.ReadUint16())
		s.modifiedInstruction = inst
	case 0x18: // dload
		inst := &loads.DLOAD{}
		inst.Index = uint(reader.ReadUint16())
		s.modifiedInstruction = inst
	case 0x19: // aload
		inst := &loads.ALOAD{}
		inst.Index = uint(reader.ReadUint16())
		s.modifiedInstruction = inst
	case 0x36: // istore
		inst := &stores.ISTORE{}
		inst.Index = uint(reader.ReadUint16())
		s.modifiedInstruction = inst
	case 0x37: // lstore
		inst := &stores.LSTORE{}
		inst.Index = uint(reader.ReadUint16())
		s.modifiedInstruction = inst
	case 0x38: // fstore
		inst := &stores.FSTORE{}
		inst.Index = uint(reader.ReadUint16())
		s.modifiedInstruction = inst
	case 0x39: // dstore
		inst := &stores.DSTORE{}
		inst.Index = uint(reader.ReadUint16())
		s.modifiedInstruction = inst
	case 0x4a: // astore
		inst := &stores.ASTORE{}
		inst.Index = uint(reader.ReadUint16())
		s.modifiedInstruction = inst
	case 0x84: // iinc
		inst := &math.IINC{}
		inst.Index = uint(reader.ReadUint16())
		inst.Const = int32(reader.ReadInt16())
		s.modifiedInstruction = inst
	case 0xa9:
		// ref
		panic("Unsupported opcode: 0xa9!")
	}
}

func (s *WIDE) Execute(frame *rtda.Frame) {
	s.modifiedInstruction.Execute(frame)
}
