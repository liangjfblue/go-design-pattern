/**
 *
 * @author liangjf
 * @create on 2020/6/23
 * @version 1.0
 */
package array_stack

import (
	"bytes"
	"fmt"
	_1_array "go-design-pattern/data-structures/01-array"
	stack2 "go-design-pattern/data-structures/03-stack"
)

type stack struct {
	data _1_array.Arrayer
}

func New(cap int) stack2.Stacker {
	return &stack{
		data: _1_array.New(cap),
	}
}

//Cap 栈大小
func (s *stack) Size() int {
	return s.data.Size()
}

//IsEmpty 队列是否为空
func (s *stack) IsEmpty() bool {
	return s.data.IsEmpty()
}

//Push 入栈
func (s *stack) Push(e interface{}) {
	s.data.AddLast(e)
}

//Pop 出栈
func (s *stack) Pop() interface{} {
	return s.data.RemoveLast()
}

//Peek 获取首元素
func (s *stack) Peek() interface{} {
	return s.data.Get(s.Size() - 1)
}

//String 打印
func (s *stack) String() string {
	var (
		buf bytes.Buffer
	)

	buf.WriteString(fmt.Sprintf("03-stack: size = %d, capacity = %d\n", s.data.Size(), s.data.Cap()))
	buf.WriteString("[")
	for i := 0; i < s.data.Size(); i++ {
		buf.WriteString(fmt.Sprint(s.data.Get(i)))
		if i != (s.data.Size() - 1) {
			buf.WriteString(", ")
		}
	}
	buf.WriteString("]")

	return buf.String()
}
