package classfile

/*
	// LineNumberTable属性表存放方法的行号信息，
	LineNumberTable_attribute {
		u2 attribute_name_index; // 在这里这个值通常是指向 "LineNumberTable" 字符串的索引。
		u4 attribute_length;
		u2 line_number_table_length;
		{
			u2 start_pc; // 字节码的偏移量
			u2 line_number; // 源代码的行号。
		}
		line_number_table[line_number_table_length];
	}
*/

// 属于调试信息，都不是运行时必需的。
// 主要作用是在调试信息中提供源代码行号和字节码行号之间的映射，以便调试器可以在调试时准确地显示源代码的行数。
// 例如，当出现异常时，调试器可以利用这些信息在源代码中准确定位异常发生的位置。
type LineNumberTableAttribute struct {
	lineNumberTable []*LineNumberTableEntry
}
type LineNumberTableEntry struct {
	startPc    uint16
	lineNumber uint16
}

func (self *LineNumberTableAttribute) readInfo(reader *ClassReader) {
	lineNumberTableLength := reader.readUint16()
	self.lineNumberTable = make([]*LineNumberTableEntry, lineNumberTableLength)
	for i := range self.lineNumberTable {
		self.lineNumberTable[i] = &LineNumberTableEntry{
			startPc:    reader.readUint16(),
			lineNumber: reader.readUint16(),
		}
	}
}
