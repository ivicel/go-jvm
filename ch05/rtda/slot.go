package rtda

// jvm 局部变量槽结构体
type Slot struct {
	num int32   // 存放整型, 长整型, 浮点数变量
	ref *Object // 存放引用类型变量
}
