package takeuforward

import (
	"reflect"
	"testing"
)

func TestHeapFunc(t *testing.T) {
	tests := []struct {
		arr      [][]int
		expected [][]int
	}{
		{arr: [][]int{{10, 1}, {5, 2}, {3, 7}, {100, 5}}, expected: [][]int{{3, 7}, {5, 2}, {10, 1}, {100, 5}}},
	}

	for _, testCase := range tests {
		t.Run("", func(t *testing.T) {
			actualResult := heapFunc(testCase.arr)
			if !reflect.DeepEqual(actualResult, testCase.expected) {
				t.Errorf("Expected %v but got %v", testCase.expected, actualResult)
			}
		})
	}
}
