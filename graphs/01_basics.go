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
	for i := 0; i < 5; i++ {
		myGraph.addVertexMethod(i)
	}
	myGraph.addVertexMethod(3)
	myGraph.addEdge(0, 4)
	myGraph.addEdge(0, 4)
	myGraph.addEdge(1, 6)
	myGraph.addEdge(3, 2)
	myGraph.printGraphAdjacencyList()

}
