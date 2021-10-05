package classfile

// 属性表结构
//
// attribute_info {
// 	u2 attribute_name_index; 			// 名称索引
// 	u4 attribute_length;				// 属性长度
// 	u1 info[attribute_length];			// 具体属性结构, 不同属性结构不相同
// }
type AttributeInfo interface {
	readInfo(reader *ClassReader)
}

// readAttributes 读取属性表
func readAttributes(reader *ClassReader, cp ConstantPool) []AttributeInfo {
	// 属性数量
	attributeCount := reader.readUint16()
	attributes := make([]AttributeInfo, attributeCount)
	for i := 0; i < int(attributeCount); i++ {
		attributes[i] = readAttribute(reader, cp)
	}

	return attributes
}

// readAttribute 读取属性
func readAttribute(reader *ClassReader, cp ConstantPool) AttributeInfo {
	attrNameIndex := reader.readUint16()
	attrName := cp.getUtf8(attrNameIndex)
	attrLen := reader.readUint32()
	attrInfo := newAttributeInfo(attrName, attrLen, cp)
	attrInfo.readInfo(reader)

	return attrInfo
}

// newAttributeInfo 组装一个属性结构
func newAttributeInfo(attrName string, attrLen uint32, cp ConstantPool) AttributeInfo {
	switch attrName {
	case "Code":
		return &CodeAttribute{cp: cp}
	case "ConstantValue":
		return &ConstantValueAttribute{}
	case "Deprecated":
		return &DeprecatedAttribute{}
	case "Exceptions":
		return &ExceptionsAttribute{}
	case "LineNumberTable":
		return &LineNumberTableAttribute{}
	case "LocalVariableTable":
		return &LocalVariableTableAttribute{}
	case "SourceFile":
		return &SourceFileAttribute{}
	case "Synthetic":
		return &SyntheticAttribute{}
	default:
		return &UnparsedAttribute{attrName, attrLen, nil}

	}
}