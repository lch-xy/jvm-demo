package constants

import (
	"jvm-demo/chapter-06/instructions/base"
	"jvm-demo/chapter-06/rtda"
)

/*
JVM 中的 nop 指令是空操作指令，它不执行任何操作，只是简单地跳过下一条指令。nop 指令的汇编代码为 nop，机器码为 00。
nop 指令可以用于以下目的：
	1.填充代码，以便代码长度是 8 字节的倍数，这样可以提高 CPU 的性能。
	2.跳过无效的代码，例如在调试时插入的断点。
	3.在程序中插入延迟，例如在等待 IO 操作完成时。
*/

type NOP struct{ base.NoOperandsInstruction }

func (self *NOP) Execute(frame *rtda.Frame) {
	// 什么也不用做
}
