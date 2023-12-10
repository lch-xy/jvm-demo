package heap

// 类、字段、方法和接口方法的符号引用，这4种类型的符号引用有一些共性，所以仍然使用继承来减少重复代码。
type SymRef struct {
	cp        *ConstantPool // 存放符号引用所在的运行时常量池指针，这样就可以通过符号引用访问到运行时常量池，进一步又可以访问到类数据。
	className string        // 存放类的完全限定名。
	class     *Class        // 缓存解析后的类结构体指针，这样类符号引用只需要解析一次就可以了，后续可以直接使用缓存值。
}

// 直接返回类指针，否则调用resolveClassRef()方法进行解析
func (self *SymRef) ResolvedClass() *Class {
	if self.class == nil {
		self.resolveClassRef()
	}
	return self.class
}

// jvms8 5.4.3.1
// 如果类D通过符号引用N引用类C的话，要解析N，先用D的类加载器加载C，然后检查D是否有权限访问C。
// 如果没有，则抛出IllegalAccessError异常。用D的类
func (self *SymRef) resolveClassRef() {
	d := self.cp.class
	c := d.loader.LoadClass(self.className)
	if !c.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}

	self.class = c
}
