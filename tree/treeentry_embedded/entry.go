package main

import (
	"fmt"
	"learngo/tree"
)

type myTreeNode struct {
	*tree.Node // Embedding 内嵌
}

func (myNode *myTreeNode) postOrder() {
	if myNode == nil || myNode.Node == nil {
		return
	}
	left := myTreeNode{myNode.Left}
	right := myTreeNode{myNode.Right}

	left.postOrder()
	right.postOrder()
	myNode.Print()
}

func (myNode *myTreeNode) Traverse() {
	fmt.Println("This method is shadowed.")
}

func main() {
	root := myTreeNode{&tree.Node{Value: 3}}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{Value: 5}
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateNode(2)
	root.Right.Left.SetValue(4)

	fmt.Println("In-order traversal: ")
	fmt.Print("root.Traverse(): ")
	root.Traverse()
	fmt.Print("root.Node.Traverse(): ")
	root.Node.Traverse()
	fmt.Println()

	fmt.Print("My own post-order traversal: ")
	root.postOrder()
	fmt.Println()
}
