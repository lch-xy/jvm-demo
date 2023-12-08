package classfile

/*
	// 用于表示字段的常量值的一种属性结构。它包含了与字段相关的常量初始值。
	// ConstantValue是定长属性，只会出现在field_info结构中，用于表示常量表达式的值
	ConstantValue_attribute {
		u2 attribute_name_index; // 指向常量池中的 "ConstantValue" 字符串的索引。
		u4 attribute_length; // attribute_length的值必须是2。因为后面就一个constantvalue_index
		u2 constantvalue_index; // 常量池索引，但具体指向哪种常量因字段类型而异。
	}
*/

type ConstantValueAttribute struct {
	constantValueIndex uint16
}

func (self *ConstantValueAttribute) readInfo(reader *ClassReader) {
	self.constantValueIndex = reader.readUint16()
}
func (self *ConstantValueAttribute) ConstantValueIndex() uint16 {
	return self.constantValueIndex
}
