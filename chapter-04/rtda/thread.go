package rtda

// Java虚拟机规范对Java虚拟机栈的约束也相当宽松。Java虚拟机栈可以是连续的空间，也可以不连续；可以是固定大小，也可以在运行时动态扩展

type Thread struct {
	pc    int
	stack *Stack
}

func NewThread() *Thread {
	return &Thread{
		// 参数表示要创建的Stack最多可以容纳多少帧
		stack: newStack(1024),
	}
}

func (self *Thread) PushFrame(frame *Frame) {
	self.stack.push(frame)
}
func (self *Thread) PopFrame() *Frame {
	return self.stack.pop()
}

func (self *Thread) CurrentFrame() *Frame {
	return self.stack.top()
}

func (self *Thread) PC() int      { return self.pc } // getter
func (self *Thread) SetPC(pc int) { self.pc = pc }   // setter
