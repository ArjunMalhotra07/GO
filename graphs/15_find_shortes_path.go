package graphs

import (
	"fmt"
	"math"
)

func FindShortestPath() {
	// ! Initialisation of Graph Vars
	matrixGraph := &GraphMatrix{}
	//! From, To, Weight
	exampleGraphStruct := &ExampleGraphStruct{array: [][]int{{1, 2, 4}, {1, 3, 4}, {2, 3, 2}, {3, 4, 3}, {3, 5, 1}, {3, 6, 6}, {4, 6, 2}, {5, 6, 3}}, maxVerticesCount: 6} //! Weighted Graphs
	// exampleGraphStruct := &ExampleGraphStruct{array: [][]int{{1, 2, 7}, {1, 3, 9}, {1, 6, 14}, {3, 6, 2}, {3, 2, 10}, {2, 4, 15}, {3, 4, 11}, {5, 4, 6}, {6, 5, 9}}, maxVerticesCount: 6}
	matrixGraph.Nodes = make([][]int, exampleGraphStruct.maxVerticesCount+1)

	for i := 0; i < exampleGraphStruct.maxVerticesCount+1; i++ {
		matrixGraph.Nodes[i] = make([]int, exampleGraphStruct.maxVerticesCount+1)
	}
	for i := 0; i < len(exampleGraphStruct.array); i++ {
		matrixGraph.addEdgeWeightInMatrixGraph(exampleGraphStruct.array[i][0], exampleGraphStruct.array[i][1], exampleGraphStruct.array[i][2])
	}
	// ! Print Graph Info
	matrixGraph.printGraphMatrix()
	// ! Finding Minimum Distance of each node from a source Vertex
	source := 1
	destination := 6
	fmt.Println(findMinPath(matrixGraph.Nodes, source, destination))
}
func findMinPath(g [][]int, source, destination int) []int {
	minDist := make([]int, len(g))
	isNodeAlreadyRelaxed := make([]bool, len(g))
	parent := make([]int, len(g))
	for i := 0; i < len(parent); i++ {
		parent[i] = i
	}
	ans := []int{}

	for i := 1; i < len(minDist); i++ {
		minDist[i] = math.MaxInt64
	}
	minDist[source] = 0
	h := &Heap{}
	if len(g) == 0 {
		return minDist
	}

	//! Adding {Distance, Node} in Min Heap
	h.addInHeap([]int{0, source})
	for len(h.heap) != 0 {
		root := h.getTopElementFromHeap()
		fmt.Println("relaxed", root)
		if !isNodeAlreadyRelaxed[root[1]] {
			isNodeAlreadyRelaxed[root[1]] = true
			for node := 1; node < len(g); node++ {
				//! Looping through each edge, getting its weight & updating dist in ans array if its less than before
				cellValue := g[root[1]][node]
				newDistance := cellValue + root[0]
				if cellValue != 0 && newDistance < minDist[node] {
					minDist[node] = newDistance
					h.addInHeap([]int{newDistance, node})
					parent[node] = root[1]
				}
			}
		}
	}
	ans = append(ans, destination)
	for parent[destination] != destination {
		ans = append(ans, parent[destination])
		destination = parent[destination]
	}
	fmt.Println(ans)
	return minDist
}
