package classfile

// 类或者接口的符号引用, 也是存放一个指向常量池的索引
//
// CONSTANT_Class_info {
// 	u1 tag;
// 	u2 name_index;
// }
type ConstantClassInfo struct {
	cp        ConstantPool
	nameIndex uint16
}

func (c *ConstantClassInfo) readInfo(reader *ClassReader) {
	c.nameIndex = reader.readUint16()
}

func (c *ConstantClassInfo) String() string {
	return c.cp.getUtf8(c.nameIndex)
}
