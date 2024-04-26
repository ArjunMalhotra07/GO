package slidingwindow

import (
	"fmt"
	"math"
)

func FindLongestSubArray() {
	arr := []int{2, 5, 1, 10, 10}
	sum := 14
	find(arr, sum)
}
func find(arr []int, sum int) {
	l, r, currentSum := 0, 0, 0
	maxLength := 0
	for r < len(arr) {
		currentSum += arr[r]
		for currentSum > sum {
			currentSum -= arr[l]
			l += 1
		}
		if currentSum <= sum {
			maxLength = int(math.Max(float64(r-l+1), float64(maxLength)))
			r += 1
		}
	}
	fmt.Println(maxLength)
}
