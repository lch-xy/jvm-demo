package heap

import "jvm-demo/chapter-06/classfile"

// 为了避免重复代码，创建一个结构体存放这些公共的信息（访问标志、名字、描述符）
type ClassMember struct {
	accessFlags uint16
	name        string
	descriptor  string
	class       *Class // class字段存放Class结构体指针，这样可以通过字段或方法访问到它所属的类。
}

// 从class文件中复制数据
func (self *ClassMember) copyMemberInfo(memberInfo *classfile.MemberInfo) {
	self.accessFlags = memberInfo.AccessFlags()
	self.name = memberInfo.Name()
	self.descriptor = memberInfo.Descriptor()
}

// 用通俗的语言描述字段访问规则。如果字段是public，则任何类都可以访问。
// 如果字段是protected，则只有子类和同一个包下的类可以访问。
// 如果字段有默认访问权限（非public，非protected，也非privated），则只有同一个包下的类可以访问。
// 否则，字段是private，只有声明这个字段的类才能访问。
func (self *ClassMember) isAccessibleTo(d *Class) bool {
	if self.IsPublic() {
		return true
	}
	c := self.class
	if self.IsProtected() {
		return d == c || d.isSubClassOf(c) ||
			c.getPackageName() == d.getPackageName()
	}
	if !self.IsPrivate() {
		return c.getPackageName() == d.getPackageName()
	}
	return d == c
}

func (self *ClassMember) IsPublic() bool {
	return 0 != self.accessFlags&ACC_PUBLIC
}
func (self *ClassMember) IsPrivate() bool {
	return 0 != self.accessFlags&ACC_PRIVATE
}
func (self *ClassMember) IsProtected() bool {
	return 0 != self.accessFlags&ACC_PROTECTED
}
func (self *ClassMember) IsStatic() bool {
	return 0 != self.accessFlags&ACC_STATIC
}
func (self *ClassMember) IsFinal() bool {
	return 0 != self.accessFlags&ACC_FINAL
}
func (self *ClassMember) IsSynthetic() bool {
	return 0 != self.accessFlags&ACC_SYNTHETIC
}

// getters
func (self *ClassMember) Name() string {
	return self.name
}
func (self *ClassMember) Descriptor() string {
	return self.descriptor
}
func (self *ClassMember) Class() *Class {
	return self.class
}
