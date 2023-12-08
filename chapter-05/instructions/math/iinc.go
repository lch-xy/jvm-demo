package math

import (
	"jvm-demo/chapter-05/instructions/base"
	"jvm-demo/chapter-05/rtda"
)

// Increment local variable by constant
type IINC struct {
	Index uint  // 索引部分是一个无符号 8 位整数，用于指定要增加的本地变量的索引。
	Const int32 // 常量部分是一个有符号 8 位整数，表示要增加的常量值。
}

func (self *IINC) FetchOperands(reader *base.BytecodeReader) {
	self.Index = uint(reader.ReadUint8())
	self.Const = int32(reader.ReadInt8())
}

func (self *IINC) Execute(frame *rtda.Frame) {
	localVars := frame.LocalVars()
	val := localVars.GetInt(self.Index)
	val += self.Const
	localVars.SetInt(self.Index, val)
}
