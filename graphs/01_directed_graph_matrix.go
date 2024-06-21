package graphs

import "fmt"

type DirectedGraph struct {
	Nodes [][]int
}

func (g *DirectedGraph) PrintGraphMatrix() {
	dummy := make([]int, len(g.Nodes))
	for i := 0; i < len(g.Nodes); i++ {
		dummy[i] = i
	}
	fmt.Print("  ", dummy)
	fmt.Println()
	for i := 0; i < len(g.Nodes); i++ {
		fmt.Print(i, g.Nodes[i])
		fmt.Println()
	}
}

func (g *DirectedGraph) AddEdgeInDirectedMatrixGraph(sourceKey, destinationkey int) {
	g.Nodes[sourceKey][destinationkey] = 1
}

func (g *DirectedGraph) AddEdgeWeightInDirectedMatrixGraph(sourceKey, destinationkey, weight int) {
	g.Nodes[sourceKey][destinationkey] = weight
	g.Nodes[destinationkey][sourceKey] = weight
}
