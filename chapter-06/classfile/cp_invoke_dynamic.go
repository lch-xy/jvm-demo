package classfile

/*
	// 用于表示方法句柄的一种常量。
	CONSTANT_MethodHandle_info {
		u1 tag;
		u1 reference_kind; // 表示方法句柄的类型，它指定了如何解释 reference_index 字段的内容。
		u2 reference_index; // 用于指向常量池中的一个方法句柄的索引。
	}
*/
// 方法句柄被用于支持动态语言、Lambda 表达式等特性。
// 它提供了一种灵活的方式，通过方法句柄可以在运行时动态地调用方法，而不是使用传统的静态方法调用。
type ConstantMethodHandleInfo struct {
	referenceKind  uint8
	referenceIndex uint16
}

func (self *ConstantMethodHandleInfo) readInfo(reader *ClassReader) {
	self.referenceKind = reader.readUint8()
	self.referenceIndex = reader.readUint16()
}

/*
	// 常量类型，用于表示方法类型（Method Type）。
	CONSTANT_MethodType_info {
	    u1 tag;
	    u2 descriptor_index; // 指向常量池中的一个 UTF-8 编码的方法类型描述符的索引。
	}
*/

// 这个常量主要用于支持 Java 7 引入的 invokedynamic 指令和动态语言功能。
// 在这些场景下，运行时需要动态地获取方法的参数类型和返回类型，而 CONSTANT_MethodType_info 提供了一种标准的方式来表示这样的方法类型信息。
// 这样的方法类型描述符在常量池中的索引，可以在运行时动态地获取方法类型信息。
type ConstantMethodTypeInfo struct {
	descriptorIndex uint16
}

func (self *ConstantMethodTypeInfo) readInfo(reader *ClassReader) {
	self.descriptorIndex = reader.readUint16()
}

/*
// 用于表示动态方法调用（invokedynamic）的一种常量类型。

	CONSTANT_InvokeDynamic_info {
	    u1 tag;
	    u2 bootstrap_method_attr_index; // 表示引导方法（bootstrap method）在 BootstrapMethods 属性表中的索引。
	    u2 name_and_type_index; // 索引，其中包含了调用方法的名称和类型。
	}
*/
// invokedynamic 指令是 Java 7 引入的一种新的字节码指令，主要用于支持在运行时绑定的动态语言特性。
// 与传统的方法调用不同，invokedynamic 的目标在运行时由引导方法确定，这使得 Java 虚拟机更加灵活，能够支持更动态的语言特性。
type ConstantInvokeDynamicInfo struct {
	bootstrapMethodAttrIndex uint16
	nameAndTypeIndex         uint16
}

func (self *ConstantInvokeDynamicInfo) readInfo(reader *ClassReader) {
	self.bootstrapMethodAttrIndex = reader.readUint16()
	self.nameAndTypeIndex = reader.readUint16()
}
