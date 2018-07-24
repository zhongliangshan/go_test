package helper

import (
	"reflect"
	"fmt"
	"container/list"
)

type Node struct {
	Key int
	Value int
	Left *Node
	Right *Node
}

type Nodes struct {
	Root *Node
	Count int
}

var Root Nodes

func (n *Nodes) Init() {
	Root.Count = 0
	Root.Root = &Node{}
}


func (n *Nodes) Insert(key , value int) {
	n.Root = n.insert(Root.Root , key , value)
}

func (n *Nodes) CreateNewNode(key , value int) *Node {
	var node Node
	node.Key = key
	node.Value = value
	node.Right = &Node{}
	node.Left = &Node{}

	return &node
}
func isEmpty(a interface{}) bool {
	v := reflect.ValueOf(a)
	if v.Kind() == reflect.Ptr {
		v=v.Elem()
	}
	return v.Interface() == reflect.Zero(v.Type()).Interface()
}


func (n *Nodes) insert(node *Node , key , value int) *Node {
	if node == nil || isEmpty(node) {
		n.Count++
		return n.CreateNewNode(key , value)
	}

	if node.Key == key {
		node.Value = value
		return node
	}

	if node.Left.Key > key {
		node.Left = n.insert(node.Left , key , value)
	}

	if node.Right.Key < key {
		node.Right = n.insert(node.Right , key , value)
	}

	return node
}

// 中序遍历 左 中 右
func (n *Nodes) PrintNodes(node *Node) {
	if isEmpty(node) {
		return
	}
	if node.Left != nil {
		n.PrintNodes(node.Left)
	}

	if node.Right != nil {
		n.PrintNodes(node.Right)
	}

	fmt.Println("key:" , node.Key, ";value:",node.Value)
}

func (n *Nodes) Contain(node *Node , key int) error {
	if node == nil || isEmpty(node) {
		return fmt.Errorf("not find")
	}

	if node.Key == key {
		return nil
	}

	if node.Left.Key > key {
		return n.Contain(node.Left , key )
	} else  {
		 return n.Contain(node.Right , key )
	}

}

func (n *Nodes) Search(node *Node , key int) (int,error) {
	if node == nil || isEmpty(node) {
		return 0 , fmt.Errorf("not find")
	}
	if node.Key == key {
		return node.Value , nil
	}

	if node.Left.Key > key {
		return n.Search(node.Left , key )
	} else {
		return n.Search(node.Right , key )
	}
}

// 前序遍历
func (n *Nodes) PreOrder(node *Node) {
	if node == nil || isEmpty(node) {
		return
	}

	fmt.Println("key:" , node.Key , ";value:" , node.Value)

	n.PreOrder(node.Left)
	n.PreOrder(node.Right)
}
// 中序遍历

func (n *Nodes) InOrder(node *Node) {
	if node == nil || isEmpty(node) {
		return
	}
	n.InOrder(node.Left)
	fmt.Println("key:" , node.Key , ";value:" , node.Value)
	n.InOrder(node.Right)
}
// 后序遍历
func (n *Nodes) LastOrder(node *Node) {
	if node == nil || isEmpty(node) {
		return
	}
	n.LastOrder(node.Left)
	n.LastOrder(node.Right)
	fmt.Println("key:" , node.Key , ";value:" , node.Value)

}

// 队列中的数据 放进去是什么 取出来后 需要转化成什么类型的数据 不然 就得不到正确的数据   单类型 不需要这样处理 复杂的需要
func (n *Nodes) LevelOrder() {
	l := list.New()
	l.PushBack(Root.Root)

	for l.Len() != 0 {
		n := l.Front()
		node , _ := (n.Value).(*Node)
		l.Remove(n)

		fmt.Println("key:" , node.Key , ";value:" , node.Value)
		if node.Left != nil && !isEmpty(node.Left) {
			l.PushBack(node.Left)
		}

		if node.Right != nil && !isEmpty(node.Right) {
			l.PushBack(node.Right)
		}
	}
}

func (n *Nodes) FindMinNode(node *Node) *Node{
	if node.Left == nil || isEmpty(node.Left) {
		return node
	}

	return n.FindMinNode(node.Left)
}


func (n *Nodes) FindMaxNode(node *Node) *Node{
	if node.Right == nil || isEmpty(node.Right) {
		return node
	}

	return n.FindMaxNode(node.Right)

}

func (n *Nodes) DelMinNode(node *Node) *Node {
	if node.Left == nil || isEmpty(node.Left) {
		nodeRight := node.Right
		n.Count --

		return nodeRight

	}

	node.Left =  n.DelMinNode(node.Left)

	return node
}

func (n *Nodes) DelMaxNode(node *Node) *Node {
	if node.Right == nil || isEmpty(node.Right) {
		nodeLeft := node.Left
		n.Count --

		return nodeLeft

	}

	node.Right =  n.DelMaxNode(node.Right)
	return node
}

func (n *Nodes) DelNode(node *Node , key int) *Node {
	if isEmpty(node) {
		return nil
	}

	if key < node.Key {
		node.Left = n.DelNode(node.Left , key)
		return node
	} else if key > node.Key {
		node.Right = n.DelNode(node.Right , key)
		return node
	} else {
		// 等于的情况 实际删除
		if node.Left == nil || isEmpty(node.Left) {
			nodeRight := node.Right
			n.Count --

			return nodeRight
		}

		if node.Right == nil || isEmpty(node.Right) {
			nodeLeft := node.Left
			n.Count --

			return nodeLeft
		}

		// 左右都不为空 找出右子树的最小的节点 然后删除
		minNode := n.FindMinNode(node.Right)

		// 删除掉 右子树中的最小节点 然后返回新的根节点给 已经赋值的最小节点的右边
		minNode.Right = n.DelMinNode(node.Right)

		minNode.Left = node.Left

		return minNode
	}
}