package classfile

// 属性 Deprecated
type DeprecatedAttribute struct {
	MarkerAttribute
}

func (attr *DeprecatedAttribute) String() string {
	return "Deprecated"
}

// 属性 Synthetic
type SyntheticAttribute struct {
	MarkerAttribute
}

func (attr *SyntheticAttribute) String() string {
	return "Synthetic"
}

// 这只是标记属性, 没有数据
type MarkerAttribute struct{}

func (m *MarkerAttribute) readInfo(reader *ClassReader) {
	// read nothing
}
