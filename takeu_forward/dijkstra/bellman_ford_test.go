// ! https://youtu.be/0vVofAhAYjc?si=gcoiUEiecUjssJMz
package takeuforward

import (
	"fmt"
	"math"
	"reflect"
	"testing"
)

func TestBellmanFord(t *testing.T) {
	tests := []struct {
		v        int
		arr      [][]int
		source   int
		expected []int
	}{
		{v: 2, arr: [][]int{{0, 1, 9}}, source: 0, expected: []int{0, 9}},
		{v: 5, arr: [][]int{{1, 3, 2}, {4, 3, -1}, {2, 4, 1}, {1, 2, 1}, {0, 1, 5}}, source: 0, expected: []int{0, 5, 6, 6, 7}},
		{v: 4, arr: [][]int{{0, 1, 4}, {1, 2, -6}, {2, 3, 5}, {3, 1, -2}}, source: 0, expected: []int{-1}},
	}
	for _, testCase := range tests {
		t.Run("", func(t *testing.T) {
			actualResult := bellmanFord(testCase.v, testCase.arr, testCase.source)
			if !reflect.DeepEqual(actualResult, testCase.expected) {
				t.Errorf("Expected %v but got %v", testCase.expected, actualResult)
			}
		})
	}
}

// ! [u,v,wt] in arr[][]
func bellmanFord(v int, arr [][]int, source int) []int {
	dist := make([]int, v)
	for i := 0; i < len(dist); i++ {
		dist[i] = math.MaxInt
	}
	dist[source] = 0
	for i := 0; i < v-1; i++ {
		for _, edge := range arr {
			u := edge[0]
			v := edge[1]
			wt := edge[2]
			if dist[u] != math.MaxInt && dist[u]+wt < dist[v] {
				dist[v] = dist[u] + wt
			}
		}
	}
	fmt.Println(dist)
	for i := 0; i < len(arr); i++ {
		edge := arr[i]
		u := edge[0]
		v := edge[1]
		wt := edge[2]
		if dist[u]+wt < dist[v] {
			return []int{-1}
		}
	}
	return dist
}
