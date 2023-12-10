package rtda

import "jvm-demo/chapter-06/rtda/heap"

type Slot struct {
	num int32        // 用于存储基本类型变量
	ref *heap.Object // 用于存储引用类型变量
}
