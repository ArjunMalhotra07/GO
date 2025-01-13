package arrays

// https://leetcode.com/problems/find-numbers-with-even-number-of-digits/

func FindNumbers(nums []int) int {
	count := 0
	for i := 0; i < len(nums); i++ {
		if checkForEven(nums[i]) {
			count += 1
		}
	}
	return count
}
func checkForEven(i int) bool {
	count := 0
	for i > 0 {
		count++
		i /= 10
	}

	return count%2 == 0
}
