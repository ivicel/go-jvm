package classfile

import "math"

// 整型常量
//
// CONSTANT_Integer_info {
// 	u1 tag;
// 	u4 bytes;
// }
type ConstantIntegerInfo struct {
	val int32 // 整型数值
}

func (c *ConstantIntegerInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint32()
	c.val = int32(bytes)
}

// 浮点常量
//
// CONSTANT_Float_info {
// 	u1 tag;
// 	u4 bytes;
// }
type ConstantFloatInfo struct {
	val float32
}

func (c *ConstantFloatInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint32()
	c.val = math.Float32frombits(bytes)
}

// 长整型
//
// CONSTANT_Long_info {
// 	u1 tag;
// 	u4 high_bytes; // 高位字节
// 	u4 low_bytes; // 低位字节
// }
type ConstantLongInfo struct {
	val int64
}

func (c *ConstantLongInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint64()
	c.val = int64(bytes)
}

// 双精度类型
//
// CONSTANT_Double_info {
// 	u1 tag;
// 	u4 high_bytes;
// 	u4 low_bytes;
// }
type ConstantDoubleInfo struct {
	val float64
}

func (c *ConstantDoubleInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint64()
	c.val = math.Float64frombits(bytes)
}