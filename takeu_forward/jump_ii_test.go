// ! https://youtu.be/7SBVnw7GSTk?si=VZ7BWY63vpMVtcu4
package takeuforward

import (
	"math"
	"testing"

	"github.com/go-playground/assert"
)

func TestJump2(t *testing.T) {
	tests := []struct {
		nums     []int
		expected int
	}{
		{nums: []int{2, 3, 1, 1, 4}, expected: 2},
		{nums: []int{2, 3, 1, 1, 4, 1, 1, 1, 2}, expected: 3},
	}
	for _, testCase := range tests {
		t.Run("", func(t *testing.T) {
			actualResult := jumpGame(testCase.nums)
			assert.Equal(t, actualResult, testCase.expected)
		})
	}
}

func jumpGame(nums []int) int {
	jumps := 0
	l, r := 0, 0
	for r < len(nums)-1 {
		farthest := 0
		for i := l; i <= r; i++ {
			farthest = int(math.Max(float64(nums[i]+i), float64(farthest)))
		}
		l = r + 1
		r = farthest
		jumps += 1
	}
	return jumps
}
