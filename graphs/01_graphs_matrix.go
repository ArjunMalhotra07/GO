package graphs

import "fmt"

type GraphMatrix struct {
	Nodes [][]int
}

func (g *GraphMatrix) printGraphMatrix() {
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

func (g *GraphMatrix) addEdgeInMatrixGraph(sourceKey, destinationkey int) {
	g.Nodes[sourceKey][destinationkey] = 1
	g.Nodes[destinationkey][sourceKey] = 1
}
