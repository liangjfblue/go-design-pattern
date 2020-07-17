/**
 *
 * @author liangjf
 * @create on 2020/7/17
 * @version 1.0
 */
package skiplist

import (
	"bytes"
	"errors"
	"fmt"
	"math/rand"
	"sync"
)

var (
	ErrSkipListNil    = errors.New("skiplist: skiplist is nil")
	ErrCanNotFindNode = errors.New("skiplist: can not find node")
)

var (
	_defaultMaxLevel = 16
	_maxLevel        = 64
	_skipListP       = 0.5
)

type Node struct {
	Key      int
	Value    interface{}
	forward  []*Node
	maxLevel int
}

//skipList 跳表(非线性安全)
type skipList struct {
	Head        *Node
	curLevel    int
	setMaxLevel int
	size        int
	mu          sync.RWMutex
}

func New() ISKipList {
	return NewWithMaxLevel(_defaultMaxLevel)
}

func NewWithMaxLevel(maxLevel int) ISKipList {
	if maxLevel < 1 || maxLevel > _maxLevel {
		panic("skiplist maxLevel must be 1~64")
	}
	s := new(skipList)

	s.curLevel = 1
	s.size = 0
	s.setMaxLevel = maxLevel
	s.Head = &Node{
		Key:     -1,
		Value:   nil,
		forward: make([]*Node, maxLevel),
	}

	return s
}

func (s *skipList) Insert(node *Node) error {
	if s.Head == nil {
		return ErrSkipListNil
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	//概率层级
	level := randomLevel()

	node.forward = make([]*Node, s.setMaxLevel)
	node.maxLevel = level

	update := make([]*Node, level)
	for i := 0; i < level; i++ {
		update[i] = s.Head
	}

	//找出合适的插入位置
	head := s.Head
	for i := level - 1; i >= 0; i-- {
		for head.forward[i] != nil && head.forward[i].Key < node.Key {
			head = head.forward[i]
		}
		update[i] = head
	}

	//已存在直接更新
	x := head.forward[0]
	if x != nil && x.Key == node.Key {
		x.Value = node.Value
		return nil
	}

	//插入新节点, 更新多有前节点的后继, 插入节点的后继
	for i := 0; i < level; i++ {
		node.forward[i] = update[i].forward[i]
		update[i].forward[i] = node
	}

	//概率层级
	if level > s.curLevel {
		s.curLevel = level
	}

	s.size++

	return nil
}

//Delete 删除node
func (s *skipList) Delete(key int) error {
	if s.Head == nil {
		return ErrSkipListNil
	}

	var (
		head   = s.Head
		update = make([]*Node, s.setMaxLevel)
	)

	s.mu.Lock()
	defer s.mu.Unlock()

	//找出合适的删除位置
	for i := s.curLevel - 1; i >= 0; i-- {
		for head.forward[i] != nil && head.forward[i].Key < key {
			head = head.forward[i]
		}
		update[i] = head
	}

	x := head.forward[0]
	if x.Key != key {
		return ErrCanNotFindNode
	}

	for i := s.curLevel - 1; i >= 0; i-- {
		if update[i].forward[i] != nil && update[i].forward[i].Key == key {
			update[i].forward[i] = update[i].forward[i].forward[i]
		}
	}

	x = nil
	for s.curLevel > 1 && s.Head.forward[s.curLevel] == nil {
		s.curLevel--
	}

	s.size--

	return nil
}

//Search 查找node
func (s *skipList) Search(key int) (interface{}, error) {
	if s.Head == nil {
		return nil, ErrSkipListNil
	}

	s.mu.RLock()
	defer s.mu.RUnlock()

	//逐层对比, 找出key大于前索引, 小于后索引, 即key肯定在二者之间
	//注意:	这里的结果是head.forward[0].Key, 是因为概率是50%,25%...递增
	//		(因为逐层1/2, 上层与下层跨度是2, 中间肯定只有一个node)
	head := s.Head
	for i := s.curLevel - 1; i >= 0; i-- {
		for head.forward[i] != nil && head.forward[i].Key < key {
			head = head.forward[i]
		}
	}

	if head.forward[0].Key != key {
		return nil, ErrCanNotFindNode
	}

	return head.forward[0].Value, nil
}

//GetLevel 获取跳表的层数
func (s *skipList) GetLevel() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.curLevel
}

//GetSize 获取跳表元素个数
func (s *skipList) GetSize() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.size
}

func (s *skipList) String() string {
	var (
		buf bytes.Buffer
	)

	buf.WriteString(fmt.Sprintf("skiplist: size = %d\n", s.size))
	buf.WriteString("head=>")
	head := s.Head
	for head.forward[0] != nil {
		buf.WriteString(fmt.Sprint(head.forward[0].Value) + "=>")
		head = head.forward[0]
	}
	buf.Truncate(buf.Len() - 2)
	return buf.String()
}

// 理论来讲，一级索引中元素个数应该占原始数据的 50%，二级索引中元素个数占 25%，三级索引12.5% ，一直到最顶层。
// 因为这里每一层的晋升概率是 50%。对于每一个新插入的节点，都需要调用 randomLevel 生成一个合理的层数。
// 该 randomLevel 方法会随机生成 1~MAX_LEVEL 之间的数，且 ：
//        50%的概率返回 1
//        25%的概率返回 2
//      12.5%的概率返回 3
//     	...
func randomLevel() int {
	level := 1
	for rand.Float64() < _skipListP && level < _defaultMaxLevel {
		level++
	}

	return level
}
