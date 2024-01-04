package graphs

import (
	"fmt"
)

type Graph struct {
	Vertices []*Vertex
}
type Vertex struct {
	Key        int
	Neighbours []*Vertex
}

// ! Adds Vertex to the Graph
func (g *Graph) addVertexMethod(value int) {
	if containsVertex(g.Vertices, value) {
		err := fmt.Errorf("cant add vertex %v as it already exists", value)
		fmt.Println(err.Error())
	} else {
		dummyVertex := &Vertex{Key: value}
		g.Vertices = append(g.Vertices, dummyVertex)
	}
}

// ! Adds Edge from Vertex A -> B
func (g *Graph) addEdge(source, destination int) {
	//! Get actual vertices (from, to)
	sourceVertex := getVertex(source, g.Vertices)
	destinationVertex := getVertex(destination, g.Vertices)
	//! Check for errors
	//! Is source or destination nil   && if edge already exists
	if sourceVertex == nil || destinationVertex == nil {
		fmt.Printf("Invalid Edge (%v - %v)\n", source, destination)
		return
	} else if containsVertex(sourceVertex.Neighbours, destination) {
		fmt.Printf("Existing Edge (%v - %v)\n", source, destination)
		return
	} else {
		sourceVertex.Neighbours = append(sourceVertex.Neighbours, destinationVertex)
		destinationVertex.Neighbours = append(destinationVertex.Neighbours, sourceVertex)
	}
}

// ! Check if Graph already has a key
func containsVertex(s []*Vertex, key int) bool {
	for _, vertex := range s {
		if key == vertex.Key {
			return true
		}
	}
	return false
}

// ! Returns Vertex Pointer that matches given Key
func getVertex(key int, vertices []*Vertex) *Vertex {
	for _, v := range vertices {
		if v.Key == key {
			return v
		}
	}
	return nil
}

// ! Prints Graph contents
func (g *Graph) printGraphAdjacencyList() {
	for _, vertex := range g.Vertices {
		fmt.Printf("\n Vertex %v:", vertex.Key)
		for _, neighbour := range vertex.Neighbours {
			fmt.Printf(" %v", neighbour.Key)
		}
	}
	fmt.Println()
}

func PerformGraphsOps() {
	myGraph := &Graph{}
	makeGraph(myGraph)
	myGraph.printGraphAdjacencyList()
	BreadthFirstSearch(myGraph)
}

func makeGraph(g *Graph) {
	exampleGraphStruct := &ExampleGraphStruct{array: [][]int{{1, 2}, {1, 6}, {2, 3}, {2, 4}, {6, 7}, {6, 8}, {4, 5}, {7, 5}}, maxVerticesCount: 8}
	// exampleGraphStruct := &ExampleGraphStruct{array: [][]int{{1, 3}, {1, 2}, {3, 4}, {2, 4}, {5, 6}, {6, 7}, {5, 7}, {8, 9}}, maxVerticesCount: 10}
	for i := 1; i <= exampleGraphStruct.maxVerticesCount; i++ {
		g.addVertexMethod(i)
	}
	for i := 0; i < len(exampleGraphStruct.array); i++ {
		g.addEdge(exampleGraphStruct.array[i][0], exampleGraphStruct.array[i][1])
	}

}

type ExampleGraphStruct struct {
	array            [][]int
	maxVerticesCount int
}
