package takeuforward

import (
	"math"
	"reflect"
	"testing"
)

func TestFloydWarshall(t *testing.T) {
	tests := []struct {
		arr      [][]int
		expected [][]int
	}{
		{
			arr: [][]int{
				{0, 2, math.MaxInt, math.MaxInt},
				{1, 0, 3, math.MaxInt},
				{math.MaxInt, math.MaxInt, 0, math.MaxInt},
				{3, 5, 4, 0},
			},
			expected: [][]int{
				{0, 2, 5, math.MaxInt},
				{1, 0, 3, math.MaxInt},
				{math.MaxInt, math.MaxInt, 0, math.MaxInt},
				{3, 5, 4, 0},
			},
		},
	}

	for _, testCase := range tests {
		t.Run("", func(t *testing.T) {
			actualResult := floydWarshall(testCase.arr)
			if !reflect.DeepEqual(actualResult, testCase.expected) {
				t.Errorf("Expected %v but got %v", testCase.expected, actualResult)
			}
		})
	}
}

func floydWarshall(cost [][]int) [][]int {
	for via := 0; via < len(cost); via++ {
		for i := 0; i < len(cost); i++ {
			for j := 0; j < len(cost); j++ {
				if cost[i][via] != math.MaxInt && cost[via][j] != math.MaxInt && cost[i][j] > cost[i][via]+cost[via][j] {
					cost[i][j] = cost[i][via] + cost[via][j]
				}
			}
		}
	}
	return cost
}
