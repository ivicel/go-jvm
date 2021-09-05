package classfile

import "fmt"

// 常量结构
// cp_info {
// 	u1 tag;
// 	u1 info[];
// }
const (
	CONSTANT_Class              = 7  // class 类型
	CONSTANT_Fieldref           = 9  // 字段
	CONSTANT_Methodref          = 10 // 方法
	CONSTANT_InterfaceMethodref = 11 // 接口
	CONSTANT_String             = 8  // 字符串
	CONSTANT_Integer            = 3  // 整型
	CONSTANT_Float              = 4  // 单精度浮点
	CONSTANT_Long               = 5  // 长整型
	CONSTANT_Double             = 6  // 双精度浮点
	CONSTANT_NameAndType        = 12 // 名称及描述符索引结构
	CONSTANT_Utf8               = 1  // utf8 字符
	CONSTANT_MethodHandle       = 15
	CONSTANT_MethodType         = 16
	CONSTANT_InvokeDynamic      = 18
)

// 常量信息
type ConstantInfo interface {
	readInfo(reader *ClassReader)
}

// TODO: readConstantInfo 读一个常量结构
func readConstantInfo(reader *ClassReader, cp ConstantPool) ConstantInfo {
	tag := reader.readUint8()
	switch tag {
	case CONSTANT_Class:
		return &ConstantClassInfo{}
	case CONSTANT_Fieldref:
		return &ConstantFieldrefInfo{}
	case CONSTANT_Methodref:
		return &ConstantMethodrefInfo{}
	case CONSTANT_InterfaceMethodref:
		return &ConstantInterfaceMethodrefInfo{}
	case CONSTANT_String:
		return &ConstantStringInfo{}
	case CONSTANT_Integer:
		return &ConstantIntegerInfo{}
	case CONSTANT_Float:
		return &ConstantFloatInfo{}
	case CONSTANT_Long:
		return &ConstantLongInfo{}
	case CONSTANT_Double:
		return &ConstantDoubleInfo{}
	case CONSTANT_NameAndType:
		return &ConstantNameAndTypeInfo{}
	case CONSTANT_Utf8:
		return &ConstantUtf8Info{}
	case CONSTANT_MethodHandle:
		return &ConstantMethodHandleInfo{}
	case CONSTANT_MethodType:
		return &ConstantMethodTypeInfo{}
	case CONSTANT_InvokeDynamic:
		return &ConstantInvokeDynamicInfo{}
	default:
		panic(fmt.Errorf("java.lang.ClassFormatError: constant pool tag[%d]!", tag))
	}
}
