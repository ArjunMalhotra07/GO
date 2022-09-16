//  AVL trees are height balancing binary search tree.
//  AVL tree checks the height of the left and the
//  right sub-trees and assures that the difference is
//  not more than 1. This difference is called the Balance Factor.
//Time Complexity O(n)
package main

import (
	"fmt"
	"math"
)

func main() {
	f := fmt.Println
	tree := &BinaryNode{data: 4, height: 1, left: nil, right: nil}
	//Balanced Tree
	tree.insert(2)
	tree.insert(1)
	tree.insert(3)
	tree.insert(6)
	tree.insert(5)
	tree.insert(7)
	//Unbalanced
	// tree.insert(5)
	// tree.insert(6)
	// tree.insert(7)
	// tree.insert(3)

	height := 0.0
	if isBalanced2(tree, &height) {
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
	} else {
		//move left
		if n.left == nil {
			n.left = &BinaryNode{data: key}

		} else {
			n.left.insert(key)
		}
	}

}

func isBalanced2(n *BinaryNode, height *float64) bool {
	lh := 0.0
	rh := 0.0
	if n == nil {
		return true
	}

	if isBalanced2(n.left, &lh) == false {
		return false
	}
	if isBalanced2(n.right, &rh) == false {
		return false
	}
	*height = math.Max(lh, rh) + 1

	var f float64 = float64(lh - rh)
	if math.Abs(f) <= 1 {

		return true
	} else {
		return false
	}

}
