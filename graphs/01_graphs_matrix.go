package graphs

import "fmt"

type GraphMatrix struct {
	Nodes [][]int
}

func (g *GraphMatrix) printGraphMatrix() {
	row := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Print("  ", row)
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
