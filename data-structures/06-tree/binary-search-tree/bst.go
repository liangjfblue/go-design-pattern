/**
 *
 * @author liangjf
 * @create on 2020/6/28
 * @version 1.0
 */
package binary_search_tree

import (
	"fmt"
	"go-design-pattern/data-structures/02-queue/ringqueue"
	_6_tree "go-design-pattern/data-structures/06-tree"
	"go-design-pattern/data-structures/utils"
)

type BST struct {
	root *_6_tree.Node
	size int
}

func New() _6_tree.Treer {
	return &BST{}
}

func (b *BST) Size() int {
	return b.size
}

func (b *BST) IsEmpty() bool {
	return b.size == 0
}

func (b *BST) Add(e interface{}) {
	b.root = b.add(b.root, e)
}

func (b *BST) add(n *_6_tree.Node, e interface{}) *_6_tree.Node {
	if n == nil {
		b.size++
		return &_6_tree.Node{
			E: e,
		}
	}

	if utils.Compare(e, n.E) > 0 {
		n.Right = b.add(n.Right, e)
	} else if utils.Compare(e, n.E) < 0 {
		n.Left = b.add(n.Left, e)
	}

	return n
}

func (b *BST) Contains(e interface{}) bool {
	return b.contains(b.root, e)
}

func (b *BST) contains(n *_6_tree.Node, e interface{}) bool {
	if n == nil {
		return false
	}

	if utils.Compare(e, n.E) > 0 {
		return b.contains(n.Right, e)
	} else if utils.Compare(e, n.E) < 0 {
		return b.contains(n.Left, e)
	} else {
		return true
	}
}

func (b *BST) PreOrder() {
	b.preOrder(b.root)
}

func (b *BST) preOrder(n *_6_tree.Node) {
	if n == nil {
		return
	}

	//TODO 业务逻辑
	fmt.Println(n.E)
	b.preOrder(n.Left)
	b.preOrder(n.Right)
}

func (b *BST) InOrder() {
	b.inOrder(b.root)
}

func (b *BST) inOrder(n *_6_tree.Node) {
	if n == nil {
		return
	}

	b.inOrder(n.Left)
	//TODO 业务逻辑
	fmt.Println(n.E)
	b.inOrder(n.Right)
}

func (b *BST) PostOrder() {
	b.postOrder(b.root)
}

func (b *BST) postOrder(n *_6_tree.Node) {
	if n == nil {
		return
	}

	b.postOrder(n.Left)
	b.postOrder(n.Right)
	//TODO 业务逻辑
	fmt.Println(n.E)
}

//LevelOrder 层序遍历
func (b *BST) LevelOrder() {
	//借助队列的先进先出特性
	queue := ringqueue.New(16)
	queue.EnQueue(b.root)
	for !queue.IsEmpty() {
		n := queue.DeQueue().(*_6_tree.Node)
		fmt.Println(n.E)

		if n.Left != nil {
			queue.EnQueue(n.Left)
		}

		if n.Right != nil {
			queue.EnQueue(n.Right)
		}
	}
}

//Minimum 返回最小节点的值
func (b *BST) Minimum() interface{} {
	if b.IsEmpty() {
		panic("bst empty")
	}
	//return b.minimumV2(b.root).E
	return b.minimum(b.root).E
}

func (b *BST) minimum(n *_6_tree.Node) *_6_tree.Node {
	if n.Left == nil {
		return n
	}

	return b.minimum(n.Left)
}

//minimumV2 minimum的非递归形式
func (b *BST) minimumV2(n *_6_tree.Node) *_6_tree.Node {
	root := n
	for root.Left != nil {
		root = root.Left
	}
	return root
}

//Maximum 返回最大节点的值
func (b *BST) Maximum() interface{} {
	if b.IsEmpty() {
		panic("bst empty")
	}
	//return b.maximumV2(b.root).E
	return b.maximum(b.root).E
}

func (b *BST) maximum(n *_6_tree.Node) *_6_tree.Node {
	if n.Right == nil {
		return n
	}

	return b.maximum(n.Right)
}

//maximumV2 maximum的非递归形式
func (b *BST) maximumV2(n *_6_tree.Node) *_6_tree.Node {
	root := n
	for root.Right != nil {
		root = root.Right
	}
	return root
}

//RemoveMin 删除最小节点
func (b *BST) RemoveMin() interface{} {
	minN := b.Minimum()
	b.root = b.removeMin(b.root)
	return minN
}

func (b *BST) removeMin(n *_6_tree.Node) *_6_tree.Node {
	if n.Left == nil {
		b.size--
		//右节点放到左节点, 保持二叉树结构(完全二叉树)
		rNode := n.Right
		return rNode
	}

	n.Left = b.removeMin(n.Left)
	return n
}

//RemoveMax 删除最大节点
func (b *BST) RemoveMax() interface{} {
	maxN := b.Maximum()
	b.root = b.removeMax(b.root)
	return maxN
}

func (b *BST) removeMax(n *_6_tree.Node) *_6_tree.Node {
	if n.Right == nil {
		b.size--
		//保持二叉树结构(完全二叉树)
		lNode := n.Left
		return lNode
	}

	n.Right = b.removeMax(n.Right)
	return n
}

//Remove 删除指定值的节点
func (b *BST) Remove(e interface{}) {
	b.root = b.remove(b.root, e)
}

func (b *BST) remove(n *_6_tree.Node, e interface{}) *_6_tree.Node {
	if n == nil {
		return nil
	}

	if utils.Compare(e, n.E) > 0 {
		n.Right = b.remove(n.Right, e)
		return n
	} else if utils.Compare(e, n.E) < 0 {
		n.Left = b.remove(n.Left, e)
		return n
	} else {
		if n.Left == nil {
			b.size--
			rNode := n.Right
			return rNode
		}

		if n.Right == nil {
			b.size--
			lNode := n.Left
			return lNode
		}

		//删除的是非叶子节点, 获取右子树的最小节点填充到被删除节点, 并删除此右子树的最小节点
		minNode := b.minimum(n.Right)
		minNode.Left = n.Left
		minNode.Right = b.removeMin(n.Right)
		n.Left = nil
		n.Right = nil
		return minNode
	}
}

func (b *BST) String() string {
	return ""
}
