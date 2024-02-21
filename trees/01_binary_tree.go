package trees

import (
	"fmt"
	"math"
)

// ! Node
type Node struct {
	Key    int
	Left   *Node
	Right  *Node
	Height float64
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
	t.Height = math.Max(height(t.Left), height(t.Right)) + 1
	balanceFactor := getBalanceFactor(t)
	fmt.Println("Balance Factor", balanceFactor, " Key", t.Key)
}
func getBalanceFactor(t *Node) float64 {
	if t == nil {
		return 0
	} else {
		return height(t.Left) - height(t.Right)
	}
}
func height(t *Node) float64 {
	if t == nil {
		return 0
	} else {
		return t.Height
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
func (t *Node) PreOrderTraversal(ans []*Node) []*Node {
	if t != nil {
		ans = append(ans, t)
	}
	if t.Left != nil {
		ans = t.Left.PreOrderTraversal(ans)
	}
	if t.Right != nil {
		ans = t.Right.PreOrderTraversal(ans)
	}
	return ans
}
func (t *Node) PostOrderTraversal(ans []*Node) []*Node {
	if t.Left != nil {
		ans = t.Left.PostOrderTraversal(ans)
	}
	if t.Right != nil {
		ans = t.Right.PostOrderTraversal(ans)
	}
	if t != nil {
		ans = append(ans, t)
	}
	return ans
}
func (t *Node) InOrderTraversal(ans []*Node) []*Node {

	if t.Left != nil {
		ans = t.Left.InOrderTraversal(ans)
	}
	if t != nil {
		ans = append(ans, t)
	}
	if t.Right != nil {
		ans = t.Right.InOrderTraversal(ans)
	}
	return ans
}

func DoTreeOperations() {

	myTree := &Node{Key: 3}
	arrayInput(myTree)
	fmt.Println("Preorder")
	printAns(myTree.PreOrderTraversal([]*Node{}))
	fmt.Println("Inorder")
	printAns(myTree.InOrderTraversal([]*Node{}))
	fmt.Println("Postorder")
	printAns(myTree.PostOrderTraversal([]*Node{}))
	fmt.Println(myTree.SearchValueInTree(160))

}

func arrayInput(treeNode *Node) {
	arr := []int{2, 1}
	for i := 0; i < len(arr); i++ {
		treeNode.InsertIntotree(arr[i])
	}
}

func printAns(incomingArray []*Node) {
	for i := 0; i < len(incomingArray); i++ {
		fmt.Print("Key ", incomingArray[i].Key, " Height ", incomingArray[i].Height)
		fmt.Println()
	}
}
