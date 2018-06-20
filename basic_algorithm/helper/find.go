package helper

import (
	"reflect"
	"fmt"
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
