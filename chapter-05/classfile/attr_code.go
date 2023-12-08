package classfile

/*
// Code是变长属性，只存在于method_info结构中。Code属性中存放字节码等方法相关信息。

	Code_attribute {
		u2 attribute_name_index; // 指向常量池中的一个 UTF-8 编码的字符串，表示属性的名称。对于 Code_attribute，这个值通常是指向 "Code" 字符串的索引。
		u4 attribute_length;
		u2 max_stack; // 操作数栈的最大深度
		u2 max_locals; // 局部变量表大小
		u4 code_length;
		u1 code[code_length]; // 字节码，存储着具体的字节码指令。
		u2 exception_table_length;
		{
			u2 start_pc;
			u2 end_pc;
			u2 handler_pc; // 表示异常处理代码的起始位置（字节码偏移量）。
			u2 catch_type; // 表示异常类的索引。它应该是一个 CONSTANT_Class_info 常量，表示被捕获的异常类型。
		}
		exception_table[exception_table_length]; // 字段是一个数组，包含了异常处理表的信息。
		u2 attributes_count;
		attribute_info attributes[attributes_count]; // 包含了该属性的附加属性信息。
	}
*/
type CodeAttribute struct {
	cp             ConstantPool
	maxStack       uint16
	maxLocals      uint16
	code           []byte
	exceptionTable []*ExceptionTableEntry
	attributes     []AttributeInfo
}

type ExceptionTableEntry struct {
	startPc   uint16
	endPc     uint16
	handlerPc uint16
	catchType uint16
}

func (self *CodeAttribute) readInfo(reader *ClassReader) {
	self.maxStack = reader.readUint16()
	self.maxLocals = reader.readUint16()
	codeLength := reader.readUint32()
	self.code = reader.readBytes(codeLength)
	self.exceptionTable = readExceptionTable(reader)
	self.attributes = readAttributes(reader, self.cp)
}

func readExceptionTable(reader *ClassReader) []*ExceptionTableEntry {
	exceptionTableLength := reader.readUint16()
	exceptionTable := make([]*ExceptionTableEntry, exceptionTableLength)
	for i := range exceptionTable {
		exceptionTable[i] = &ExceptionTableEntry{
			startPc:   reader.readUint16(),
			endPc:     reader.readUint16(),
			handlerPc: reader.readUint16(),
			catchType: reader.readUint16(),
		}
	}
	return exceptionTable
}

func (self *CodeAttribute) MaxStack() uint {
	return uint(self.maxStack)
}
func (self *CodeAttribute) MaxLocals() uint {
	return uint(self.maxLocals)
}
func (self *CodeAttribute) Code() []byte {
	return self.code
}
func (self *CodeAttribute) ExceptionTable() []*ExceptionTableEntry {
	return self.exceptionTable
}

func (self *ExceptionTableEntry) StartPc() uint16 {
	return self.startPc
}
func (self *ExceptionTableEntry) EndPc() uint16 {
	return self.endPc
}
func (self *ExceptionTableEntry) HandlerPc() uint16 {
	return self.handlerPc
}
func (self *ExceptionTableEntry) CatchType() uint16 {
	return self.catchType
}
