package references

import (
	"jvm-demo/chapter-06/instructions/base"
	"jvm-demo/chapter-06/rtda"
	"jvm-demo/chapter-06/rtda/heap"
)

// checkcast指令和instanceof指令很像，区别在于：
// instanceof指令会改变操作数栈（弹出对象引用，推入判断结果）
// checkcast则不改变操作数栈（如果判断失败，直接抛出ClassCastException异常）

type CHECK_CAST struct{ base.Index16Instruction }

func (self *CHECK_CAST) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	// 先从操作数栈中弹出对象引用，再推回去，这样就不会改变操作数栈的状态。如果引用是null，则指令执行结束。
	ref := stack.PopRef()
	stack.PushRef(ref)
	if ref == nil {
		return
	}

	cp := frame.Method().Class().ConstantPool()
	classRef := cp.GetConstant(self.Index).(*heap.ClassRef)
	class := classRef.ResolvedClass()
	if !ref.IsInstanceOf(class) {
		panic("java.lang.ClassCastException")
	}
}
