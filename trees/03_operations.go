package trees

import (
	"fmt"
	"math"
)

func getBalanceFactor(t *Node) float64 {
	if t == nil {
		return 0
	} else {
		return getHeight(t.Left) - getHeight(t.Right)
	}
}
func getHeight(t *Node) float64 {
	if t == nil {
		return 0
	} else {
		return t.Height
	}
}

// ! Insert
func (t *Node) InsertIntotree(value int) *Node {

	if value < t.Key {
		//! Move left
		if t.Left == nil {
			t.Left = &Node{Key: value}
		} else {
			t.Left.InsertIntotree(value)
		}
	} else if value > t.Key {
		//! Move Right
		if t.Right == nil {
			t.Right = &Node{Key: value}
		} else {
			t.Right.InsertIntotree(value)
		}
	}
	//! Height of Node where Key is added and not the Key Node (It's already zero)
	height := math.Max(getHeight(t.Left), getHeight(t.Right)) + 1
	t.Height = height
	// balanceFactor := getBalanceFactor(t)
	// if balanceFactor > 1 {
	// 	fmt.Println("here")
	// 	if value < t.Left.Key {
	// 		return t.rightRotate()
	// 	} else if value > t.Left.Key {
	// 		t.Left = t.Left.leftRotate()
	// 		return t.rightRotate()
	// 	}
	// }
	// if balanceFactor < -1 {
	// 	if value > t.Right.Key {
	// 		return t.leftRotate()
	// 	} else if value < t.Right.Key {
	// 		t.Right = t.rightRotate()
	// 		return t.leftRotate()
	// 	}
	// }
	return t
}

func (y *Node) rightRotate() *Node {
	x := y.Left
	t2 := x.Right
	x.Right = y
	y.Left = t2
	y.Height = math.Max(getHeight(y.Left), getHeight(y.Right))
	x.Height = math.Max(getHeight(x.Left), getHeight(x.Right)) + 1
	return x
}
func (x *Node) leftRotate() *Node {
	y := x.Right
	t2 := y.Left
	y.Left = x
	x.Right = t2
	y.Height = math.Max(getHeight(y.Left), getHeight(y.Right)) + 1
	x.Height = math.Max(getHeight(x.Left), getHeight(x.Right))
	return y
}

// ! Search
func (t *Node) SearchValueInTree(n int) bool {
	if t == nil {
		return false
	}
	fmt.Println("NODE VALUE", t.Key)
	if n < t.Key {
		return t.Left.SearchValueInTree(n)
	} else if n > t.Key {
		return t.Right.SearchValueInTree(n)
	} else {
		return true
	}
}
