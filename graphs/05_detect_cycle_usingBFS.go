package graphs

import "fmt"

func PerformCycleDetectionUsingBFS() {
	//!Cycle Example Graph
	exampleGraphStruct := &ExampleGraphStruct{array: [][]int{{1, 2}, {1, 3}, {3, 4}, {3, 6}, {6, 7}, {7, 5}, {2, 5}}, maxVerticesCount: 7}
	// exampleGraphStruct := &ExampleGraphStruct{array: [][]int{{1, 2}, {1, 4}, {4, 3}, {2, 3}}, maxVerticesCount: 4} //!Cycle Example Graph
	listGraph := &Graph{}
	listGraph.makeListGraph(*exampleGraphStruct)
	listGraph.printGraphAdjacencyList()
	fmt.Println(listGraph.DetectCycleUsingBFS())
}
func (graph *Graph) DetectCycleUsingBFS() bool {
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
