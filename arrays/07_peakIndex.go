package arrays

// https://leetcode.com/problems/peak-index-in-a-mountain-array/description/
func PeakIndexInMountainArray(arr []int) int {
	s := 0
	e := len(arr) - 1
	for s < e {
		mid := s + (e-s)/2
		if arr[mid] > arr[mid+1] {
			e = mid
		} else {
			s = mid + 1
		}
	}
	return s
}
