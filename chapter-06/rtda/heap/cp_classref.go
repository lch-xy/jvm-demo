package heap

import "jvm-demo/chapter-06/classfile"

// 对于类符号引用，只要有类名，就可以解析符号引用。
type ClassRef struct {
	SymRef
}

// 根据class文件中存储的类常量创建ClassRef实例
func newClassRef(cp *ConstantPool, classInfo *classfile.ConstantClassInfo) *ClassRef {
	ref := &ClassRef{}
	ref.cp = cp
	ref.className = classInfo.Name()
	return ref
}
