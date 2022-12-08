package main

import "fmt"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func (n *Node) Insert(val int) {
	if n == nil {
		return
	} else if val <= n.Val {
		if n.Left == nil {
			n.Left = &Node{Val: val}
		} else {
			n.Left.Insert(val)
		}
	} else {
		if n.Right == nil {
			n.Right = &Node{Val: val}
		} else {
			n.Right.Insert(val)
		}
	}
}

func (n *Node) Contains(val int) bool {
	if n == nil {
		return false
	}
	if n.Val == val {
		return true
	} else if val < n.Val {
		return n.Left.Contains(val)
	} else {
		return n.Right.Contains(val)
	}
}

func (n *Node) PrintInOrder() {
	if n == nil {
		return
	}
	n.Left.PrintInOrder()
	fmt.Println(n.Val)
	n.Right.PrintInOrder()
}

func main() {
	root := &Node{Val: 10}
	root.Insert(5)
	root.Insert(15)
	root.Insert(2)
	root.Insert(7)
	root.Insert(12)
	root.Insert(17)
	root.PrintInOrder()
	fmt.Println(root.Contains(7))
	fmt.Println(root.Contains(13))
}
