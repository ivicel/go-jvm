package classfile

import "fmt"

const JAVA_MAGIC uint32 = 0xCAFEBABE

// class 文件结构
//
//	ClassFile {
//		u4 magic;												// 魔法值
//		u4 minor_version;										// 次版本
//		u4 major_version;										// 主版本
//		u2 constant_pool_length;								// 常量池大小, 个数为 n, 下标从 1 开始
//		cp_info constant_pool[constant_pool_length-1];			// 常量池结构数组, 数量至多 n-1 个, 但 long, double 类型算两个
//		u2 access_flags;										// 类访问控制
//		u2 this_class;											// 类名称下标
//		u2 super_class;											// 父类名称下标
//		u2 interface_count;										// 接口数量
//		u2 interfaces[interface_count];							// 接口结构数组
//		u2 fields_count;										// 字段数量
//		field_info fields[fields_count];						// 字段结构数组
//		u2 methods_count;										// 方法数量
//		method_info methods[methods_count];						// 方法结构数组
//		u2 attribute_count;										// 属性数量
//		attribute_info attributes[attribute_count];				// 属性数组
//	}
type ClassFile struct {
	magic        uint32          // 魔法常数
	minorVersion uint16          // 次版本号
	majorVersion uint16          // 主版本号
	constantPool ConstantPool    // 常量池
	accessFlag   uint16          // class 文件的访问权限
	thisClass    uint16          // 类文件名指向
	superClass   uint16          // 父类指向
	interfaces   []uint16        // 指向实现的接口名称
	fields       []*MemberInfo   // 类字段
	methods      []*MemberInfo   // 方法
	attributes   []AttributeInfo // 属性信息
}

// Parse 将 class 文件解析成 ClassFile 结构体
func Parse(classData []byte) (cf *ClassFile, err error) {
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("解析 class 文件出错: %v", r)
			}
		}
	}()

	cr := &ClassReader{classData}
	cf = &ClassFile{}
	cf.read(cr)
	return
}

func (c *ClassFile) read(r *ClassReader) {
	c.readAndCheckMagic(r)
	c.readAndCheckVersion(r)
	c.constantPool = readConstantPool(r)
	c.accessFlag = r.readUint16()
	c.thisClass = r.readUint16()
	c.superClass = r.readUint16()
	c.interfaces = r.readUint16s()
	c.fields = readMembers(r, c.constantPool)
	c.methods = readMembers(r, c.constantPool)

}

// readAndCheckMagic 检查 class 文件的魔法值
func (c *ClassFile) readAndCheckMagic(r *ClassReader) {
	c.magic = r.readUint32()
	if c.magic != JAVA_MAGIC {
		panic("java.lang.ClassFormatError: magic!")
	}
}

// readAndCheckVersion 检查 class 文件版本号
func (c *ClassFile) readAndCheckVersion(r *ClassReader) {
	c.minorVersion = r.readUint16()
	c.majorVersion = r.readUint16()

	// TODO: 只支持 java 8
	switch c.majorVersion {
	case 45:
		return
	case 46, 47, 48, 49, 50, 51, 52:
		// if c.minorVersion == 0 {
		// return
		// }
		return
	}

	panic("java.lang.UnsupportedClassVersionError!")
}

// MinorVersion 获取次版本号
func (c *ClassFile) MinorVersion() uint16 {
	return c.minorVersion
}

// MajorVersion 获取主版本号
func (c *ClassFile) MajorVersion() uint16 {
	return c.majorVersion
}

// ConstantPool 获取常量池
func (c *ClassFile) ConstantPool() ConstantPool {
	return c.constantPool
}

// AccessFlag 获取类/接口访问权限
func (c *ClassFile) AccessFlag() uint16 {
	return c.accessFlag
}

// Fields 获取类成员变量
func (c *ClassFile) Fields() []*MemberInfo {
	return c.fields
}

// Methods 获取成员方法
func (c *ClassFile) Methods() []*MemberInfo {
	return c.methods
}

// ClassName 获取类名
func (c *ClassFile) ClassName() string {
	return c.constantPool.getClassName(c.thisClass)
}

// SuperClassName 获取父类名
func (c *ClassFile) SuperClassName() string {
	if c.superClass > 0 {
		return c.constantPool.getClassName(c.superClass)
	}

	// 超类为 java.lang.Object
	return ""
}

// InterfaceNames 获取实现的接口名称
func (c *ClassFile) InterfaceNames() []string {
	interfaceNames := make([]string, len(c.interfaces))
	for i, cpIndex := range c.interfaces {
		interfaceNames[i] = c.constantPool.getClassName(cpIndex)
	}

	return interfaceNames
}

func (c *ClassFile) Attributes() []AttributeInfo {
	return c.attributes
}
