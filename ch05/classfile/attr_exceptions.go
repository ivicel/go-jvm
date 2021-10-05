package classfile

// Exceptions 属性
//
// Exceptions_attribute {
// 	u2 attribute_name_index;
// 	u4 attribute_length;
// 	u2 number_of_exceptions;
// 	u2 exception_index_table[number_of_exceptions];
// }
type ExceptionsAttribute struct {
	exceptioniIndexTable []uint16
}

func (attr *ExceptionsAttribute) readInfo(reader *ClassReader) {
	attr.exceptioniIndexTable = reader.readUint16s()
}

func (attr *ExceptionsAttribute) ExceptionIndexTable() []uint16 {
	return attr.exceptioniIndexTable
}

func (attr *ExceptionsAttribute) String() string {
	return "Exceptions"
}
