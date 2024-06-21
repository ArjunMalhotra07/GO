package graphs

import "fmt"

func PerformDFS() {
	exampleGraphStruct := &ExampleGraphStruct{array: [][]int{{1, 2}, {1, 6}, {2, 3}, {2, 4}, {6, 7}, {6, 8}, {4, 5}, {7, 5}}, maxVerticesCount: 8}
	listGraph := &Graph{}
	listGraph.makeListGraph(*exampleGraphStruct)
	listGraph.printGraphAdjacencyList()
	DepthFirstSearch(listGraph)
}
func DepthFirstSearch(graph *Graph) {
	if len(graph.Vertices) == 0 {
		return
	}
	visitedArray := make([]int, len(graph.Vertices)+1)
	myStack := &Stack{}
	myStack.Push(graph.Vertices[0])
	ansArray := []int{}
	for len(myStack.vertices) != 0 {
		poppedVertex := myStack.Pop()
		if visitedArray[poppedVertex.Key] != 1 {
			visitedArray[poppedVertex.Key] = 1
			ansArray = append(ansArray, poppedVertex.Key)
			for _, v := range poppedVertex.Neighbours {
				if visitedArray[v.Key] != 1 {
					myStack.Push(v)
				}
			}
		}
	}
	fmt.Println("DFS:", ansArray)
}

type Stack struct {
	vertices []*Vertex
}

func (s *Stack) Push(vertex *Vertex) {
	s.vertices = append(s.vertices, vertex)
}

func (s *Stack) Pop() *Vertex {
	length := len(s.vertices)
	if length >= 0 {
		removedvertex := s.vertices[length-1]
		s.vertices = s.vertices[:length-1]
		return removedvertex
	} else {
		fmt.Println("Empty Stack")
	}
	return nil
}
