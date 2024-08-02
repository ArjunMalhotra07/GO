package graphs

import "fmt"

func BFSMatrix() {
	exampleGraphStruct := &ExampleGraphStruct{array: [][]int{{1, 2}, {1, 6}, {2, 3}, {2, 4}, {6, 7}, {6, 8}, {4, 5}, {7, 5}}, maxVerticesCount: 8}
	graph := &GraphMatrix{}
	graph.makeMatrixGraph(*exampleGraphStruct)
	graph.printGraphMatrix()
	myQ := &q{}
	visited := make([]bool, len(graph.Nodes))
	myQ.enQ(1)
	ans := []int{}

	for len(myQ.vertices) != 0 {
		pop := myQ.deQ()
		if !visited[pop] {
			ans = append(ans, pop)
			visited[pop] = true
			for node, isNeighbour := range graph.Nodes[pop] {
				if isNeighbour == 1 && !visited[node] {
					myQ.enQ(node)
				}
			}
		}
	}
	fmt.Println(ans)
}

type q struct {
	vertices []int
}

func (qu *q) enQ(vertex int) {
	qu.vertices = append(qu.vertices, vertex)
}
func (qu *q) deQ() int {
	if len(qu.vertices) == 0 {
		return -1
	}
	firstElement := qu.vertices[0]
	if len(qu.vertices) != 1 {
		qu.vertices = qu.vertices[1:]
	} else {
		qu.vertices = []int{}
	}
	return firstElement
}
