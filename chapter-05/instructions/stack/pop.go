package stack

import (
	"jvm-demo/chapter-05/instructions/base"
	"jvm-demo/chapter-05/rtda"
)

type POP struct{ base.NoOperandsInstruction }

// pop指令只能用于弹出int、float等占用一个操作数栈位置的变量。
func (self *POP) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	stack.PopSlot()
}

type POP2 struct{ base.NoOperandsInstruction }

// double和long变量在操作数栈中占据两个位置，需要使用pop2指令弹出。
func (self *POP2) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	stack.PopSlot()
	stack.PopSlot()
}
