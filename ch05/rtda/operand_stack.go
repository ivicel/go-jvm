package rtda

import "math"

// 操作数栈
type OperandStack struct {
	size  uint   // 栈大小
	slots []Slot // 槽列表
}

func newOperandStack(maxStack uint) *OperandStack {
	if maxStack > 0 {
		return &OperandStack{
			slots: make([]Slot, maxStack),
		}
	}

	return nil
}

// PushInt 往栈顶压入一个 int
func (o *OperandStack) PushInt(val int32) {
	o.slots[o.size].num = val
	o.size++
}

// PopInt 出栈
func (o *OperandStack) PopInt() int32 {
	o.size--
	return o.slots[o.size].num
}

// PushFloat 入栈单精度浮点数
func (o *OperandStack) PushFloat(val float32) {
	bits := math.Float32bits(val)
	o.slots[o.size].num = int32(bits)
	o.size++
}

func (o *OperandStack) PopFloat() float32 {
	o.size--
	bits := uint32(o.slots[o.size].num)
	return math.Float32frombits(bits)
}

// PushLong 入栈 long
func (o *OperandStack) PushLong(val int64) {
	o.slots[o.size].num = int32(val)
	o.slots[o.size+1].num = int32(val >> 32)
	o.size += 2
}

func (o *OperandStack) PopLong() int64 {
	o.size -= 2
	low := uint32(o.slots[o.size].num)
	high := uint32(o.slots[o.size+1].num)
	return int64(low) | int64(high)<<32
}

// PushDouble 入栈 double
func (o *OperandStack) PushDobule(val float64) {
	bits := math.Float64bits(val)
	o.PushLong(int64(bits))
}

func (o *OperandStack) PopDobule() float64 {
	bits := o.PopLong()
	return math.Float64frombits(uint64(bits))
}

// PushRef 入栈引用类型
func (o *OperandStack) PushRef(ref *Object) {
	o.slots[o.size].ref = ref
	o.size++
}

func (o *OperandStack) PopRef() *Object {
	o.size--
	ref := o.slots[o.size].ref
	o.slots[o.size].ref = nil
	return ref
}

// PushSlot 压入一个变量槽
func (o *OperandStack) PushSlot(slot Slot) {
	o.slots[o.size] = slot
	o.size++
}

func (o *OperandStack) PopSlot() Slot {
	o.size--
	return o.slots[o.size]
}
