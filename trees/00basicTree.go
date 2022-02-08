package main

import "fmt"

var count int

func main() {
	f := fmt.Println
	tree := &BinaryNode{data: 100}
	tree.insert(56)
	tree.insert(54)
	tree.insert(123)
	tree.insert(125)
	tree.insert(58)
	tree.insert(118)
	tree.insert(500)
	tree.insert(121)

	printPreOrder(tree)
	f(tree.search(500))

	f(count)
}

type BinaryNode struct {
	data  int
	left  *BinaryNode
	right *BinaryNode
}

// type BinaryTree struct {
// 	root *BinaryNode
// }

func (n *BinaryNode) insert(key int) {
	if n.data < key {
		//move right
		if n.right == nil {
			n.right = &BinaryNode{data: key}
		} else {
			n.right.insert(key)
		}
	} else {
		//move left
		if n.left == nil {
			n.left = &BinaryNode{data: key}
		} else {
			n.left.insert(key)
		}
	}
}

func (Node *BinaryNode) search(key int) bool {
	count++
	if Node == nil {
		return false
	}
	if Node.data == key {
		return true
	}

	if Node.data < key {
		return Node.right.search(key)
	}

	if Node.data > key {
		return Node.left.search(key)
	}
	return false
}
func printPreOrder(Node *BinaryNode) {
	f := fmt.Println
	f(Node.data)
	if Node.left != nil {
		printPreOrder(Node.left)
	}

	if Node.right != nil {
		printPreOrder(Node.right)
	}

}
