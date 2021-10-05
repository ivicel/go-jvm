package base

type BytecodeReader struct {
	code []byte // 字节码
	pc   int    // 记录读取到哪个字节
}

func (r *BytecodeReader) Reset(code []byte, pc int) {
	r.code = code
	r.pc = pc
}

// ReadUint8 读取无符号一个字节
func (r *BytecodeReader) ReadUint8() uint8 {
	i := r.code[r.pc]
	r.pc++
	return i
}

// ReadInt8 读取无符号一个字节
func (r *BytecodeReader) ReadInt8() int8 {
	return int8(r.ReadUint8())
}

// ReadUint16 读取无符号双字节
func (r *BytecodeReader) ReadUint16() uint16 {
	byte1 := uint16(r.ReadUint8())
	byte2 := uint16(r.ReadUint8())
	return (byte1 << 8) | byte2
}

// ReadInt16 读取有符号双字节
func (r *BytecodeReader) ReadInt16() int16 {
	return int16(r.ReadUint16())
}

// ReadInt32 读取有符号四字节
func (r *BytecodeReader) ReadInt32() int32 {
	byte1 := int32(r.ReadUint8())
	byte2 := int32(r.ReadUint8())
	byte3 := int32(r.ReadUint8())
	byte4 := int32(r.ReadUint8())

	return (byte1 << 24) | (byte2 << 16) | (byte3 << 8) | byte4
}

// SkipPadding tableswitch 操作码后面有 0-3 字节的 padding, 以保证 defaultOffset
// 在字节码中的地址是 4 的倍数
func (r *BytecodeReader) SkipPadding() {
	if r.pc%4 != 0 {
		r.ReadInt8()
	}
}

// ReadInt32s 读取 n 个 int32
func (r *BytecodeReader) ReadInt32s(n int32) []int32 {
	ints := make([]int32, n)
	for i := range ints {
		ints[i] = r.ReadInt32()
	}

	return ints
}

func (r *BytecodeReader) PC() int {
	return r.pc
}
