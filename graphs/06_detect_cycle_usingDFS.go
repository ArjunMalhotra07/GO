package graphs

func DetectCycleUsingDFS(graph *Graph) bool {
	if len(graph.Vertices) == 0 {
		return false
	}
	visitedArray := make([]int, len(graph.Vertices)+1)
	myStack := &Stack{}
	myStack.Push(graph.Vertices[0])
	for len(myStack.vertices) != 0 {
		poppedVertex := myStack.Pop()
		if visitedArray[poppedVertex.Key] != 1 {
			visitedArray[poppedVertex.Key] = 1
			for _, v := range poppedVertex.Neighbours {
				if visitedArray[v.Key] != 1 {
					myStack.Push(v)
				}
			}
		} else {
			return true
		}
	}
	return false
}
