package graphs

import "fmt"

func PerformTopoSort() {
	// arr := [][]int{{0, 0, 0, 0, 0, 0}, {0, 0, 0, 0, 0, 0}, {0, 0, 0, 1, 0, 0}, {0, 1, 0, 0, 0, 0}, {1, 1, 0, 0, 0, 0}, {1, 0, 1, 0, 0, 0}} //! Non cyclic example
	arr := [][]int{{0, 0, 0, 0, 0, 1}, {0, 0, 0, 1, 0, 0}, {0, 0, 0, 1, 0, 0}, {1, 0, 0, 0, 0, 0}, {1, 1, 0, 0, 0, 0}, {0, 0, 1, 0, 0, 0}} //! Cyclic example

	vis := make([]bool, len(arr))

	s := &_S{}
	for i := 0; i < len(arr); i++ {
		if !vis[i] {
			s.dfs(i, arr, vis)
		}
	}

	fmt.Print("Topologoical sort order ")
	for len(s.nums) != 0 {
		fmt.Print(s.pop(), " ")
	}
	fmt.Println()
}
func (s *_S) dfs(node int, graph [][]int, vis []bool) {
	vis[node] = true
	for i := 0; i < len(graph[node]); i++ {
		if !vis[i] && graph[node][i] == 1 {
			s.dfs(i, graph, vis)
		}
	}
	s.push(node)
}

type _S struct {
	nums []int
}

func (s *_S) push(i int) {
	s.nums = append(s.nums, i)
}
func (s *_S) pop() int {
	if len(s.nums) != 0 {
		top := s.nums[len(s.nums)-1]
		if len(s.nums) == 1 {
			s.nums = []int{}
		} else {
			s.nums = s.nums[:len(s.nums)-1]
		}
		return top
	}
	return -1
}
