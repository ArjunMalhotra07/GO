package arrays

import "fmt"

func getMaximumPartitions(performance []int) int {
	n := len(performance)
	partitions := 0
	currentAnd := performance[0]
	for i := 1; i < n; i++ {
		currentAnd &= performance[i]
		if currentAnd == 0 {
			partitions++
			if i+1 < n {
				currentAnd = performance[i+1]
			}
		}
	}
	return partitions
}

func DogetMaximumPartitions() {
	performance1 := []int{3, 1, 2, 3}
	performance2 := []int{5, 2, 6, 1, 1, 4}
	performance3 := []int{2, 3, 1, 5, 2}
	performance4 := []int{1, 2, 3}
	fmt.Println(getMaximumPartitions(performance1))
	fmt.Println(getMaximumPartitions(performance2))
	fmt.Println(getMaximumPartitions(performance3))
	fmt.Println(getMaximumPartitions(performance4))
}
