package trees

import (
	"fmt"
)

// ! Node
type Node struct {
	Key    int
	Left   *Node
	Right  *Node
	Height float64
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
	// balanceFactor := getBalanceFactor(myTree)
	// fmt.Println("Balance Factor", balanceFactor)
}

func arrayInput(treeNode *Node) {
	arr := []int{2, 1}
	for i := 0; i < len(arr); i++ {
		treeNode.InsertIntotree(arr[i])
	}
}
