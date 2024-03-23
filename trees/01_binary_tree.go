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

}

func arrayInput(root *Node) {
	arr := []int{4, 3, 2, 1}
	for i := 0; i < len(arr); i++ {
		root = root.InsertIntotree(arr[i])
	}
	fmt.Println("Preorder")
	printAns(root.PreOrderTraversal([]*Node{}))
	fmt.Println("Inorder")
	printAns(root.InOrderTraversal([]*Node{}))
	// fmt.Println("Inorder")
	// printAns(root.InOrderTraversal2())

	fmt.Println("Postorder")
	printAns(root.PostOrderTraversal([]*Node{}))
	fmt.Println(root.SearchValueInTree(160))

}
