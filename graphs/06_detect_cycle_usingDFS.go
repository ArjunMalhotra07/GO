package graphs

import "fmt"

func PerformCycleDetectionUsingDFS() {
	//!Cycle Example Graph
	exampleGraphStruct := &ExampleGraphStruct{array: [][]int{{1, 2}, {1, 3}, {3, 4}, {3, 6}, {6, 7}, {7, 5}, {2, 5}}, maxVerticesCount: 7}
	// exampleGraphStruct := &ExampleGraphStruct{array: [][]int{{1, 2}, {1, 4}, {4, 3}, {2, 3}}, maxVerticesCount: 4} //!Cycle Example Graph
	listGraph := &Graph{}
	listGraph.makeListGraph(*exampleGraphStruct)
	listGraph.printGraphAdjacencyList()
	fmt.Println(listGraph.DetectCycleUsingDFS())
}
func (graph *Graph) DetectCycleUsingDFS() bool {
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
