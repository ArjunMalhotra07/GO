package graphs

import "fmt"

func PerformBellmanFord() {
	exampleGraphStruct := &ExampleGraphStruct{array: [][]int{{0, 1, 5}, {1, 5, -3}, {5, 3, 1}, {1, 2, -2}, {3, 2, 6}, {2, 4, 3}, {3, 4, -2}}, maxVerticesCount: 6}
	bellmanFord(exampleGraphStruct, 0)

}

func bellmanFord(g *ExampleGraphStruct, source int) {
	fmt.Println(g.array)
	fmt.Println(g.maxVerticesCount)
	distArray := make([]int, g.maxVerticesCount)
	for i := 0; i < g.maxVerticesCount; i++ {
		distArray[i] = 1e18
	}
	distArray[source] = 0

	for i := 0; i < g.maxVerticesCount; i++ {
		for _, edge := range g.array {
			u := edge[0]
			v := edge[1]
			wt := edge[2]
			if distArray[u] != 1e18 && distArray[u]+wt < distArray[v] {
				distArray[v] = distArray[u] + wt
			}

		}
	}
	fmt.Println(distArray)
}
