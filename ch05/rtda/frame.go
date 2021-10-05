package rtda

// jvm 帧结构
type Frame struct {
	lower        *Frame        // 指向栈的下一个帧
	localVars    LocalVars     // 本地变量表
	operandStack *OperandStack // 操作数栈
	thread       *Thread       // 当前线程
	nextPC       int           // 下一条指令计数器
}

// newFrame 生成新的栈帧
func newFrame(maxLocals, maxStack uint, thread *Thread) *Frame {
	return &Frame{
		localVars:    newLocalVars(maxLocals),
		operandStack: newOperandStack(maxStack),
		thread: thread,
	}
}

func (f *Frame) LocalVars() *LocalVars {
	return &f.localVars
}

func (f *Frame) OperandStack() *OperandStack {
	return f.operandStack
}

func (f *Frame) Thread() *Thread {
	return f.thread
}

func (f *Frame) NextPC() int {
	return f.nextPC
}

func (f *Frame) SetNextPC(nextPC int) {
	f.nextPC = nextPC
}
