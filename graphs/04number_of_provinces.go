package graphs

func GetNumberOfProvinces(isConnected [][]int) int {
	ans := 0
	length := len(isConnected)
	if length == 0 {
		return ans
	}
	visitedArray := make([]int, length+1)
	myQueue := &ProvinceQueue{}
	for i := 1; i < len(visitedArray)-1; i++ {
		if visitedArray[i] == 0 {
			ans += 1
			visitedArray[i] = 1
			myQueue.Enqueue(i)
			for len(myQueue.nums) != 0 {
				dequeued := myQueue.Dequeue()
				for j := 1; j < len(isConnected[dequeued]); j++ {
					if isConnected[dequeued][j] == 1 && visitedArray[j] == 0 {
						visitedArray[j] = 1
						myQueue.Enqueue(j)
					}
				}
			}
		}
	}
	return ans
}

type ProvinceQueue struct {
	nums []int
}

func (q *ProvinceQueue) Enqueue(value int) {
	q.nums = append(q.nums, value)
}

func (q *ProvinceQueue) Dequeue() int {
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
