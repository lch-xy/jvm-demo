package rtda

// maxSize字段保存栈的容量（最多可以容纳多少帧）, size字段保存栈的当前大小，_top字段保存栈顶指针。
type Stack struct {
	maxSize uint
	size    uint
	_top    *Frame
}

func newStack(maxSize uint) *Stack {
	return &Stack{
		maxSize: maxSize,
	}
}

// 把帧推入栈顶
func (self *Stack) push(frame *Frame) {
	if self.size >= self.maxSize {
		panic("java.lang.StackOverflowError")
	}
	if self._top != nil {
		frame.lower = self._top
	}
	self._top = frame
	self.size++
}

// 把栈顶帧弹出
func (self *Stack) pop() *Frame {
	if self._top == nil {
		panic("jvm stack is empty! ")
	}

	top := self._top
	self._top = top.lower
	top.lower = nil
	self.size--

	return top
}

// 返回栈顶帧，但并不弹出
func (self *Stack) top() *Frame {
	if self._top == nil {
		panic("jvm stack is empty! ")
	}
	return self._top
}
