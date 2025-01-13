package stacksqueues

import "fmt"

type Queue struct {
	nums []int
}

func (q *Queue) Enqueue(n int) {
	q.nums = append(q.nums, n)
}

func (q *Queue) DeQueue() int {
	length := len(q.nums)
	if length > 0 {
		removedItem := q.nums[0]
		if length != 1 {
			q.nums = q.nums[1:]
		} else {
			q.nums = []int{}
		}
		return removedItem
	} else {
		fmt.Println("Empty Queue")
	}
	return -1
}

func DoQueueOperations() {
	myQueue := &Queue{}
	myQueue.Enqueue(5)
	myQueue.Enqueue(10)
	myQueue.Enqueue(15)
	myQueue.Enqueue(20)
	fmt.Println(myQueue.nums)
	myQueue.DeQueue()
	myQueue.DeQueue()
	myQueue.DeQueue()
	fmt.Println(myQueue.nums)
	myQueue.DeQueue()
	fmt.Println(myQueue.nums)
}
