package main

import "fmt"

type Node struct {
	Val    int
	Left   *Node
	Right  *Node
	Height int
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
	n.Height = 1 + max(height(n.Left), height(n.Right))
	n = balance(n)
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

func height(n *Node) int {

	if n == nil {
		return -1
	}
	return n.Height
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func balance(n *Node) *Node {

	if n == nil {
		return n
	}
	if height(n.Left)-height(n.Right) > 1 {
		if height(n.Left.Left) >= height(n.Left.Right) {
			n = rotateRight(n)
		} else {
			n.Left = rotateLeft(n.Left)
			n = rotateRight(n)
		}
	} else if height(n.Right)-height(n.Left) > 1 {
		if height(n.Right.Right) >= height(n.Right.Left) {
			n = rotateLeft(n)
		} else {
			n.Right = rotateRight(n.Right)
			n = rotateLeft(n)
		}
	}
	return n
}

func rotateLeft(n *Node) *Node {

	newRoot := n.Right
	n.Right = newRoot.Left
	newRoot.Left = n
	n.Height = 1 + max(height(n.Left), height(n.Right))
	newRoot.Height = 1 + max(height(newRoot.Left), height(newRoot.Right))
	return newRoot
}

func rotateRight(n *Node) *Node {

	newRoot := n.Left
	n.Left = newRoot.Right
	newRoot.Right = n
	n.Height = 1 + max(height(n.Left), height(n.Right))
	newRoot.Height = 1 + max(height(newRoot.Left), height(newRoot.Right))
	return newRoot
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

}
