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
	left := 0
	right := 0
	maxLength := 0
	for r < len(arr) {
		currentSum += arr[r]
		for currentSum > sum {
			currentSum -= arr[l]
			l += 1
		}
		if currentSum <= sum {
			if r-l+1 > maxLength {
				left = l
				right = r
				maxLength = int(math.Max(float64(r-l+1), float64(maxLength)))
			}
			r += 1
		}
	}
	fmt.Println(maxLength)
	fmt.Println("left Index", left)
	fmt.Println("Index right", right)
}
