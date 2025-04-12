//! https://www.w3schools.com/dsa/dsa_algo_insertionsort.php
package basics

import (
	"reflect"
	"testing"
)

func sort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		num := arr[i]
		for j := i - 1; j >= 0; j-- {
			if arr[j] > num {
				arr[j], arr[i] = arr[i], arr[j]
				i = j
			} else {
				break
			}
		}
	}
	return arr
}

func TestInsertionSort(t *testing.T) {
	tests := []struct {
		arr      []int
		expected []int
	}{
		{arr: []int{-7, 5, 10, 1, 0, 5, -10, -37, 21}, expected: []int{-37, -10, -7, 0, 1, 5, 5, 10, 21}},
	}
	for _, test := range tests {
		t.Run("Test Insertion Sort", func(t *testing.T) {
			actualResult := sort(test.arr)
			if !reflect.DeepEqual(test.expected, actualResult) {
				t.Errorf("Expectd %v got %v", test.expected, actualResult)
			}
		})
	}
}
