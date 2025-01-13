package arrays

// https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/description/

func SearchRange(nums []int, target int) []int {
	ans := []int{-1, -1}
	ans[0] = search(nums, target, true)
	ans[1] = search(nums, target, false)

	return ans
}

func search(nums []int, target int, isFindingStart bool) int {
	s := 0
	e := len(nums) - 1
	ans := -1
	for s <= e {
		mid := s + (e-s)/2
		if target < nums[mid] {
			e = mid - 1
		} else if target > nums[mid] {
			s = mid + 1
		} else {
			ans = mid
			if isFindingStart {
				e = mid - 1
			} else {
				s = mid + 1
			}
		}
	}
	return ans
}
