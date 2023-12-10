package classfile

/*
	CONSTANT_NameAndType_info {
		u1 tag;
		u2 name_index; // 字段或方法名的索引
		u2 descriptor_index; // 字段或方法的描述符的索引
	}
*/

// Java语言支持方法重载（override），不同的方法可以有相同的名字，只要参数列表不同即可。
// 这就是为什么CONSTANT_NameAndType_info结构要同时包含名称和描述符的原因。那么字段呢？
// Java是不能定义多个同名字段的，哪怕它们的类型各不相同。这只是Java语法的限制而已，从class文件的层面来看，是完全可以支持这点的。

// CONSTANT_Class_info和CONSTANT_NameAndType_info加在一起可以唯一确定一个字段或者方法。
type ConstantNameAndTypeInfo struct {
	nameIndex       uint16
	descriptorIndex uint16
}

func (self *ConstantNameAndTypeInfo) readInfo(reader *ClassReader) {
	self.nameIndex = reader.readUint16()
	self.descriptorIndex = reader.readUint16()
}
