package classfile

//	LineNumberTable_attribute {
//		u2 attribute_name_index;
//		u4 attribute_length;
//		u2 line_number_table_length;
//		{
//			u2 start_pc;
//			u2 line_number;
//		}
//		line_number_table[line_number_table_length];
//	}

type LocalVariableTableAttribute struct {
	lineNumberTable []*LocalVariableTableEntry
}
type LocalVariableTableEntry struct {
	startPc    uint16
	lineNumber uint16
}

func (self *LocalVariableTableAttribute) readInfo(reader *ClassReader) {
	lineNumberTableLength := reader.readUint16()
	self.lineNumberTable = make([]*LocalVariableTableEntry, lineNumberTableLength)
	for i := range self.lineNumberTable {
		self.lineNumberTable[i] = &LocalVariableTableEntry{
			startPc:    reader.readUint16(),
			lineNumber: reader.readUint16(),
		}
	}
}
