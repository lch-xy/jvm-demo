package classfile

/*
	// SourceFile是可选定长属性，只会出现在ClassFile结构中，用于指出源文件名。
	SourceFile_attribute {
		u2 attribute_name_index; // 在这里这个值通常是指向 "SourceFile" 字符串的索引。
		u4 attribute_length; // attribute_length的值必须是2。
		u2 sourcefile_index; // sourcefile_index是常量池索引，指向CONSTANT_Utf8_info常量
	}
*/

// 主要用于指定编译后的字节码文件对应的源文件的名称。
// 这个属性对于调试和诊断非常有用，因为它能够帮助开发者追踪字节码文件与源代码的对应关系。
type SourceFileAttribute struct {
	cp              ConstantPool
	sourceFileIndex uint16
}

func (self *SourceFileAttribute) readInfo(reader *ClassReader) {
	self.sourceFileIndex = reader.readUint16()
}
func (self *SourceFileAttribute) FileName() string {
	return self.cp.getUtf8(self.sourceFileIndex)
}
