//  AVL trees are height balancing binary search tree.
//  AVL tree checks the height of the left and the
//  right sub-trees and assures that the difference is
//  not more than 1. This difference is called the Balance Factor.

//Time Complexity O(n^2)
package main

import (
	"fmt"
)

func main() {

	root := &BinaryNode{data: -10}
	root = insert2(root, -3)
	root = insert2(root, 0)
	root = insert2(root, 5)

	root = insert2(root, 9)
	// root = insert2(root, -11)
	// root = insert2(root, 11)
	// root = insert2(root, 118)
	// root = insert2(root, 500)
	// root = insert2(root, 121)
	// root = insert2(root, 117)
	printPreOrder1(root)

}
func printPreOrder1(n *BinaryNode) {
	f := fmt.Printf
	f("%v ", n.data)
	if n.left != nil {
		printPreOrder1(n.left)
	}

	if n.right != nil {
		printPreOrder1(n.right)
	}

}

type BinaryNode struct {
	root   *BinaryNode
	data   int
	left   *BinaryNode
	right  *BinaryNode
	height int
}

func insert2(root *BinaryNode, key int) *BinaryNode {
	current := root
	if current == nil {
		return &BinaryNode{data: key}

	} else {
		if key > current.data {
			//move right
			if current.right == nil {
				current.right = &BinaryNode{data: key}
				return root
			} else {
				insert2(current.right, key)
			}
		} else if key < current.data {
			//move left
			if current.left == nil {
				current.left = &BinaryNode{data: key}
				return root
			} else {
				insert2(current.left, key)
			}
		} else {
			fmt.Println("Duplicates Not Allowed")
			return root
		}

	}
	root.height = 1 + max(int(height(root.left)), int(height(root.right)))
	bf := getBalanceFactor(current)
	if bf > 1 && key < root.left.data {
		return rightRotate(root)
	}
	if bf < -1 && key > root.right.data {
		return leftRotate(root)
	}
	if bf > 1 && key > root.left.data {
		root.left = leftRotate(root.left)
		rightRotate(root)
	}
	if bf < -1 && key < root.right.data {
		root.right = rightRotate(root.right)
		leftRotate(root)
	}

	return root
}

func height(n *BinaryNode) int {
	if n == nil {
		return 0
	}
	return int(n.height)

}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
func rightRotate(y *BinaryNode) *BinaryNode {
	var x *BinaryNode = y.left
	var t2 *BinaryNode = x.right
	x.right = y
	y.left = t2
	y.height = max(int(height(y.left)), int(height(y.right))) + 1
	x.height = max(int(height(x.left)), int(height(x.right))) + 1
	return x
}

func leftRotate(x *BinaryNode) *BinaryNode {
	y := x.right
	t2 := y.left
	y.left = x
	x.right = t2
	x.height = max(int(height(x.left)), int(height(x.right))) + 1
	y.height = max(int(height(y.left)), int(height(y.right))) + 1
	return y
}

func getBalanceFactor(current *BinaryNode) int {
	if current == nil {
		return 0
	} else {
		return int(height(current.left)) - int(height(current.right))
	}

}
