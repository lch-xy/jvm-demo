package classfile

/*
	// Exceptions是变长属性，记录方法抛出的异常表，其结构定义如下:
	Exceptions_attribute {
		u2 attribute_name_index; // 在这里这个值通常是指向 "Exceptions" 字符串的索引。
		u4 attribute_length; // 属性长度
		u2 number_of_exceptions; // 异常数量
		u2 exception_index_table[number_of_exceptions]; // 表示异常类在常量池中的索引。
	}
*/

type ExceptionsAttribute struct {
	exceptionIndexTable []uint16
}

func (self *ExceptionsAttribute) readInfo(reader *ClassReader) {
	self.exceptionIndexTable = reader.readUint16s()
}
func (self *ExceptionsAttribute) ExceptionIndexTable() []uint16 {
	return self.exceptionIndexTable
}
