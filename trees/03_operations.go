package trees

import (
	"fmt"
	"math"
)

func getBalanceFactor(t *Node) float64 {
	if t == nil {
		return 0
	} else {
		fmt.Println(t.Key)

		fmt.Println(getHeight(t))
		fmt.Println(getHeight(t))
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
func (t *Node) InsertIntotree(value int) {

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
	height := math.Max(getHeight(t.Left), getHeight(t.Right)) + 1
	t.Height = height
	fmt.Println("Inserted", value, "Height", height)

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
