package slidingwindow

import "fmt"

func ContantWindowQues() {
	arr := []int{-1, 2, 3, 3, 4, 5, -1}
	length := 4
	findSubArrays(arr, length)
}

func findSubArrays(arr []int, size int) {
	sum := 0
	left := 0
	right := size
	for i := left; i < right; i++ {
		sum += arr[i]
	}
	fmt.Println("Sum of Subarray [", arr[left], arr[right], "] -", sum)
	for right < len(arr)-1 {
		sum -= arr[left]
		right += 1
		left += 1
		sum += arr[right]
		fmt.Println("Sum of Subarray [", arr[left], arr[right], "] -", sum)

	}
}
