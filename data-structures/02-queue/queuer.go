/**
 *
 * @author liangjf
 * @create on 2020/6/23
 * @version 1.0
 */
package _2_queue

type Queuer interface {
	Cap() int
	Size() int
	IsEmpty() bool
	EnQueue(interface{})
	DeQueue() interface{}
	Front() interface{}
	String() string
}
