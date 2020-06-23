/**
 *
 * @author liangjf
 * @create on 2020/6/23
 * @version 1.0
 */
package _4_list

type Lister interface {
	Size() int
	IsEmpty() bool
	AddFirst(interface{})
	AddLast(interface{})
	Remove(int) interface{}
	RemoveFirst() interface{}
	RemoveLast() interface{}
	Contains(interface{}) bool
	Find(interface{}) int
	FindAll(interface{}) []int
	RemoveElement(interface{}) bool
	RemoveAllElement(interface{}) bool
	HasRing() int
	String() string
}
