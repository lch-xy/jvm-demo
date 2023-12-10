package classfile

// 一个通用的结构，用于表示在 Java 字节码中出现的未解析（或未识别）的属性。
// 由于 Java 字节码规范允许在类文件的属性表中包含用户自定义的属性，而这些属性可能是新版本的 Java 编译器或其他工具所特有的。
// 因此不同的工具可能会在类文件中添加自定义的属性，而这些属性对于旧版本的 Java 编译器或解释器可能是未知的。
type UnparsedAttribute struct {
	name   string // 未解析属性的名称
	length uint32 // 属性的长度
	info   []byte // 未解析的属性信息（原始字节码）
}

func (self *UnparsedAttribute) readInfo(reader *ClassReader) {
	self.info = reader.readBytes(self.length)
}
