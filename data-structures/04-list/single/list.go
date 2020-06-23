/**
 *
 * @author liangjf
 * @create on 2020/6/23
 * @version 1.0
 */
package single

import (
	"bytes"
	"fmt"
)

type Node struct {
	value interface{}
	next  *Node
}

type list struct {
	head *Node
	size int
}

func (l *list) Size() int {
	return l.size
}

func (l *list) IsEmpty() bool {
	return l.head.next == nil
}

func (l *list) AddFirst(e interface{}) {
	panic("implement me")
}

func (l *list) AddLast(e interface{}) {
	panic("implement me")
}

func (l *list) Remove(e int) interface{} {
	panic("implement me")
}

func (l *list) RemoveFirst() interface{} {
	panic("implement me")
}

func (l *list) RemoveLast() interface{} {
	panic("implement me")
}

func (l *list) Contains(e interface{}) bool {
	panic("implement me")
}

func (l *list) Find(e interface{}) int {
	panic("implement me")
}

func (l *list) FindAll(e interface{}) []int {
	panic("implement me")
}

func (l *list) RemoveElement(e interface{}) bool {
	panic("implement me")
}

func (l *list) RemoveAllElement(e interface{}) bool {
	panic("implement me")
}

func (l *list) HasRing() int {
	panic("implement me")
}

func (l *list) String() string {
	var (
		buf bytes.Buffer
	)

	buf.WriteString(fmt.Sprintf("04-list: size = %d\n", l.Size()))
	buf.WriteString("[")

	head := l.head.next
	for head.next != nil {
		buf.WriteString(fmt.Sprint(l.head.value))
		head = head.next
	}
	buf.WriteString("]")

	return buf.String()
}
