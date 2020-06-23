/**
 *
 * @author liangjf
 * @create on 2020/6/23
 * @version 1.0
 */
package _1_array

type Arrayer interface {
	Cap() int
	Size() int
	IsEmpty() bool
	Add(int, interface{})
	AddLast(interface{})
	AddFirst(interface{})
	Get(int) interface{}
	Set(int, interface{})
	Contains(interface{}) bool
	Find(interface{}) int
	FindAll(interface{}) []int
	Remove(int) interface{}
	RemoveFirst() interface{}
	RemoveLast() interface{}
	RemoveElement(interface{}) bool
	RemoveAllElement(interface{}) bool
	String() string
}
