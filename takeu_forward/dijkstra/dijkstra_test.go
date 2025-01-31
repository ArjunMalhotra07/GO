// ! https://youtu.be/V6H1qAeB-l4?si=PVaz5MH3NkaSOvyo
package takeuforward

import (
	"math"
	"reflect"
	"testing"
)

func TestDijskstra(t *testing.T) {
	tests := []struct {
		arr      [][]int
		expected []int
		source   int
	}{
		{arr: [][]int{{0, 4, 4, 0, 0, 0}, {4, 0, 2, 0, 0, 0}, {4, 2, 0, 6, 1, 3}, {0, 0, 6, 0, 3, 2}, {0, 0, 1, 3, 0, 0}, {0, 0, 3, 2, 0, 0}}, expected: []int{0, 4, 4, 8, 5, 7}, source: 0},
		{
			arr: [][]int{
				{0, 4, 0, 0, 0, 0},
				{4, 0, 8, 0, 0, 0},
				{0, 8, 0, 7, 0, 4},
				{0, 0, 7, 0, 9, 14},
				{0, 0, 0, 9, 0, 10},
				{0, 0, 4, 14, 10, 0},
			},
			expected: []int{0, 4, 12, 19, 26, 16}, source: 0,
		},
		{
			arr: [][]int{{0}}, expected: []int{0}, source: 0,
		},
		{
			arr: [][]int{
				{0, 1, 1, 0, 0},
				{1, 0, 1, 1, 0},
				{1, 1, 0, 0, 1},
				{0, 1, 0, 0, 1},
				{0, 0, 1, 1, 0},
			},
			expected: []int{0, 1, 1, 2, 2}, source: 0,
		},
		{
			arr: [][]int{
				{0, 2, 4, 0},
				{2, 0, 1, 5},
				{4, 1, 0, 1},
				{0, 5, 1, 0},
			},
			expected: []int{0, 2, 3, 4}, source: 0,
		},
		{
			arr: [][]int{
				{0, 1, 2, 3, 4},
				{1, 0, 1, 2, 3},
				{2, 1, 0, 1, 2},
				{3, 2, 1, 0, 1},
				{4, 3, 2, 1, 0},
			},
			expected: []int{0, 1, 2, 3, 4}, source: 0,
		},
	}
	for _, testCase := range tests {
		t.Run("", func(t *testing.T) {
			actualResult := dijskstra(testCase.arr, testCase.source)
			if !reflect.DeepEqual(actualResult, testCase.expected) {
				t.Errorf("Expected %v but got %v", testCase.expected, actualResult)
			}
		})
	}
}

func dijskstra(arr [][]int, source int) []int {
	leastDistance := make([]int, len(arr))
	for i := 0; i < len(leastDistance); i++ {
		leastDistance[i] = math.MaxInt
	}
	h := &heap{}
	h.add([]int{source, 0})
	leastDistance[source] = 0
	for len(h.nums) != 0 {
		top := h.extract()
		for i := 0; i < len(arr); i++ {
			if top[1] != i && arr[top[1]][i] != 0 && arr[top[1]][i]+top[0] < leastDistance[i] {
				leastDistance[i] = arr[top[1]][i] + top[0]
				h.add([]int{leastDistance[i], i})
			}
		}
	}
	return leastDistance
}
