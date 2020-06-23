/**
 *
 * @author liangjf
 * @create on 2020/6/23
 * @version 1.0
 */
package ringqueue

import (
	"bytes"
	"fmt"
	queue2 "go-design-pattern/data-structures/02-queue"
	"go-design-pattern/data-structures/utils"
)

type queue struct {
	head, tail, count int
	data              []interface{}
}

func New(cap uint32) queue2.Queuer {
	return &queue{
		data: make([]interface{}, utils.Min2(cap)),
	}
}

//Cap 队列容量
func (q *queue) Cap() int {
	return cap(q.data)
}

//Size 队列大小
func (q *queue) Size() int {
	return q.count
}

//IsEmpty 队列是否为空
func (q *queue) IsEmpty() bool {
	return q.count == 0
}

//EnQueue 入队
func (q *queue) EnQueue(e interface{}) {
	if q.count == len(q.data) {
		q.resize(true)
	}

	q.data[q.tail] = e
	q.tail = (q.tail + 1) & (len(q.data) - 1)
	q.count++
}

//DeQueue 出队
func (q *queue) DeQueue() interface{} {
	e := q.data[q.head]
	q.data[q.head] = nil
	q.head = (q.head + 1) & (len(q.data) - 1)
	q.count--

	//1/4时触发缩容为1/2(避免缩容马上又扩容)
	if len(q.data) > 0 && q.count<<2 <= len(q.data) {
		q.resize(false)
	}

	return e
}

//Front 获取首个元素
func (q *queue) Front() interface{} {
	if q.count == 0 {
		panic("02-queue empty")
	}

	return q.data[0]
}

//Get 获取指定位置元素
func (q *queue) get(idx int) interface{} {
	if q.data == nil {
		panic("02-queue empty")
	}
	return q.data[(q.head+idx)&(len(q.data)-1)]
}

//resize 调整queue
func (q *queue) resize(scala bool) {
	count := uint32(q.count)
	if scala {
		count = count + 1
	}
	nEm := make([]interface{}, utils.Min2(count))

	if q.tail > q.head {
		copy(nEm, q.data[q.head:q.tail])
	} else {
		n := copy(nEm, q.data[q.head:])
		copy(nEm[n:], q.data[:q.tail])
	}

	q.head = 0
	q.tail = q.count
	q.data = nEm
	return
}

//String 打印队列
func (q *queue) String() string {
	var (
		buf bytes.Buffer
	)

	buf.WriteString(fmt.Sprintf("02-queue: size = %d, capacity = %d\n", q.Size(), q.Cap()))
	buf.WriteString("[")
	for i := 0; i < q.Size(); i++ {
		buf.WriteString(fmt.Sprint(q.get(i)))
		if i != (q.Size() - 1) {
			buf.WriteString(", ")
		}
	}
	buf.WriteString("]")

	return buf.String()
}
