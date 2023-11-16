package classfile

// MUTF-8编码方式和UTF-8大致相同，但并不兼容。
// 差别有两点：
//  一是null字符（代码点U+0000）会被编码成2字节：0xC0、0x80；
//  二是补充字符（SupplementaryCharacters，代码点大于U+FFFF的Unicode字符）是按UTF-16拆分为代理对（Surrogate Pair）分别编码的。

type ConstantUtf8Info struct {
	str string
}

func (self *ConstantUtf8Info) readInfo(reader *ClassReader) {
	length := uint32(reader.readUint16())
	bytes := reader.readBytes(length)
	self.str = decodeMUTF8(bytes)
}

// Java序列化机制也使用了MUTF-8编码。
// java.io.DataInput和java.io.DataOutput接口分别定义了readUTF()和writeUTF()方法，可以读写MUTF-8编码的字符串。
// decodeMUTF8()函数的代码就是笔者根据java.io.DataInputStream.readUTF()方法改写的。
// 代码很长，解释起来也很乏味，所以这里就不详细解释了。
// 因为Go语言字符串使用UTF-8编码，所以如果字符串中不包含null字符或补充字符，下面这个简化版的readMUTF8()也是可以工作的。
// 简化版，完整版请阅读本章源代码
func decodeMUTF8(bytes []byte) string {
	return string(bytes)
}
