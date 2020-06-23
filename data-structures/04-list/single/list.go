/**
 *
 * @author liangjf
 * @create on 2020/6/23
 * @version 1.0
 */
package single

import (
	"bytes"
	"errors"
	"fmt"
	list2 "go-design-pattern/data-structures/04-list"
)

var (
	ErrListEmpty = errors.New("list empty")
)

type list struct {
	head *list2.Node
	size int
}

func New() list2.Lister {
	head := new(list2.Node)
	head.Value = 0
	head.Next = nil

	return &list{
		head: head,
		size: 0,
	}
}

func (l *list) Size() int {
	return l.size
}

func (l *list) IsEmpty() bool {
	return l.head.Next == nil
}

func (l *list) Next() *list2.Node {
	return l.head
}

func (l *list) AddFirst(e interface{}) {
	defer func() {
		l.size++
	}()

	if l.head.Next == nil {
		l.head.Next = &list2.Node{
			Value: e,
			Next:  nil,
		}
		return
	}

	tmpHead := l.head
	node := &list2.Node{
		Value: e,
		Next:  tmpHead.Next,
	}

	l.head.Next = node
}

func (l *list) AddLast(e interface{}) {
	defer func() {
		l.size++
	}()

	if l.head.Next == nil {
		l.head.Next = &list2.Node{
			Value: e,
			Next:  nil,
		}
		return
	}

	pre := l.head
	for pre.Next != nil {
		pre = pre.Next
	}

	pre.Next = &list2.Node{
		Value: e,
		Next:  nil,
	}
}

func (l *list) Remove(n *list2.Node) interface{} {
	if l.head.Next == nil {
		return ErrListEmpty
	}

	defer func() {
		l.size--
	}()

	if l.head.Next.Next == nil {
		tmp := l.head.Next
		l.head.Next = nil
		tmp.Next = nil
		return tmp.Value
	}

	pre := l.head
	for pre.Next != nil {
		if n == pre.Next {
			break
		}
		pre = pre.Next
	}

	tmp := pre.Next
	pre.Next = pre.Next.Next
	tmp.Next = nil

	return tmp.Value
}

func (l *list) RemoveFirst() interface{} {
	return l.Remove(l.head.Next)
}

func (l *list) RemoveLast() interface{} {
	cur := l.head
	for cur.Next != nil {
		cur = cur.Next
	}

	return l.Remove(cur)
}

func (l *list) Contains(e interface{}) bool {
	pre := l.head
	for pre.Next != nil {
		if e == pre.Value {
			return true
		}
		pre = pre.Next
	}
	return false
}

func (l *list) Find(e interface{}) int {
	index := 0
	pre := l.head
	for pre.Next != nil {
		index++
		if e == pre.Value {
			return index
		}
		pre = pre.Next
	}
	return -1
}

func (l *list) FindAll(e interface{}) []int {
	idx := 0
	idxs := make([]int, 0)
	pre := l.head

	for pre.Next != nil {
		idx++
		if e == pre.Value {
			idxs = append(idxs, idx)
		}
		pre = pre.Next
	}
	return idxs
}

func (l *list) RemoveElement(e interface{}) bool {
	pre := l.head
	for pre != nil {
		if e == pre.Value {
			l.Remove(pre)
			return true
		}
		pre = pre.Next
	}
	return false
}

func (l *list) RemoveAllElement(e interface{}) bool {
	hadRemove := false
	pre := l.head
	for pre != nil {
		tmp := pre.Next
		if e == pre.Value {
			hadRemove = true
			l.Remove(pre)
		}
		pre = tmp
	}
	return hadRemove
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
