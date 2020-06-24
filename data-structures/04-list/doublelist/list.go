/**
 *
 * @author liangjf
 * @create on 2020/6/24
 * @version 1.0
 */
package doublelist

import (
	"bytes"
	"fmt"
	list2 "go-design-pattern/data-structures/04-list"
)

type list struct {
	head  *list2.Node
	tail  *list2.Node
	value interface{}
	size  int
}

func New() list2.Lister {
	l := new(list)
	node := new(list2.Node)
	l.head.Next = node
	l.tail.Next = node
	return l
}

func (l *list) Size() int {
	return l.size
}

func (l *list) IsEmpty() bool {
	return l.size == 0
}

func (l *list) Next() *list2.Node {
	return l.head
}

func (l *list) AddFirst(e interface{}) {
	defer func() {
		l.size++
	}()
}

func (l *list) AddLast(e interface{}) {
	panic("implement me")
}

func (l *list) Remove(node *list2.Node) interface{} {
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
	//TODO
	return 0
}

func (l *list) String() string {
	var (
		buf bytes.Buffer
	)

	buf.WriteString(fmt.Sprintf("04-list: size = %d\t\t", l.Size()))
	buf.WriteString("[")

	head := l.head.Next
	for head != nil {
		buf.WriteString(fmt.Sprint(head.Value))
		head = head.Next
		buf.WriteString(",")
	}
	buf.Truncate(buf.Len() - 1)
	buf.WriteString("]")

	return buf.String()
}
