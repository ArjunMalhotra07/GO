package graphs

import "fmt"

func DetectCycleUsingBFS(graph *Graph) bool {
	if len(graph.Vertices) == 0 {
		fmt.Println("Empty graph")
		return false
	}

	visitedArray := make([]int, len(graph.Vertices)+1)
	myQueue := &Queue{}
	myQueue.Enqueue(graph.Vertices[0])
	for len(myQueue.vertices) != 0 {
		node := myQueue.DeQueue()
		if visitedArray[node.Key] != 1 {
			visitedArray[node.Key] = 1
			for _, v := range node.Neighbours {
				if visitedArray[v.Key] != 1 {
					myQueue.Enqueue(v)
				}
			}
		} else {
			return true
		}
	}
	return false
}
