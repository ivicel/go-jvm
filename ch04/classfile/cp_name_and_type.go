package classfile

// 字段或方法的名称和描述符(即类型)
//
// CONSTANT_NameAndType_info {
// 	u1 tag;
// 	u2 name_index;
// 	u2 descriptor_index;
// }
//
// 描述符说明:
//
// 1. 基本类型: byte(B), short(S), char(C), int(I), long(J), float(F), double(D) 为单字母, 注意 long 为 J 而不是 L
// 2. 引用类型: L+类全限定名+分号, 如: Ljava.lang.String;
// 3. 数组类型: [+数组元素全限制名, 如: [java.lang.String
// 4. void 类型是 V
// 5. 方法的描述符为: (用分号分隔的参数描述符)+返回值类型描述符,
// 如: void fun() 				-> ()V
// String toString() 			-> ()Ljava.lang.String;
// void main(String[] args) 	-> ([java.lang.String;])V
// int max(float x, float y)	-> (FF)I
type ConstantNameAndTypeInfo struct {
	nameIndex       uint16
	descriptorIndex uint16
}

func (c *ConstantNameAndTypeInfo) readInfo(reader *ClassReader) {
	c.nameIndex = reader.readUint16()
	c.descriptorIndex = reader.readUint16()
}
