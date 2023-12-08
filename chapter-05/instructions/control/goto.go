package control

import (
	"jvm-demo/chapter-05/instructions/base"
	"jvm-demo/chapter-05/rtda"
)

// Branch always
type GOTO struct{ base.BranchInstruction }

func (self *GOTO) Execute(frame *rtda.Frame) {
	base.Branch(frame, self.Offset)
}
