package rtda

// jvm 帧结构
type Frame struct {
	lower        *Frame        // 指向栈的下一个帧
	localVars    LocalVars     // 本地变量表
	operandStack *OperandStack // 操作数栈
}

// NewFrame 生成新的栈帧
func NewFrame(maxLocals, maxStack uint) *Frame {
	return &Frame{
		localVars:    newLocalVars(maxLocals),
		operandStack: newOperandStack(maxStack),
	}
}

func (f *Frame) LocalVars() LocalVars {
	return f.localVars
}

func (f *Frame) OperandStack() *OperandStack {
	return f.operandStack
}
