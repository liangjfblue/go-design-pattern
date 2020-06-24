/**
 *
 * @author liangjf
 * @create on 2020/6/24
 * @version 1.0
 */
package listqueue

import (
	"bytes"
	"errors"
	"fmt"
	queue2 "go-design-pattern/data-structures/02-queue"
	list2 "go-design-pattern/data-structures/04-list"
	"go-design-pattern/data-structures/04-list/singlelist"
)

type queue struct {
	data list2.Lister
}

func New() queue2.Queuer {
	return &queue{
		data: singlelist.New(),
	}
}

func (q *queue) Cap() int {
	//list方式无容量限制
	return 0
}

func (q *queue) Size() int {
	return q.data.Size()
}

func (q *queue) IsEmpty() bool {
	return q.data.IsEmpty()
}

func (q *queue) EnQueue(e interface{}) {
	q.data.AddLast(e)
}

func (q *queue) DeQueue() interface{} {
	return q.data.RemoveFirst()
}

func (q *queue) Front() interface{} {
	if q.data.Next().Next == nil {
		return errors.New("queue empty")
	}

	return q.data.Next().Next.Value
}

func (q *queue) String() string {
	var (
		buf bytes.Buffer
	)

	buf.WriteString(fmt.Sprintf("02-queue: size = %d\n", q.data.Size()))
	buf.WriteString("head=>")
	head := q.data.Next().Next
	for head != nil {
		buf.WriteString(fmt.Sprint(head.Value))
		buf.WriteString("=>")
		head = head.Next
	}
	buf.WriteString("NULL")

	return buf.String()
}
