package classfile

import "encoding/binary"

type ClassReader struct {
	data []byte
}

// 读取 u1
func (r *ClassReader) readUint8() uint8 {
	val := r.data[0]
	r.data = r.data[1:]
	return val
}

// 读取 u2
func (r *ClassReader) readUint16() uint16 {
	val := binary.BigEndian.Uint16(r.data)
	r.data = r.data[2:]
	return val
}

// 读取多个 u2
func (r *ClassReader) readUint16s() []uint16 {
	length := r.readUint16()
	arr := make([]uint16, length)
	for i := range arr {
		arr[i] = r.readUint16()
	}

	return arr
}

// 读取 u4
func (r *ClassReader) readUint32() uint32 {
	val := binary.BigEndian.Uint32(r.data)
	r.data = r.data[4:]
	return val
}

// 读取 u8
func (r *ClassReader) readUint64() uint64 {
	val := binary.BigEndian.Uint64(r.data)
	r.data = r.data[8:]
	return val
}

func (r *ClassReader) readBytes(n uint32) []byte {
	val := r.data[:n]
	r.data = r.data[n:]
	return val
}
