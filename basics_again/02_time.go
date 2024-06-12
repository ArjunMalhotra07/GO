package basicsagain

import (
	"fmt"
	"time"
)

func AddToArray() {
	var n int = 1000000
	slice1 := make([]int, 0, n)
	slice2 := []int{}
	fmt.Println("Time To Add 1M values after preallocating capacity:", add(slice1, n))
	fmt.Println("Time To Add 1M values not preallocating capacity:", add(slice2, n))
}

func add(slice []int, n int) time.Duration {
	now := time.Now()
	for i := 0; i < n; i++ {
		slice = append(slice, 1)
	}
	return time.Since(now)
}
