package graphs

import "fmt"

func PerformKahnSort() {
	directedGraph := &DirectedGraph{}
	exampleGraphStruct := &ExampleGraphStruct{array: [][]int{{1, 2}, {2, 3}, {4, 2}, {3, 10}, {4, 7}, {7, 10}, {7, 11}, {1, 5}, {5, 6}, {4, 6}, {6, 8}, {8, 9}}, maxVerticesCount: 11}
	directedGraph.makeDirectedMatrixGraph(*exampleGraphStruct)
	directedGraph.KahnAlgo()
}

func (g *DirectedGraph) KahnAlgo() {
	indegree := make([]int, len(g.Nodes))
	for i := 0; i < len(g.Nodes); i++ {
		for j := 0; j < len(g.Nodes[0]); j++ {
			if g.Nodes[i][j] == 1 {
				indegree[j] += 1
			}
		}
	}
	//! Using this becuase I only need q with []int; not [][]int
	q := &ProvinceQueue{}
	for i := 1; i < len(indegree); i++ {
		if indegree[i] == 0 {
			q.Enqueue(i)
		}
	}
	sorted := []int{}
	for len(q.nums) != 0 {
		popped := q.Dequeue()
		sorted = append(sorted, popped)
		for i := 0; i < len(g.Nodes[popped]); i++ {
			if g.Nodes[popped][i] == 1 {
				indegree[i] -= 1
				if indegree[i] == 0 {
					q.Enqueue(i)
				}
			}
		}
	}
	fmt.Println("Kahn Algorithm", sorted)
}
