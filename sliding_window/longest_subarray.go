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
	left, right := 0, 0
	currentSum := 0
	maxLength := 0
	for right < len(arr) {
		currentSum += arr[right]
		for currentSum > sum {
			currentSum -= arr[left]
			left += 1
		}
		if currentSum <= sum {
			fmt.Println(arr[left : right+1])
			maxLength = int(math.Max(float64(maxLength), float64(right-left+1)))
			right += 1
		}
	}

	fmt.Println(maxLength)

}
