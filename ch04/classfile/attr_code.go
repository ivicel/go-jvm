package classfile

// Code_attribute {
// 	u2 attribute_name_index;
// 	u4 attribute_length;
// 	u2 max_stack;									// 操作数栈的最大深度
// 	u2 max_locals;									// 局部变量表大小
// 	u4 code_length;									// 字节码长度
// 	u1 code[code_length];							// 字节码
// 	u2 exception_table_length;						// 异常表长度
//
// 	{
// 		u2 start_pc;
// 		u2 end_pc;
// 		u2 handler_pc;
// 		u2 catch_type;
// 	} exception_table[exception_table_length];		// 异常表结构
//
// 	u2 attributes_count;							// 方法属性下还有其他属性, 这些属性长度
// 	attribute_info attributes[attributes_count];
// }
type CodeAttribute struct {
	cp             ConstantPool
	maxStack       uint16                 // 操作数栈深度
	maxLocals      uint16                 // 局部变量表大小
	code           []byte                 // 字节码
	exceptionTable []*ExceptionTableEntry // 异常表
	attributes     []AttributeInfo        // 其他属性
}

func (attr *CodeAttribute) readInfo(reader *ClassReader) {
	attr.maxStack = reader.readUint16()
	attr.maxLocals = reader.readUint16()
	attr.code = reader.readBytes(uint32(reader.readUint32()))
	attr.exceptionTable = readExceptionTable(reader)
	attr.attributes = readAttributes(reader, attr.cp)
}

type ExceptionTableEntry struct {
	startPc   uint16
	endPc     uint16
	handlerPc uint16
	catchType uint16
}

func readExceptionTable(reader *ClassReader) []*ExceptionTableEntry {
	exceptionTableLength := reader.readUint16()
	exceptionTable := make([]*ExceptionTableEntry, exceptionTableLength)
	for i := 0; i < int(exceptionTableLength); i++ {
		exceptionTable[i] = &ExceptionTableEntry{
			startPc:   reader.readUint16(),
			endPc:     reader.readUint16(),
			handlerPc: reader.readUint16(),
			catchType: reader.readUint16(),
		}
	}

	return exceptionTable
}

func (attr *CodeAttribute) String() string {
	return "Code"
}
