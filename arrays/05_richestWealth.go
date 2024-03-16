// https://leetcode.com/problems/richest-customer-wealth/
package arrays

func MaximumWealth(accounts [][]int) int {
	maxSum := 0
	for i := 0; i < len(accounts); i++ {
		sum := 0
		for j := 0; j < len(accounts[i]); j++ {
			sum += accounts[i][j]
		}
		if sum > maxSum {
			maxSum = sum
		}
		sum = 0
	}
	return maxSum
}
