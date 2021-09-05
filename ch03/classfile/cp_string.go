package classfile

// java.lang.String 字面量, 这个结构并存放数据, 有一个指向 utf8 结构位置索引
//
// CONSTANT_String_info {
// 	u1 tag;
// 	u2 string_index;
// }
type ConstantStringInfo struct {
	cp          ConstantPool
	stringIndex uint16
}

func (c *ConstantStringInfo) readInfo(reader *ClassReader) {
	c.stringIndex = reader.readUint16()
}

func (c *ConstantStringInfo) String() string {
	return c.cp.getUtf8(c.stringIndex)
}
