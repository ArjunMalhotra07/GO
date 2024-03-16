package arrays

import "fmt"

func DoSumOfSubsets() {
	// Example array with numbers
	nums := []int{10, 20, 30, 40, 50}
	// Target sum
	target := 50

	// Variable to store the result
	var result [][]int

	// Find subsets with the given target sum
	subsetSum(nums, target, []int{}, 0, target, &result)

	// Print the result
	fmt.Println("Subsets with sum", target, "are:", result)
}

func subsetSum(nums []int, target int, currentSubset []int, currentIndex int, remainingSum int, result *[][]int) {
	// Check if the remaining sum is zero, indicating that the currentSubset sums up to the target
	if remainingSum == 0 {
		// Append a copy of the currentSubset to the result
		*result = append(*result, append([]int{}, currentSubset...))
	}

	// Explore two options: include the current element or skip it
	for i := currentIndex; i < len(nums); i++ {
		// Skip the current element if it exceeds the remaining sum
		if remainingSum-nums[i] < 0 {
			continue
		}

		// Include the current element in the subset
		currentSubset = append(currentSubset, nums[i])

		// Recursively explore subsets with the current element included
		subsetSum(nums, target, currentSubset, i+1, remainingSum-nums[i], result)

		// Exclude the current element to explore subsets without it
		currentSubset = currentSubset[:len(currentSubset)-1]
	}
}
