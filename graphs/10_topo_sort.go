package graphs

import "fmt"

func (g *DirectedGraph) TopoLogicalSort() {
	ans := []int{}
	s := &S{}
	visited := make([]bool, len(g.Nodes))
	for i := 1; i < len(g.Nodes); i++ {
		// fmt.Println(i, visited[i])
		if !visited[i] {
			dfs(i, g, &visited, s)
			s.push(i)
		}
	}
	for len(s.vertices) != 0 {
		ans = append(ans, s.pop())
	}
	fmt.Println("Topological sort order", ans)
}
func dfs(node int, g *DirectedGraph, visited *[]bool, stack *S) {

	(*visited)[node] = true
	for j := 0; j < len(g.Nodes[node]); j++ {
		if g.Nodes[node][j] == 1 && !(*visited)[j] {
			dfs(j, g, visited, stack)
			stack.push(j)
			// fmt.Println("Pushed", j)
		}
	}
}

type S struct {
	vertices []int
}

func (s *S) push(vertex int) {
	s.vertices = append(s.vertices, vertex)
}
func (s *S) pop() int {
	length := len(s.vertices)
	if length > 0 {
		popped := s.vertices[length-1]
		if length != 1 {
			s.vertices = s.vertices[:length-1]
		} else {
			s.vertices = []int{}
		}
		return popped
	}
	return -1
}
