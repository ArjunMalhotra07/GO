package graphs

import "fmt"

func PerformGraphsOps() {
	myGraph := &Graph{}
	matrixGraph := &GraphMatrix{}
	makeGraph(myGraph, matrixGraph)
	myGraph.printGraphAdjacencyList()
	matrixGraph.printGraphMatrix()

	BreadthFirstSearch(myGraph)
	DepthFirstSearch(myGraph)
	fmt.Println(GetNumberOfProvinces(matrixGraph.Nodes))
}

func makeGraph(g *Graph, matrixGraph *GraphMatrix) {
	// exampleGraphStruct := &ExampleGraphStruct{array: [][]int{{1, 2}, {1, 6}, {2, 3}, {2, 4}, {6, 7}, {6, 8}, {4, 5}, {7, 5}}, maxVerticesCount: 8}
	// exampleGraphStruct := &ExampleGraphStruct{array: [][]int{{1, 2}, {1, 3}, {2, 5}, {2, 6}, {3, 4}, {3, 7}, {7, 8}, {4, 8}}, maxVerticesCount: 8}
	// exampleGraphStruct := &ExampleGraphStruct{array: [][]int{{1, 2}, {2, 3}, {4, 5}, {5, 6}, {7, 8}}, maxVerticesCount: 8} //!Provinces Example Graph
	// exampleGraphStruct := &ExampleGraphStruct{array: [][]int{{1, 2}, {2, 3}, {4, 1}, {4, 3}, {5, 6}, {7, 6}, {8, 9}, {9, 10}, {11, 12}}, maxVerticesCount: 13} //!Provinces Example Graph
	exampleGraphStruct := &ExampleGraphStruct{array: [][]int{{1, 1}, {2, 2}, {3, 3}}, maxVerticesCount: 3} //!Provinces Example Graph

	matrixGraph.Nodes = make([][]int, exampleGraphStruct.maxVerticesCount+1)
	for i := 1; i <= exampleGraphStruct.maxVerticesCount; i++ {
		g.addVertexMethod(i)
	}
	for i := 0; i < exampleGraphStruct.maxVerticesCount+1; i++ {
		matrixGraph.Nodes[i] = make([]int, exampleGraphStruct.maxVerticesCount+1)
	}
	for i := 0; i < len(exampleGraphStruct.array); i++ {
		g.addEdge(exampleGraphStruct.array[i][0], exampleGraphStruct.array[i][1])
		matrixGraph.addEdgeInMatrixGraph(exampleGraphStruct.array[i][0], exampleGraphStruct.array[i][1])
	}

}

type ExampleGraphStruct struct {
	array            [][]int
	maxVerticesCount int
}
