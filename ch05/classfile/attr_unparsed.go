package classfile

// 无法解析的属性结构
type UnparsedAttribute struct {
	name   string
	length uint32
	info   []byte
}

func (a *UnparsedAttribute) readInfo(reader *ClassReader) {
	a.info = reader.readBytes(a.length)
}

// 23 种预定义属性可以分为三组
// 1. java 虚拟机所必需的, 5 种
// 2. java 类库所必需的, 15 种
// 3. 提供给工具使用的, 6 种, 可选的. 比如 LineNumberTable 在异常堆栈中显示行号

func (attr *UnparsedAttribute) String() string {
	return "UnparsedAttribute"
}
