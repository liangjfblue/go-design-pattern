/**
 *
 * @author liangjf
 * @create on 2020/6/23
 * @version 1.0
 */
package _3_stack

type Stacker interface {
	Size() int
	IsEmpty() bool
	Push(interface{})
	Pop() interface{}
	Peek() interface{}
	String() string
}
