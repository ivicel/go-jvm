package classfile

// 行号属性
//
// LineNumberTable_attribute {
// 	u2 attribute_name_index;
// 	u2 attribute_length;
// 	u2 line_number_table_length;
// 	{
// 		u2 start_pc;
// 		u2 line_number;
// 	} line_number_table[line_number_table_length];
// }
type LineNumberTableAttribute struct {
	lineNumberTable []*LineNumberTableEntry
}

func (attr *LineNumberTableAttribute) readInfo(reader *ClassReader) {
	lineNumberTableLength := reader.readUint16()
	attr.lineNumberTable = make([]*LineNumberTableEntry, lineNumberTableLength)
	for i := 0; i < int(lineNumberTableLength); i++ {
		attr.lineNumberTable[i] = &LineNumberTableEntry{
			startPc:    reader.readUint16(),
			lineNumber: reader.readUint16(),
		}
	}
}

type LineNumberTableEntry struct {
	startPc    uint16
	lineNumber uint16
}

func (attr *LineNumberTableEntry) String() string {
	return "LineNumberTable"
}
