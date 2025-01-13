package arrays

import "fmt"

func LinearSearch(array []int, target int) {
	for i := 0; i < len(array); i++ {
		if array[i] == target {
			fmt.Println(i)
			return
		}
	}
	fmt.Println("-1")
}

func LinearSearchOnlyInRange(input []int, target, start, end int) {
	for i := start; i < end; i++ {
		if input[i] == target {
			fmt.Println(i)
			return
		}
	}
	fmt.Println("-1")
}
