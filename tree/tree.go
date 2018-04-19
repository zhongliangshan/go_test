package main

import "fmt"

type Node struct {
	Value       int
	Left, Right *Node
}

func main() {
	var tree = Node{2, nil, nil}
	tree.Left = &Node{}
	tree.Left.Value = 3
	tree.Right = &Node{}
	tree.Right.Value = 4

	fmt.Println(tree)
}
