package trees

import "fmt"

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

func printAns(incomingArray []*Node) {
	for i := 0; i < len(incomingArray); i++ {
		fmt.Print("Key ", incomingArray[i].Key, " Height ", incomingArray[i].Height)
		fmt.Println()
	}
}
