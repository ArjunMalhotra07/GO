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
	tree.insert(117)
	f("Pre Order Traversal")
	printPreOrder(tree)
	f()
	f("Post Order Traversal")
	printPostOrder(tree)
	f()
	f("In Order Traversal")
	printInOrder(tree)
	f()
	f(tree.search(123))

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

func (n *BinaryNode) search(key int) bool {
	count++
	if n == nil {
		return false
	}
	if n.data == key {
		return true
	}

	if n.data < key {
		return n.right.search(key)
	}

	if n.data > key {
		return n.left.search(key)
	}
	return false
}
func printPreOrder(n *BinaryNode) {
	f := fmt.Printf
	f("%v ", n.data)
	if n.left != nil {
		printPreOrder(n.left)
	}

	if n.right != nil {
		printPreOrder(n.right)
	}

}
func printPostOrder(n *BinaryNode) {
	f := fmt.Printf
	if n.left != nil {
		printPostOrder(n.left)
	}
	if n.right != nil {
		printPostOrder(n.right)
	}
	f("%v ", n.data)

}
func printInOrder(n *BinaryNode) {
	f := fmt.Printf
	if n.left != nil {
		printInOrder(n.left)
	}
	f("%v ", n.data)

	if n.right != nil {
		printInOrder(n.right)
	}

}
