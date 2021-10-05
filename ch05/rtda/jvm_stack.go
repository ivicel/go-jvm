package rtda

/**
 * 栈内元素指针结构指向
 *
 * ┌─────────┐   ┌─────────┐     ┌─────────┐
 * │ Stack   │ ┌►│ Stack   │  ┌─►│ Stack   │
 * ├─────────┤ │ ├─────────┤  │  ├─────────┤
 * │  _top   ├─┘ │  lower  ├──┘  │  lower  │
 * └─────────┘   └─────────┘     └─────────┘
 *
 */

// jvm 栈结构
type Stack struct {
	maxSize uint   // 栈的最大深度
	size    uint   // 当前栈的大小
	_top    *Frame // 指向栈顶
}

// newStack 生成新的栈
func newStack(maxSize uint) *Stack {
	return &Stack{
		maxSize: maxSize,
	}
}

// push 入栈
func (s *Stack) push(frame *Frame) {
	// 超出栈的深度
	if s.size >= s.maxSize {
		panic("java.lang.StackOverflowError")
	}

	if s._top != nil {
		frame.lower = s._top
	}

	s._top = frame
	s.size++
}

// pop 出栈
func (s *Stack) pop() *Frame {
	if s._top == nil {
		panic("jvm stack is empty!")
	}

	top := s._top
	s._top = top.lower
	// 将这个帧的内存回收
	top.lower = nil
	return top
}

// top 返回入栈顶
func (s *Stack) top() *Frame {
	if s._top == nil {
		panic("jvm stack is empty!")
	}

	return s._top
}
