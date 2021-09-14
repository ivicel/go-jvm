package rtda

// 线程结构表示
type Thread struct {
	PC    int    // 程序计数器, 指向下一行要执行的命令
	Stack *Stack // 线程虚拟机栈
}

// newThread 构建新的线程
func newThread() *Thread {
	return &Thread{
		Stack: newStack(1024),
	}
}

// PushFrame 压入的的栈帧
func (t *Thread) PushFrame(frame *Frame) {
	t.Stack.push(frame)
}

// PopFrame 将栈帧出栈
func (t *Thread) PopFrame() *Frame {
	return t.Stack.pop()
}

// CurrentFrame 获取当前栈顶帧
func (t *Thread) CurrentFrame() *Frame {
	return t.Stack.top()
}
