// https://leetcode.com/problems/two-sum/submissions/
package arrays

func TwoSum(nums []int, target int) []int {
	ans := []int{-1, -1}
	c := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		req := target - nums[i]
		_, ok := c[req]
		if ok {
			ans[0] = c[req]
			ans[1] = i
			break
		} else {
			c[nums[i]] = i
		}
	}
	return ans
}
