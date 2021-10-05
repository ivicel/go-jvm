package classfile

// java 字段或方法结构表示
// filed_info {
// 	u2 access_flags;
// 	u2 name_index;
// 	u2 descriptor_index;
// 	attribute_info attributes[attributes_count];
// }

// java 字段或方法结构体
type MemberInfo struct {
	cp              ConstantPool    // 所指向常量池
	accessFlags     uint16          // 访问权限
	nameIndex       uint16          // 名称指向
	descriptorIndex uint16          // 描述
	attributes      []AttributeInfo // 属性
}

// readMembers 读取字段或方法数据
func readMembers(reader *ClassReader, cp ConstantPool) []*MemberInfo {
	length := reader.readUint16()
	members := make([]*MemberInfo, length)
	for i := 0; i < int(length); i++ {
		members[i] = &MemberInfo{
			cp:              cp,
			accessFlags:     reader.readUint16(),
			nameIndex:       reader.readUint16(),
			descriptorIndex: reader.readUint16(),
			attributes:      readAttributes(reader, cp),
		}
	}

	return members
}

// Name 返回字段或者方法的名称
func (m *MemberInfo) Name() string {
	return m.cp.getUtf8(m.nameIndex)
}

// Descriptor 返回字段或方法描述符
func (m *MemberInfo) Descriptor() string {
	return m.cp.getUtf8(m.descriptorIndex)
}

func (m *MemberInfo) CodeAttribute() *CodeAttribute {
	for _, attrInfo := range m.attributes {
		if a, ok := attrInfo.(*CodeAttribute); ok {
			return a
		}
	}

	return nil
}
