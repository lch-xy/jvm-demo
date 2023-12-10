package classfile

/*
	// LocalVariableTable属性表中存放方法的局部变量信息。
	LocalVariableTable_attribute {
		u2 attribute_name_index; // 在这里这个值通常是指向 "LocalVariableTable" 字符串的索引。
		u4 attribute_length;
		u2 local_variable_table_length;
		{
			u2 start_pc; // 局部变量的起始字节码偏移量。
			u2 length; // 局部变量的作用域范围，即在字节码中的长度。
			u2 name_index; // 局部变量的名称在常量池中的索引。
			u2 descriptor_index; // 局部变量的描述符在常量池中的索引。
			u2 index; // 局部变量在局部变量表中的索引。
		}
		local_variable_table[local_variable_table_length];
	}
*/
// 属于调试信息，都不是运行时必需的。
// 为调试器提供局部变量的信息，以便在调试时能够查看局部变量的名称和值。
// 这对于调试复杂的 Java 程序非常有帮助。这个属性通常在编译 Java 源代码时由编译器生成。
type LocalVariableTableAttribute struct {
	localVariableTable []*LocalVariableTableEntry
}

type LocalVariableTableEntry struct {
	startPc         uint16
	length          uint16
	nameIndex       uint16
	descriptorIndex uint16
	index           uint16
}

//func (self *LocalVariableTableAttribute) readInfo(reader *ClassReader) {
//	lineNumberTableLength := reader.readUint16()
//	self.lineNumberTable = make([]*LocalVariableTableEntry, lineNumberTableLength)
//	for i := range self.lineNumberTable {
//		self.lineNumberTable[i] = &LocalVariableTableEntry{
//			startPc:    reader.readUint16(),
//			lineNumber: reader.readUint16(),
//		}
//	}
//}

func (self *LocalVariableTableAttribute) readInfo(reader *ClassReader) {
	localVariableTableLength := reader.readUint16()
	self.localVariableTable = make([]*LocalVariableTableEntry, localVariableTableLength)
	for i := range self.localVariableTable {
		self.localVariableTable[i] = &LocalVariableTableEntry{
			startPc:         reader.readUint16(),
			length:          reader.readUint16(),
			nameIndex:       reader.readUint16(),
			descriptorIndex: reader.readUint16(),
			index:           reader.readUint16(),
		}
	}
}
