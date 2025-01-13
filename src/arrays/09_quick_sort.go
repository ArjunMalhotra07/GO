package arrays

import "fmt"

func DoQuickSort() {
	fmt.Println("Quick Sort func")

	arr := []int{35, 50, 15, 25, 80, 20, 90, 45, -1, 2, -43, 78, 3}
	quickSorrt(arr, 0, len(arr)-1)
	fmt.Println(arr)

}
func quickSorrt(nums []int, low, high int) {
	if low < high {
		pivot := helper(nums, low, high)
		quickSorrt(nums, low, pivot-1)
		quickSorrt(nums, pivot+1, high)
	}
}
func helper(nums []int, low, high int) int {
	pivot := nums[low]
	i := low
	j := high
	for i < j {
		for i <= high && nums[i] <= pivot {
			i += 1
		}
		for j >= low && nums[j] > pivot {
			j -= 1
		}
		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	nums[j], nums[low] = nums[low], nums[j]
	return j
}
