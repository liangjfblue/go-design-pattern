/**
 *
 * @author liangjf
 * @create on 2020/6/24
 * @version 1.0
 */
package _5_heap

/**
insert
sift_up —— 用于插入元素
get_max —— 返回最大值但不移除元素
get_size() —— 返回存储的元素数量
is_empty() —— 若堆为空则返回 true
extract_max —— 返回最大值并移除
sift_down —— 用于获取最大值元素
remove(i) —— 删除指定索引的元素
heapify —— 构建堆，用于堆排序
heap_sort() —— 拿到一个未排序的数组，然后使用大顶堆进行就地排序
*/

type Heaper interface {
	Size() int
	Cap() int
	IsEmpty() bool
	Insert(interface{})
	GetMax() interface{}
	ExtractMax() interface{}
	Replace(interface{}) interface{}
	BuildHeap([]interface{}) Heaper
	String() string
}
