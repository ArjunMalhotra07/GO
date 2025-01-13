package trees

import (
	"math"
)

func getBalanceFactor(t *Node) float64 {
	return getHeight(t.Left) - getHeight(t.Right)
}
func getHeight(t *Node) float64 {
	if t == nil {
		return 0
	} else {
		return t.Height
	}
}

// ! Insert
func (root *Node) InsertIntotree(value int) *Node {

	if value < root.Key {
		//! Move left
		if root.Left == nil {
			root.Left = &Node{Key: value, Height: 1}
		} else {
			root.Left = root.Left.InsertIntotree(value)
		}
	} else if value > root.Key {
		//! Move Right
		if root.Right == nil {
			root.Right = &Node{Key: value, Height: 1}
		} else {
			root.Right = root.Right.InsertIntotree(value)
		}
	}
	//! Height of Node where Key is added and not the Key Node (It's already zero)
	height := math.Max(getHeight(root.Left), getHeight(root.Right)) + 1
	root.Height = height
	balanceFactor := getBalanceFactor(root)
	if balanceFactor > 1 {
		if value < root.Left.Key {
			return root.rightRotate()
		} else if value > root.Left.Key {
			root.Left = root.Left.leftRotate()
			return root.rightRotate()
		}
	}
	if balanceFactor < -1 {
		if value > root.Right.Key {
			return root.leftRotate()
		} else if value < root.Right.Key {
			root.Right = root.rightRotate()
			return root.leftRotate()
		}
	}
	return root
}

func (y *Node) rightRotate() *Node {
	x := y.Left
	t2 := x.Right
	x.Right = y
	y.Left = t2
	y.Height = math.Max(getHeight(y.Left), getHeight(y.Right)) + 1
	x.Height = math.Max(getHeight(x.Left), getHeight(x.Right)) + 1
	return x
}
func (x *Node) leftRotate() *Node {
	y := x.Right
	t2 := y.Left
	y.Left = x
	x.Right = t2
	y.Height = math.Max(getHeight(y.Left), getHeight(y.Right)) + 1
	x.Height = math.Max(getHeight(x.Left), getHeight(x.Right)) + 1
	return y
}

// ! Search
func (t *Node) SearchValueInTree(n int) bool {
	if t == nil {
		return false
	}
	if n < t.Key {
		return t.Left.SearchValueInTree(n)
	} else if n > t.Key {
		return t.Right.SearchValueInTree(n)
	} else {
		return true
	}
}
