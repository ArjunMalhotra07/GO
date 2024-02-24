package trees

import "fmt"

// ! Node
type Node struct {
	Key    int
	Left   *Node
	Right  *Node
	Height float64
}

func DoTreeOperations() {
	myTree := &Node{Key: 5}
	arrayInput(myTree)
	fmt.Println("Preorder")
	printAns(myTree.PreOrderTraversal([]*Node{}))
	fmt.Println("Inorder")
	printAns(myTree.InOrderTraversal([]*Node{}))
	fmt.Println("Postorder")
	printAns(myTree.PostOrderTraversal([]*Node{}))
	fmt.Println(myTree.SearchValueInTree(160))

}

func arrayInput(root *Node) {
	arr := []int{4, 3, 2, 1}
	for i := 0; i < len(arr); i++ {
		root = root.InsertIntotree(arr[i])
		fmt.Println(root.Key)
	}
	// root = root.InsertIntotree(2)
	// root.InsertIntotree(1)

}
