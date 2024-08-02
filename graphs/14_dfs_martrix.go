package graphs

import "fmt"

func DFSMatrix() {
	exampleGraphStruct := &ExampleGraphStruct{array: [][]int{{1, 2}, {1, 6}, {2, 3}, {2, 4}, {6, 7}, {6, 8}, {4, 5}, {7, 5}}, maxVerticesCount: 8}
	graph := &GraphMatrix{}
	graph.makeMatrixGraph(*exampleGraphStruct)
	graph.printGraphMatrix()
	myQ := &stack{}
	visited := make([]bool, len(graph.Nodes))
	myQ.push(1)
	ans := []int{}

	for len(myQ.vertices) != 0 {
		popped := myQ.pop()
		if !visited[popped] {
			ans = append(ans, popped)
			visited[popped] = true
			for node, isNeighbour := range graph.Nodes[popped] {
				if isNeighbour == 1 && !visited[node] {
					myQ.push(node)
				}
			}
		}
	}
	fmt.Println(ans)
}

type stack struct {
	vertices []int
}

func (s *stack) push(vertex int) {
	s.vertices = append(s.vertices, vertex)
}
func (s *stack) pop() int {
	if len(s.vertices) == 0 {
		return -1
	}
	lastElement := s.vertices[len(s.vertices)-1]
	if len(s.vertices) != 1 {
		s.vertices = s.vertices[:len(s.vertices)-1]
	} else {
		s.vertices = []int{}
	}
	return lastElement
}
