/**
 *
 * @author liangjf
 * @create on 2020/6/24
 * @version 1.0
 */
package big_heap

import (
	_1_array "go-design-pattern/data-structures/01-array"
	_5_heap "go-design-pattern/data-structures/05-heap"
	"go-design-pattern/data-structures/utils"
)

type heap struct {
	data _1_array.Arrayer
}

func New(cap int) _5_heap.Heaper {
	return &heap{
		data: _1_array.New(cap),
	}
}

//Size 堆大小
func (h *heap) Size() int {
	return h.data.Size()
}

//Size 堆容量
func (h *heap) Cap() int {
	return h.data.Cap()
}

//IsEmpty 堆是否为空
func (h *heap) IsEmpty() bool {
	return h.data.IsEmpty()
}

func (h *heap) Insert(e interface{}) {
	h.data.AddLast(e)
	h.shiftUp(h.Size() - 1)
}

func (h *heap) GetMax() interface{} {
	return h.getMax()
}

func (h *heap) ExtractMax() interface{} {
	max := h.getMax()

	//把最后一个提上来,然后下沉找到合适位置
	idx0V := h.data.Get(0)
	h.data.Set(0, h.data.Get(h.Size()-1))
	h.data.Set(h.Size()-1, idx0V)

	h.data.RemoveLast()
	h.heapDown(0)

	return max
}

func (h *heap) Replace(e interface{}) interface{} {
	ret := h.getMax()

	h.data.Set(0, e)
	h.heapDown(0)

	return ret
}

//String 堆形打印
func (h *heap) String() string {
	return ""
}

//BuildHeap arr的元素必须是可比较的(非切片, 非map...)
func (h *heap) BuildHeap(arr []interface{}) _5_heap.Heaper {
	for i := 0; i < len(arr); i++ {
		h.data.AddLast(arr[i])
		h.shiftUp(h.Size() - 1)
	}

	return h
}

func (h *heap) parent(idx int) int {
	if idx == 0 {
		panic("idx 0 no parent")
	}
	return (idx - 1) / 2
}

func (h *heap) leftChild(idx int) int {
	return idx*2 + 1
}

func (h *heap) rightChild(idx int) int {
	return idx*2 + 2
}

func (h *heap) getMax() interface{} {
	if h.Size() == 0 {
		panic("heap empty")
	}
	return h.data.Get(0)
}

//shiftUp 向上寻找合适位置(找大)
func (h *heap) shiftUp(i int) {
	for i > 0 && utils.Compare(h.data.Get(i), h.data.Get(h.parent(i))) > 0 {
		iV := h.data.Get(i)
		h.data.Set(i, h.data.Get(h.parent(i)))
		h.data.Set(h.parent(i), iV)
		i = h.parent(i)
	}
}

//heapDown 向下寻找合适位置(找小)
func (h *heap) heapDown(i int) {
	for h.leftChild(i) < h.Size() {
		idx := h.leftChild(i)

		//是否有右孩子, 找出左右孩子哪个大
		if idx+1 < h.Size() && utils.Compare(h.data.Get(idx), h.data.Get(idx+1)) < 0 {
			idx++
		}

		//若下沉比左右孩子都小就停止
		if utils.Compare(h.data.Get(i), h.data.Get(idx)) > 0 {
			break
		}

		//交换父亲和孩子
		iV := h.data.Get(i)
		h.data.Set(i, h.data.Get(idx))
		h.data.Set(idx, iV)

		//更新比较index
		i = idx
	}
}
