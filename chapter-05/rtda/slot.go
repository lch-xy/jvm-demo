package rtda

type Slot struct {
	num int32   // 用于存储基本类型变量
	ref *Object // 用于存储引用类型变量
}
