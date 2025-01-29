// ! https://youtu.be/tZAa_jJ3SwQ?si=FynkTPl_O0zTLJUP
package takeuforward

import (
	"testing"

	"github.com/go-playground/assert"
)

func TestJumpGame(t *testing.T) {
	tests := []struct {
		name     string
		arr      []int
		expected bool
	}{
		{name: "ShouldFail", arr: []int{1, 2, 3, 1, 1, 0, 2, 5}, expected: false},
		{name: "ShouldFail", arr: []int{3, 2, 1, 0, 4}, expected: false},
		{name: "ShouldPass", arr: []int{2, 3, 1, 0, 4}, expected: true},
		{name: "ShouldPass ", arr: []int{1, 2, 4, 1, 1, 0, 2, 5}, expected: true}}
	for _, testCase := range tests {
		t.Run("", func(t *testing.T) {
			actualResult := JumpGame(testCase.arr)
			assert.Equal(t, actualResult, testCase.expected)
		})
	}
}

func JumpGame(steps []int) bool {
	max := 0
	for i := 0; i < len(steps); i++ {
		if i > max {
			return false
		}
		if steps[i]+i >= len(steps) {
			return true
		}
		if steps[i]+i > max {
			max = steps[i] + i
		}
	}
	return true
}
