package math

import (
	"go-jvm/ch05/instructions/base"
	"go-jvm/ch05/rtda"
)

// 将变量表中的 int 数据加上某个常量, 再保存回该变量中
// 相当 i += 3
type IINC struct {
	Index uint  // 本地变量表索引
	Const int32 // 常量数值
}

func (s *IINC) FetchOperands(reader *base.BytecodeReader) {
	s.Index = uint(reader.ReadInt8())
	s.Const = int32(reader.ReadInt8())
}

func (s *IINC) Execute(frame *rtda.Frame) {
	localVars := frame.LocalVars()
	val := localVars.GetInt(s.Index)
	localVars.SetInt(s.Index, val+s.Const)
}
