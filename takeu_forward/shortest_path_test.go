package takeuforward

import (
	"math"
	"reflect"
	"testing"
)

func getShortestPath(graph [][]int, source int, dest int) []int {
	h := &heap{}
	paths := make([]int, len(graph))
	for i := 0; i < len(paths); i++ {
		paths[i] = math.MaxInt
	}
	paths[source] = 0
	paths[0] = 0
	from := make([]int, len(graph))
	from[source] = source

	h.add([]int{0, source})
	for len(h.nums) != 0 {
		top := h.extract()
		for i := 0; i < len(graph); i++ {
			if graph[top[1]][i] != 0 && graph[top[1]][i]+top[0] < paths[i] {
				paths[i] = graph[top[1]][i] + top[0]
				from[i] = top[1]
				h.add([]int{graph[top[1]][i] + top[0], i})
			}
		}
	}
	path := []int{}
	path = append([]int{dest}, path...)
	i := dest
	for from[i] != i {
		path = append([]int{from[i]}, path...)
		i = from[i]
	}
	return path
}

func TestShortestPath(t *testing.T) {
	tests := []struct {
		arr      [][]int
		expected []int
		source   int
		dest     int
	}{
		//! vertices are from 1->5
		{arr: [][]int{{0, 0, 0, 0, 0, 0}, {0, 0, 2, 0, 1, 0}, {0, 2, 0, 4, 0, 5}, {0, 0, 4, 0, 3, 1}, {0, 1, 0, 3, 0, 0}, {0, 0, 5, 1, 0, 0}}, expected: []int{1, 4, 3, 5}, source: 1, dest: 5},
		{arr: [][]int{
			{0, 0, 0, 0, 0, 0, 0},
			{0, 0, 2, 0, 4, 0, 0},
			{0, 2, 0, 1, 0, 3, 0},
			{0, 0, 1, 0, 2, 5, 0},
			{0, 4, 0, 2, 0, 0, 6},
			{0, 0, 3, 5, 0, 0, 1},
			{0, 0, 0, 0, 6, 1, 0},
		}, expected: []int{1, 2, 5, 6}, source: 1, dest: 6},

		// Test case 2: A fully connected graph where direct paths exist between all nodes
		{arr: [][]int{
			{0, 0, 0, 0, 0, 0, 0},
			{0, 0, 2, 9, 3, 7, 8},
			{0, 2, 0, 6, 4, 2, 5},
			{0, 9, 6, 0, 1, 2, 3},
			{0, 3, 4, 1, 0, 6, 7},
			{0, 7, 2, 2, 6, 0, 4},
			{0, 8, 5, 3, 7, 4, 0},
		}, expected: []int{1, 4, 3}, source: 1, dest: 3},
	}
	for _, testCase := range tests {
		t.Run("", func(t *testing.T) {
			actualResult := getShortestPath(testCase.arr, testCase.source, testCase.dest)
			if !reflect.DeepEqual(actualResult, testCase.expected) {
				t.Errorf("Expected %v but got %v", testCase.expected, actualResult)
			}
		})
	}
}
