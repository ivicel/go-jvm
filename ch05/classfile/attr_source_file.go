package classfile

// 用于指出源文件名
//
// SourceFile_attribute {
// 	u2 attribute_name_index; // 指向常量 SourceFile
// 	u2 attribute_length; // 名称长度
// 	u2 sourcefile_index; // 指向名称索引
// }

type SourceFileAttribute struct {
	cp              ConstantPool
	sourceFileIndex uint16
}

func (attr *SourceFileAttribute) readInfo(reader *ClassReader) {
	attr.sourceFileIndex = reader.readUint16()
}

func (attr *SourceFileAttribute) FileName() string {
	return attr.cp.getUtf8(attr.sourceFileIndex)
}

func (attr *SourceFileAttribute) String() string {
	return "SourceFile"
}
