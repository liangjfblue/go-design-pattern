/**
 *
 * @author liangjf
 * @create on 2020/7/17
 * @version 1.0
 */
package skiplist

type ISKipList interface {
	Insert(*Node) error
	Delete(int) error
	Search(int) (interface{}, error)
	GetLevel() int
	GetSize() int
}
