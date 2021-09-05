package classfile

// ConstantValue_attribute {
// 	u2 attribute_name_index;
// 	u4 attribute_length;
// 	u2 constantvalue_index;
// }

// 常量属性
type ConstantValueAttribute struct {
	constantValueIndex uint16
}

func (attr *ConstantValueAttribute) readInfo(reader *ClassReader) {
	attr.constantValueIndex = reader.readUint16()
}

func (attr *ConstantValueAttribute) ConstantValueIndex() uint16 {
	return attr.constantValueIndex
}

func (attr *ConstantValueAttribute) String() string {
	return "ConstantValue"
}
