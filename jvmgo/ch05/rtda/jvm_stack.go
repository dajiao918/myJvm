package rtda

/*Stack 线程栈*/
type Stack struct {
	maxSize int
	size    int
	_top    *Frame
}

func newStack(maxSize int) *Stack {
	return &Stack{
		maxSize: maxSize,
	}
}

func (self *Stack) push(frame *Frame) {
	if self.size == self.maxSize {
		panic("java.lang.StackOverflowError")
	}
	if self._top != nil {
		self._top.lower = frame
	}
	self._top = frame
	self.size++
}

func (self *Stack) pop() *Frame {
	if self._top != nil {
		frame := self._top
		self._top = self._top.lower
		frame.lower = nil
		self.size--
		return frame
	}
	panic("jvm stack is empty")
}

func (self *Stack) top() *Frame {
	return self._top
}
