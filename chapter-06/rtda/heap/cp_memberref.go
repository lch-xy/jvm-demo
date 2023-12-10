package heap

import "jvm-demo/chapter-06/classfile"

// 在Java中，我们并不能在同一个类中定义名字相同但类型不同的两个字段，这只是虚拟机进行了限制，底层是支持的
// MemberRef结构体来存放字段和方法符号引用共有的信息
type MemberRef struct {
	SymRef
	name       string
	descriptor string
}

// 从class文件内存储的字段或方法常量中提取数据
func (self *MemberRef) copyMemberRefInfo(refInfo *classfile.ConstantMemberrefInfo) {
	self.className = refInfo.ClassName()
	self.name, self.descriptor = refInfo.NameAndDescriptor()
}

func (self *MemberRef) Name() string {
	return self.name
}
func (self *MemberRef) Descriptor() string {
	return self.descriptor
}
