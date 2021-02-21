package tree

import (
	"fmt"
)

type Node struct {
	Value       int
	Left, Right *Node
}

//定义value访问函数
func (node Node) Print() { //参数为node Node
	fmt.Print(node.Value, " ")
}

//定义改变value的函数
func (node *Node) ChangeValue(val int) { //前面的括号为函数调用者
	node.Value = val
}

func (node *Node) SetValue(value int) {
	if node == nil {
		fmt.Println("Setting Value to nil" +
			"node. Ignored")
		return
	}
	node.Value = value
}

//工厂函数自定义node的创建
func CreateNode(value int) *Node {
	return &Node{Value: value} //注意返回的是局部变量的地址！
}
