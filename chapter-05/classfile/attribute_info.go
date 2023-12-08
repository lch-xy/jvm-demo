package classfile

/*
	// 字节码中表示类、字段、方法等各种元素的属性的通用结构。这个结构用于包含与元素相关的额外信息。
	attribute_info {
		u2 attribute_name_index; // 指向常量池中的一个 UTF-8 编码的字符串，表示属性的名称。
		u4 attribute_length; // 属性长度
		u1 info[attribute_length]; //数组，里面存储着具体的属性信息。
	}
*/

// 各种属性表达的信息也各不相同，因此无法用统一的结构来定义。所以这里定义一个底层的接口，需要由具体的属性实现。
type AttributeInfo interface {
	readInfo(reader *ClassReader)
}

func readAttributes(reader *ClassReader, cp ConstantPool) []AttributeInfo {
	// 直接读取attributesCount，这都是定义好的结构
	attributesCount := reader.readUint16()
	attributes := make([]AttributeInfo, attributesCount)
	for i := range attributes {
		attributes[i] = readAttribute(reader, cp)
	}
	return attributes
}

// readAttribute()先读取属性名索引，根据它从常量池中找到属性名。
// 然后读取属性长度，接着调用newAttributeInfo()函数创建具体的属性实例。
func readAttribute(reader *ClassReader, cp ConstantPool) AttributeInfo {
	attrNameIndex := reader.readUint16()
	attrName := cp.getUtf8(attrNameIndex)
	attrLen := reader.readUint32()
	attrInfo := newAttributeInfo(attrName, attrLen, cp)
	attrInfo.readInfo(reader)
	return attrInfo
}

func newAttributeInfo(attrName string, attrLen uint32,
	cp ConstantPool) AttributeInfo {
	switch attrName {

	case "Code":
		return &CodeAttribute{cp: cp}
	case "ConstantValue":
		return &ConstantValueAttribute{}
	case "Deprecated":
		return &DeprecatedAttribute{}
	case "Exceptions":
		return &ExceptionsAttribute{}
	case "LineNumberTable":
		return &LineNumberTableAttribute{}
	case "LocalVariableTable":
		return &LocalVariableTableAttribute{}
	case "SourceFile":
		return &SourceFileAttribute{cp: cp}
	case "Synthetic":
		return &SyntheticAttribute{}
	default:
		return &UnparsedAttribute{attrName, attrLen, nil}
	}
}
