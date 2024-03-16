package arrays

func FindInInfiniteArray(nums []int, target int) int {
	start := 0
	end := 1
	for target > nums[end] {
		temp := end + 1
		end = end + (end-start+1)*2
		start = temp
	}
	return DoBinarySearch(nums, target)
}

func DoBinarySearch(nums []int, target int) int {
	s := 0
	e := len(nums) - 1
	for s <= e {
		mid := s + (e-s)/2
		if target < nums[mid] {
			e = mid - 1
		} else if target > nums[mid] {
			s = mid + 1
		} else {
			return mid
		}
	}
	return -1
}
