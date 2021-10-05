package classfile

type ConstantMemberInfo struct {
	cp               ConstantPool
	classIndex       uint16 // 类的索引
	nameAndTypeIndex uint16 // name_and_type 结构的索引
}

func (c *ConstantMemberInfo) readInfo(reader *ClassReader) {
	c.classIndex = reader.readUint16()
	c.nameAndTypeIndex = reader.readUint16()
}

// ClassName 类名称
func (c *ConstantMemberInfo) ClassName() string {
	return c.cp.getClassName(c.classIndex)
}

// NameAndDescriptor 描述符说明
func (c *ConstantMemberInfo) NameAndDescriptor() (string, string) {
	return c.cp.getNameAndType(c.nameAndTypeIndex)
}

// 字段符号引用
//
// CONSTANT_Fieldref_info {
// 	u1 tag;
// 	u2 class_index;
// 	u2 name_and_type_index;
// }
type ConstantFieldrefInfo = ConstantMemberInfo

// 普通(非接口)方法符号引用
//
// CONSTANT_Methodref_info {
// 	u1 tag;
// 	u2 class_index;
// 	u2 name_and_type_index;
// }
type ConstantMethodrefInfo = ConstantMemberInfo

// 接口方法符号引用
//
// CONSTANT_InterfaceMethodref_info {
// 	u1 tag;
// 	u2 class_index;
// 	u2 name_and_type_index;
// }
type ConstantInterfaceMethodrefInfo = ConstantMemberInfo
