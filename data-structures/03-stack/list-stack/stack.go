/**
 *
 * @author liangjf
 * @create on 2020/6/24
 * @version 1.0
 */
package list_stack

import (
	"bytes"
	"errors"
	"fmt"
	stack2 "go-design-pattern/data-structures/03-stack"
	list2 "go-design-pattern/data-structures/04-list"
	"go-design-pattern/data-structures/04-list/singlelist"
)

type stack struct {
	data list2.Lister
}

func New() stack2.Stacker {
	return &stack{
		data: singlelist.New(),
	}
}

func (s stack) Size() int {
	return s.data.Size()
}

func (s stack) IsEmpty() bool {
	return s.data.IsEmpty()
}

func (s stack) Push(e interface{}) {
	s.data.AddFirst(e)
}

func (s stack) Pop() interface{} {
	return s.data.RemoveFirst()
}

func (s stack) Peek() interface{} {
	if s.data.Next().Next == nil {
		return errors.New("stack empty")
	}

	return s.data.Next().Next.Value
}

func (s stack) String() string {
	var (
		buf bytes.Buffer
	)

	buf.WriteString(fmt.Sprintf("02-queue: size = %d\n", s.data.Size()))
	buf.WriteString("head=>")
	head := s.data.Next().Next
	for head != nil {
		buf.WriteString(fmt.Sprint(head.Value))
		buf.WriteString("=>")
		head = head.Next
	}
	buf.WriteString("NULL")

	return buf.String()
}
