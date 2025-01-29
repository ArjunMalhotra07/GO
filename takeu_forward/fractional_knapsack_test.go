package takeuforward

import (
	"fmt"
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
		// Case with exact weight match
		{arr: [][]int{{100, 20}, {60, 10}, {120, 30}}, weight: 60, expected: 280}, // Takes all items
		// Case where knapsack weight is 0
		{arr: [][]int{{100, 20}, {200, 30}}, weight: 0, expected: 0}, // Cannot pick anything

		// Case where all items have the same value-to-weight ratio
		{arr: [][]int{{50, 10}, {100, 20}, {150, 30}}, weight: 50, expected: 250}, // Takes all

		// Case with only one item
		{arr: [][]int{{300, 50}}, weight: 30, expected: 180}, // Takes part of the single item

		// Case where all items fit perfectly
		{arr: [][]int{{100, 10}, {200, 20}, {300, 30}}, weight: 60, expected: 600}, // All items fit exactly

		// Edge case: Empty array
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
		return float64(arr[i][0])/float64(arr[i][1]) > float64(arr[j][0])/float64(arr[j][1])
	})
	fmt.Println(arr)
	total := 0
	for i := 0; i < len(arr); i++ {
		temp := arr[i]
		if wt <= 0 {
			return total
		}
		weightToAdd := 0

		if wt-temp[1] >= 0 {
			weightToAdd = temp[0]
			wt -= temp[1]
		} else {
			weightToAdd = int(wt*temp[0]) / temp[1]
			wt = 0
		}
		total += weightToAdd
	}
	return total
}
