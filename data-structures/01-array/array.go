/**
 *
 * @author liangjf
 * @create on 2020/6/23
 * @version 1.0
 */
package _1_array

import (
	"bytes"
	"fmt"
	"go-design-pattern/data-structures/utils"
)

var (
	initCap = 16
)

type array struct {
	size int
	data []interface{}
}

func New(cap int) Arrayer {
	if cap < initCap {
		cap = initCap
	}

	return &array{
		size: 0,
		data: make([]interface{}, cap),
	}
}

//Cap 获取数组的容量
func (a *array) Cap() int {
	return len(a.data)
}

//Size 获得数组中的元素个数
func (a *array) Size() int {
	return a.size
}

//IsEmpty 返回数组是否为空
func (a *array) IsEmpty() bool {
	return a.size == 0
}

//Add 在第index个位置插入一个新元素 e
func (a *array) Add(idx int, e interface{}) {
	if idx < 0 || idx > a.size {
		panic("out of range")
	}

	if len(a.data) == a.size {
		a.resize(int(utils.Min2(uint32(a.size) + 1)))
	}

	for i := a.size - 1; i >= idx; i-- {
		a.data[i+1] = a.data[i]
	}

	a.data[idx] = e
	a.size++
}

//AddLast 向所有元素后添加一个新元素
func (a *array) AddLast(e interface{}) {
	a.Add(a.size, e)
}

//AddFirst 向所有元素前添加一个新元素
func (a *array) AddFirst(e interface{}) {
	a.Add(0, e)
}

//Get 获取index索引位置的元素
func (a *array) Get(idx int) interface{} {
	if idx < 0 || idx >= a.size {
		panic("out of range")
	}

	return a.data[idx]
}

//Set 修改index索引位置的元素
func (a *array) Set(idx int, e interface{}) {
	if idx < 0 || idx >= a.size {
		panic("out of range")
	}

	a.data[idx] = e
}

//Contains 查找数组中是否有元素 e
func (a *array) Contains(e interface{}) bool {
	if a.size == 0 {
		return false
	}

	for _, v := range a.data {
		if e == v {
			return true
		}
	}

	return false
}

//Find 查找数组中元素 e 所在的索引，不存在则返回 -1
func (a *array) Find(e interface{}) int {
	if a.size == 0 {
		return -1
	}

	for idx, v := range a.data {
		if e == v {
			return idx
		}
	}

	return -1
}

//FindAll 查找数组中元素 e 所有的索引组成的切片，不存在则返回 []int{}
func (a *array) FindAll(e interface{}) (indexes []int) {
	var (
		idxs = make([]int, 0)
	)
	if a.size == 0 {
		return idxs
	}

	for idx, v := range a.data {
		if e == v {
			idxs = append(idxs, idx)
		}
	}

	return idxs
}

//Remove 从数组中删除index位置的元素，返回删除的元素
func (a *array) Remove(idx int) interface{} {
	if idx < 0 || idx >= a.size {
		panic("out of range")
	}

	e := a.data[idx]
	for i := idx + 1; i < a.size; i++ {
		a.data[i-1] = a.data[i]
	}

	//后面移动那位置空
	a.data[a.size] = nil
	a.size--

	if a.size == len(a.data)>>2 && len(a.data)>>1 != 0 {
		a.resize(len(a.data) / 2)
	}

	return e
}

//RemoveFirst 从数组中删除第一个元素，返回删除的元素
func (a *array) RemoveFirst() interface{} {
	return a.Remove(0)
}

//RemoveLast 从数组中删除最后一个元素，返回删除的元素
func (a *array) RemoveLast() interface{} {
	return a.Remove(a.size - 1)
}

//RemoveElement 从数组中删除一个元素 e
func (a *array) RemoveElement(e interface{}) bool {
	idx := a.Find(e)
	if idx == -1 {
		return false
	}
	a.Remove(idx)
	return true
}

//RemoveAllElement 从数组中删除所有元素 e
func (a *array) RemoveAllElement(e interface{}) bool {
	idxs := a.FindAll(e)
	if len(idxs) > 0 {
		for _, idx := range idxs {
			a.Remove(idx)
		}
		return true
	}
	return false
}

//resize 为数组扩容
func (a *array) resize(cap int) {
	nData := make([]interface{}, cap)
	for i := 0; i < a.size; i++ {
		nData[i] = a.data[i]
	}
	a.data = nData
}

//String 打印数组
func (a *array) String() string {
	var (
		buf bytes.Buffer
	)

	buf.WriteString(fmt.Sprintf("01-array: size = %d, capacity = %d\n", a.size, len(a.data)))
	buf.WriteString("[")
	for i := 0; i < a.size; i++ {
		buf.WriteString(fmt.Sprint(a.data[i]))
		if i != (a.size - 1) {
			buf.WriteString(", ")
		}
	}
	buf.WriteString("]")

	return buf.String()
}
