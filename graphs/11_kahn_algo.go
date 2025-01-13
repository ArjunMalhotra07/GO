package graphs

import "fmt"

func PerformKahnSort() {
	//! Calculate indegree
	// arr := [][]int{{0, 0, 0, 0, 0, 0}, {0, 0, 0, 0, 0, 0}, {0, 0, 0, 1, 0, 0}, {0, 1, 0, 0, 0, 0}, {1, 1, 0, 0, 0, 0}, {1, 0, 1, 0, 0, 0}} //! Non cyclic example Kahn works
	arr := [][]int{{0, 0, 0, 0, 0, 1}, {0, 0, 0, 1, 0, 0}, {0, 0, 0, 1, 0, 0}, {1, 0, 0, 0, 0, 0}, {1, 1, 0, 0, 0, 0}, {0, 0, 1, 0, 0, 0}} //! Cyclic example
	indegree := make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[0]); j++ {
			if arr[i][j] == 1 {
				indegree[j] += 1
			}
		}
	}
	//! Enq nodes with indegree = 0
	q := &_Q{}
	for i := 1; i < len(indegree); i++ {
		if indegree[i] == 0 {
			q.enQ(i)
		}
	}
	sorted := []int{}
	//! Pop values to get sorted order
	for len(q.nums) != 0 {
		popped := q.deQ()
		sorted = append(sorted, popped)
		for i := 0; i < len(arr[popped]); i++ {
			if arr[popped][i] == 1 {
				indegree[i] -= 1
				if indegree[i] == 0 {
					q.enQ(i)
				}
			}
		}
	}
	fmt.Print("Kahn Algorithm ", sorted)
	if len(sorted) != len(arr) {
		fmt.Println("Cyclic", len(sorted))
	} else {
		fmt.Println("Non cyclic")
	}
}

type _Q struct {
	nums []int
}

func (q *_Q) enQ(value int) {
	q.nums = append(q.nums, value)
}

func (q *_Q) deQ() int {
	if len(q.nums) > 0 {
		removedElement := q.nums[0]
		if len(q.nums) == 1 {
			q.nums = []int{}
		} else {
			q.nums = q.nums[1:]
		}
		return removedElement
	}
	return -1
}
