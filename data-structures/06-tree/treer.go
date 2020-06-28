/**
 *
 * @author liangjf
 * @create on 2020/6/28
 * @version 1.0
 */
package _6_tree

type Treer interface {
	Size() int
	IsEmpty() bool
	Add(interface{})
	Contains(interface{}) bool
	PreOrder()
	InOrder()
	PostOrder()
	LevelOrder()
	Minimum() interface{}
	Maximum() interface{}
	RemoveMin() interface{}
	RemoveMax() interface{}
	Remove(interface{})
	String() string
}

type Node struct {
	E     interface{}
	Left  *Node
	Right *Node
}
