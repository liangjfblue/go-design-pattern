/**
 *
 * @author liangjf
 * @create on 2020/6/24
 * @version 1.0
 */
package priorityqueue

import (
	"bytes"
	"fmt"
	queue2 "go-design-pattern/data-structures/02-queue"
	_5_heap "go-design-pattern/data-structures/05-heap"
	big_heap "go-design-pattern/data-structures/05-heap/big-heap"
)

type queue struct {
	maxHeap _5_heap.Heaper
}

func New(cap int) queue2.Queuer {
	return &queue{
		maxHeap: big_heap.New(cap),
	}
}

func (q *queue) Cap() int {
	return q.maxHeap.Cap()
}

func (q *queue) Size() int {
	return q.maxHeap.Size()
}

func (q *queue) IsEmpty() bool {
	return q.maxHeap.IsEmpty()
}

func (q *queue) EnQueue(e interface{}) {
	q.maxHeap.Insert(e)
}

func (q *queue) DeQueue() interface{} {
	return q.maxHeap.ExtractMax()
}

func (q *queue) Front() interface{} {
	return q.maxHeap.GetMax()
}

func (q *queue) String() string {
	var (
		buf bytes.Buffer
	)

	buf.WriteString(fmt.Sprintf("04-list: size = %d, cap = %d\t\t", q.Size(), q.Cap()))
	buf.WriteString("[")

	for !q.IsEmpty() {
		buf.WriteString(fmt.Sprint(q.maxHeap.ExtractMax()))
		buf.WriteString(",")
	}
	buf.Truncate(buf.Len() - 1)
	buf.WriteString("]")

	return buf.String()
}
