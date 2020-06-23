/**
 *
 * @author liangjf
 * @create on 2020/6/23
 * @version 1.0
 */
package singlequeue

import (
	"bytes"
	"fmt"
	_1_array "go-design-pattern/data-structures/01-array"
	queue2 "go-design-pattern/data-structures/02-queue"
)

type queue struct {
	data _1_array.Arrayer
}

func New(cap int) queue2.Queuer {
	return &queue{
		data: _1_array.New(cap),
	}
}

//Cap 队列容量
func (q *queue) Cap() int {
	return q.data.Cap()
}

//Size 队列大小
func (q *queue) Size() int {
	return q.data.Size()
}

//IsEmpty 队列是否为空
func (q *queue) IsEmpty() bool {
	return q.data.IsEmpty()
}

//EnQueue 入队
func (q *queue) EnQueue(e interface{}) {
	q.data.AddLast(e)
}

//DeQueue 出队
func (q *queue) DeQueue() interface{} {
	return q.data.RemoveFirst()
}

//Front 获取首个元素
func (q *queue) Front() interface{} {
	return q.data.Get(0)
}

//String 打印队列
func (q *queue) String() string {
	var (
		buf bytes.Buffer
	)

	buf.WriteString(fmt.Sprintf("02-queue: size = %d, capacity = %d\n", q.data.Size(), q.data.Cap()))
	buf.WriteString("[")
	for i := 0; i < q.data.Size(); i++ {
		buf.WriteString(fmt.Sprint(q.data.Get(i)))
		if i != (q.data.Size() - 1) {
			buf.WriteString(", ")
		}
	}
	buf.WriteString("]")

	return buf.String()
}
