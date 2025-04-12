package questions

import (
	"reflect"
	"testing"
)

func movePositives(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		if arr[i] < 0 {
			continue
		}
		for j := i - 1; j >= 0; j-- {
			if arr[j] > 0 {
				break
			}
			arr[j], arr[i] = arr[i], arr[j]
			i = j
		}
	}
	return arr
}
func TestMoveElements(t *testing.T) {
	tests := []struct {
		arr      []int
		expected []int
	}{
		{arr: []int{-10, 7, 2, -5, 3, -27, 4}, expected: []int{7, 2, 3, 4, -10, -5, -27}},
	}
	for _, test := range tests {
		t.Run("Move Positive Elements to front", func(t *testing.T) {
			actualResult := movePositives(test.arr)
			if !reflect.DeepEqual(actualResult, test.expected) {
				t.Errorf("Expected %v, got %v ", test.expected, actualResult)
			}
		})
	}
}
