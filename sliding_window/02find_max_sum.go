package slidingwindow

import "fmt"
//! Max Sum of a fixed size window (k) in an array
func GetMaxSum() {
	fmt.Println(GetMax([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 3))
}
func GetMax(a []int, k int) int {
	sum := 0
	for i := 0; i < k; i++ {
		sum += a[i]
	}
	max := sum

	for i := k; i < len(a); i++ {
		sum += a[i] - a[i-k]
		if sum > max {
			max = sum
		}
	}
	return max
}
