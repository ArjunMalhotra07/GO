package trees

import "fmt"

// ! Node
type Node struct {
	Key   int
	Left  *Node
	Right *Node
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

// ! Print Tree
func (t *Node) PreOrderTraversal(ans []int) []int {
	if t != nil {
		ans = append(ans, t.Key)
	}
	if t.Left != nil {
		ans = t.Left.PreOrderTraversal(ans)
	}
	if t.Right != nil {
		ans = t.Right.PreOrderTraversal(ans)
	}
	return ans
}
func (t *Node) PostOrderTraversal(ans []int) []int {
	if t.Left != nil {
		ans = t.Left.PostOrderTraversal(ans)
	}
	if t.Right != nil {
		ans = t.Right.PostOrderTraversal(ans)
	}
	if t != nil {
		ans = append(ans, t.Key)
	}
	return ans
}
func (t *Node) InOrderTraversal(ans []int) []int {

	if t.Left != nil {
		ans = t.Left.InOrderTraversal(ans)
	}
	if t != nil {
		ans = append(ans, t.Key)
	}
	if t.Right != nil {
		ans = t.Right.InOrderTraversal(ans)
	}
	return ans
}

func DoTreeOperations() {
	myTree := &Node{Key: 100}
	myTree.InsertIntotree(50)
	myTree.InsertIntotree(150)
	myTree.InsertIntotree(140)
	myTree.InsertIntotree(160)
	myTree.InsertIntotree(40)
	myTree.InsertIntotree(60)

	fmt.Println(myTree.PreOrderTraversal([]int{}))
	fmt.Println(myTree.InOrderTraversal([]int{}))
	fmt.Println(myTree.PostOrderTraversal([]int{}))
	fmt.Println(myTree.SearchValueInTree(160))

}
