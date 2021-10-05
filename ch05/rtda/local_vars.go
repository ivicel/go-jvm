package rtda

import "math"

// 本地变量表数组
type LocalVars []Slot

// newLocalVars 生成个的本地变量表
func newLocalVars(maxLocals uint) LocalVars {
	if maxLocals > 0 {
		return make(LocalVars, maxLocals)
	}

	return nil
}

// SetInt 设置本地整数变量
func (l *LocalVars) SetInt(index uint, val int32) {
	(*l)[index].num = val
}

// GetInt 获取本地整数变量
func (l *LocalVars) GetInt(index uint) int32 {
	return (*l)[index].num
}

// SetFloat 设置浮点数
func (l *LocalVars) SetFloat(index uint, val float32) {
	bits := math.Float32bits(val)
	(*l)[index].num = int32(bits)
}

// GetFloat 获取浮点数变量
func (l *LocalVars) GetFloat(index uint) float32 {
	bits := uint32((*l)[index].num)
	return math.Float32frombits(bits)
}

// SetLong 设置长整型, 需要用两个位来存储
func (l *LocalVars) SetLong(index uint, val int64) {
	(*l)[index].num = int32(val)
	(*l)[index+1].num = int32(val >> 32)
}

func (l *LocalVars) GetLong(index uint) int64 {
	low := uint32((*l)[index].num)
	high := uint32((*l)[index+1].num)

	return int64(low) | int64(high)<<32
}

// SetDouble 设置 double 类型
func (l *LocalVars) SetDouble(index uint, val float64) {
	bits := math.Float64bits(val)
	l.SetLong(index, int64(bits))
}

func (l *LocalVars) GetDouble(index uint) float64 {
	bits := l.GetLong(index)

	return math.Float64frombits(uint64(bits))
}

func (l *LocalVars) SetRef(index uint, ref *Object) {
	(*l)[index].ref = ref
}

func (l *LocalVars) GetRef(index uint) *Object {
	return (*l)[index].ref
}
