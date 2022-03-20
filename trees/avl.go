//  AVL trees are height balancing binary search tree.
//  AVL tree checks the height of the left and the
//  right sub-trees and assures that the difference is
//  not more than 1. This difference is called the Balance Factor.

//Time Complexity O(n^2)
package main

import (
	"fmt"
	"math"
)

func main() {
	f := fmt.Println
	tree := &BinaryNode{data: 4, height: 1, left: nil, right: nil}
	tree.insert(2)
	tree.insert(1)
	tree.insert(3)
	tree.insert(6)
	tree.insert(5)
	tree.insert(7)
	tree.insert(4)

	if isBalanced(tree) {
		f("Balanced")
	} else {
		f("Not Balanced")
	}
}

type BinaryNode struct {
	data   int
	left   *BinaryNode
	right  *BinaryNode
	height int
}

func (n *BinaryNode) insert(key int) {

	if n.data < key {
		//move right
		if n.right == nil {
			n.right = &BinaryNode{data: key}

		} else {
			n.right.insert(key)
		}
	} else if n.data > key {
		//move left
		if n.left == nil {
			n.left = &BinaryNode{data: key}

		} else {
			n.left.insert(key)
		}
	} else {
		fmt.Println("Duplicates Not Allowed")
	}

}

func isBalanced(n *BinaryNode) bool {
	if n == nil {
		return true
	}

	if isBalanced(n.left) == false {
		return false
	}
	if isBalanced(n.right) == false {
		return false
	}
	lh := height(n.left)
	rh := height(n.right)
	var f float64 = float64(lh - rh)
	if math.Abs(f) <= 1 {
		return true
	} else {
		return false
	}

}

func height(n *BinaryNode) float64 {
	if n == nil {
		return 0
	}
	lh := height(n.left)
	rh := height(n.right)
	return math.Max(lh, rh) + 1
}
