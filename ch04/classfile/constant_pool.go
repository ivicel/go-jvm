package classfile

// 常量池
type ConstantPool []ConstantInfo

func (c *ConstantPool) getClassName(index uint16) string {
	classInfo := (*c)[index].(*ConstantClassInfo)
	return c.getUtf8(classInfo.nameIndex)
}

// 读一个 utf8 字符串
func (c *ConstantPool) getUtf8(index uint16) string {
	utf8Info := (*c)[index].(*ConstantUtf8Info)
	return utf8Info.str
}

// 读一个 name_and_type 描述符
func (c *ConstantPool) getNameAndType(index uint16) (string, string) {
	ntInfo := (*c)[index].(*ConstantNameAndTypeInfo)
	name := c.getUtf8(ntInfo.nameIndex)
	_type := c.getUtf8(ntInfo.descriptorIndex)
	return name, _type
}

func readConstantPool(reader *ClassReader) ConstantPool {
	// 常量池的大小
	cpCount := reader.readUint16()
	cp := make([]ConstantInfo, cpCount)

	// 常量池总数为 n-1 个, 下标从 1 开始
	for i := 1; i < int(cpCount); i++ {
		cp[i] = readConstantInfo(reader, cp)
		cp[i].readInfo(reader)
		switch cp[i].(type) {
		// long, double 类型算两个
		case *ConstantLongInfo, *ConstantDoubleInfo:
			i++
		}
	}

	return cp
}
