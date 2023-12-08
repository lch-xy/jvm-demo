package main

import (
	"fmt"
	"jvm-demo/chapter-05/classfile"
	"jvm-demo/chapter-05/instructions"
	"jvm-demo/chapter-05/instructions/base"
	"jvm-demo/chapter-05/rtda"
)

func interpret(methodInfo *classfile.MemberInfo) {
	codeAttr := methodInfo.CodeAttribute()
	maxLocals := codeAttr.MaxLocals()
	maxStack := codeAttr.MaxStack()
	bytecode := codeAttr.Code()

	thread := rtda.NewThread()
	frame := thread.NewFrame(maxLocals, maxStack)
	thread.PushFrame(frame)

	defer catchErr(frame)
	loop(thread, bytecode)
}

func catchErr(frame *rtda.Frame) {
	if r := recover(); r != nil {
		fmt.Printf("LocalVars:%v\n", frame.LocalVars())
		fmt.Printf("OperandStack:%v\n", frame.OperandStack())
		panic(r)
	}
}

// loop()函数循环执行“计算pc、解码指令、执行指令”这三个步骤，直到遇到错误！
func loop(thread *rtda.Thread, bytecode []byte) {
	frame := thread.PopFrame()
	reader := &base.BytecodeReader{}

	// 首先获取当前 PC，然后使用 BytecodeReader 读取字节码并解码为具体的指令。
	// 接着，执行指令的 FetchOperands 方法用于获取指令操作数，然后更新栈帧的下一条指令的 PC。
	// 最后，打印当前 PC、指令类型和具体指令，并调用指令的 Execute 方法执行指令。
	for {
		pc := frame.NextPC()
		thread.SetPC(pc)

		// decode
		reader.Reset(bytecode, pc)
		opcode := reader.ReadUint8()
		inst := instructions.NewInstruction(opcode)
		inst.FetchOperands(reader)
		frame.SetNextPC(reader.PC())

		// execute
		fmt.Printf("pc:%2d inst:%T %v\n", pc, inst, inst)
		inst.Execute(frame)
	}
}
