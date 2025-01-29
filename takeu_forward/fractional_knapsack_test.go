package takeuforward

import (
	"sort"
	"testing"

	"github.com/go-playground/assert"
)

func TestFractionalKnapsack(t *testing.T) {
	tests := []struct {
		arr      [][]int
		weight   int
		expected int
	}{
		{arr: [][]int{{100, 20}, {60, 10}, {100, 50}, {200, 50}}, weight: 90, expected: 380},
		{arr: [][]int{{100, 20}, {60, 10}, {120, 30}}, weight: 60, expected: 280},
		{arr: [][]int{{100, 20}, {200, 30}}, weight: 0, expected: 0},
		{arr: [][]int{{50, 10}, {100, 20}, {150, 30}}, weight: 50, expected: 250},
		{arr: [][]int{{300, 50}}, weight: 30, expected: 180},
		{arr: [][]int{{100, 10}, {200, 20}, {300, 30}}, weight: 60, expected: 600},
		{arr: [][]int{}, weight: 50, expected: 0},
	}
	for _, testCase := range tests {
		t.Run("", func(t *testing.T) {
			actualResult := fractionalKnapsack(testCase.arr, testCase.weight)
			assert.Equal(t, actualResult, testCase.expected)
		})
	}
}

func fractionalKnapsack(arr [][]int, wt int) int {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i][0]*arr[j][1] > arr[j][0]*arr[i][1]
	})
	total := 0
	for i := 0; i < len(arr); i++ {
		temp := arr[i]
		if wt <= 0 {
			break
		}
		if wt >= temp[1] {
			total += temp[0]
			wt -= temp[1]
		} else {
			total += wt * temp[0] / temp[1]
			wt = 0
		}
	}
	return total
}
