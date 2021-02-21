package main

import (
	"fmt"
	"learngo/gocode/basic/tree"
)

//类似面向对象的继承
type myTreeNode struct {
	node *tree.Node
}

func (myNode *myTreeNode) postOrder() {
	if myNode == nil || myNode.node == nil {
		return
	}

	left := myTreeNode{myNode.node.Left}
	//myTreeNode表示left是myTreeNode的对象，可以调用该对象的方法
	right := myTreeNode{myNode.node.Right}
	left.postOrder()
	right.postOrder()
	myNode.node.Print()
	/*
		由于postOrder为myTreeNode的方法，所以left和right需要用
		myTreeNode{}来定义
		而Print为Node的方法，只需要用.node就可找到Node
	*/
}

//重载Traverse
func (node *myTreeNode) Traverse() {
	fmt.Println("This method is shadowed.")
}

func main() {
	root := myTreeNode{&tree.Node{Value: 3}}
	root.node.Left = &tree.Node{}
	root.node.Right = &tree.Node{5, nil, nil}
	root.node.Right.Left = new(tree.Node)
	root.node.Left.Right = tree.CreateNode(2)
	root.node.Right.Left.SetValue(4)

	fmt.Println("In-order traversal: ")
	fmt.Print("root.Traverse(): ")
	root.Traverse() //调用重载函数
	fmt.Print("root.node.Traverse(): ")
	root.node.Traverse() //调用原始函数
	fmt.Println()

	fmt.Print("My own post-order traversal: ")
	root.postOrder()
	fmt.Println()
}
