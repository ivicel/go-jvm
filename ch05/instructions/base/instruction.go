package base

import "go-jvm/ch05/rtda"

/**
 * 常量指令: constants
 * 加载指令: load
 * 存储指令: store
 * 操作数栈指令: stack
 * 数学指令: math
 * 转换指令: conversions
 * 比较指令: comparisions 指令
 * 控制指令: control
 * 引用指令: references
 * 扩展指令: extended
 * 保留指令: reserved
 *
 * 保留指令只有 3 条: breakpoint(0xCA) 用于断点调试, impdepl(0xFE), impdepl(0xFF) 虚拟机内部使用
 */

// 操作数指令
type Instruction interface {
	// FetchOperands 从字节码中提取操作数
	FetchOperands(reader *BytecodeReader)

	// Execute 执行指令逻辑
	Execute(frame *rtda.Frame)
}

// NoOperandsInstruction 表示没有操作数的指令
type NoOperandsInstruction struct{}

// FetchOperands 什么也不做
func (s *NoOperandsInstruction) FetchOperands(reader *BytecodeReader) {}

// 表示跳转指令
type BranchInstruction struct {
	Offset int // 存放跳转偏移量
}

func (s *BranchInstruction) FetchOperands(reader *BytecodeReader) {
	s.Offset = int(reader.ReadInt16())
}

// 存储和加载类指令需要根据索引存取局部变量, 索引由单字节操作数给出
type Index8Instruction struct {
	Index uint // 局部变量表索引
}

func (s *Index8Instruction) FetchOperands(reader *BytecodeReader) {
	s.Index = uint(reader.ReadUint8())
}

type Index16Instruction struct {
	Index uint
}

func (s *Index16Instruction) FetchOperands(reader *BytecodeReader) {
	s.Index = uint(reader.ReadUint16())
}
