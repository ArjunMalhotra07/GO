package heaps

import "fmt"

func GetKSortedElements01() {
	// nums := []int{1789, 87434, 343, 6, 7, 8, 9, 10, 15, 200, 99, 4560}
	// nums := []int{10, 4, 8, 13, 20}
	nums2 := []int{1, 5, 4, 6, 8, 9, 2}
	heapvariable := &Minheap{}
	for _, v := range nums2 {
		heapvariable.insert(v)
		// fmt.Println(heapvariable.nums)
	}
	// minConsumption := 0
	sortedArray := []int{}
	for i := 0; i < len(nums2); i++ {
		sortedArray = append(sortedArray, heapvariable.extractRoot())
		// minConsumption += heapvariable.extractRoot()
	}
	// fmt.Println(minConsumption)
	fmt.Println(sortedArray)
}

type Minheap struct {
	nums []int
}

func (h *Minheap) insert(key int) {
	h.nums = append(h.nums, key)
	h.minHeapifyUp(len(h.nums) - 1)
}

func (h *Minheap) minHeapifyUp(index int) {
	for h.nums[getParentIndex(index)] > h.nums[index] {
		h.swap(getParentIndex(index), index)
		index = getParentIndex(index)
	}
}

func (h *Minheap) swap(i1, i2 int) {
	h.nums[i1], h.nums[i2] = h.nums[i2], h.nums[i1]
}

func (h *Minheap) extractRoot() int {
	extracted := 0
	length := len(h.nums)
	if length > 0 {
		extracted = h.nums[0]
		h.nums[0] = h.nums[length-1]
		h.nums = h.nums[:length-1]
		h.minHeapifyDown(0)
	}
	return extracted
}

func (h *Minheap) minHeapifyDown(index int) {
	length := len(h.nums)
	if length > 0 {
		leftChildIndex := getLeftChildIndex(index)
		rightChildIndex := getRightChildIndex(index)
		childToCompare := 0
		for leftChildIndex <= length-1 {
			if leftChildIndex == length-1 {
				childToCompare = leftChildIndex
			} else if h.nums[leftChildIndex] < h.nums[rightChildIndex] {
				childToCompare = leftChildIndex
			} else {
				childToCompare = rightChildIndex
			}

			if h.nums[index] > h.nums[childToCompare] {
				h.swap(index, childToCompare)
				index = childToCompare
				leftChildIndex, rightChildIndex = getLeftChildIndex(index), getRightChildIndex(index)
			} else {
				return
			}
		}

	}
}
