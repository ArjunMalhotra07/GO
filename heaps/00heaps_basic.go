package heaps

import "fmt"

func makeHeapTree() {
	/*
			parent i
		left Child 2i + 1
		right Child 2i + 2
	*/
	m :=
		&MaxHeap{}
	buildHeap := []int{99, -10, 8, 45, 90, 100, -100, 56}
	for _, v := range buildHeap {
		m.Insert(v)
		fmt.Println(m)
	}

	ansArray := []int{}
	for i := 0; i < 3; i++ {
		ansArray = append(ansArray, m.extractRoot())
	}
	fmt.Println(ansArray)
}

type MaxHeap struct {
	array []int
}

func (h *MaxHeap) Insert(key int) {
	h.array = append(h.array, key)
	h.maxHeapifyUp(len(h.array) - 1)
}
func (h *MaxHeap) maxHeapifyUp(i int) {
	for h.array[parentIndex(i)] < h.array[i] {
		h.swapTwoValues(parentIndex(i), i)
		i = parentIndex(i)
	}
}
func (h *MaxHeap) extractRoot() int {
	extracted := -1
	if len(h.array) > 0 {
		extracted = h.array[0]
		lastIndex := len(h.array) - 1
		h.array[0] = h.array[lastIndex]
		h.array = h.array[:lastIndex]
		h.maxHeapifyDown(0)
	}
	return extracted
}
func (h *MaxHeap) maxHeapifyDown(index int) {
	if len(h.array) > 0 {
		leftChild := leftChildIndex(index)
		rightChild := rightChildIndex(index)
		lastIndex := len(h.array) - 1
		childToCompare := -1
		for leftChild <= lastIndex {
			//! Until atleast one child remains
			if leftChild == lastIndex {
				childToCompare = leftChild
			} else if h.array[leftChild] > h.array[rightChild] {
				childToCompare = leftChild
			} else {
				childToCompare = rightChild
			}

			//! Swap
			if h.array[index] < h.array[childToCompare] {
				h.swapTwoValues(index, childToCompare)
				index = childToCompare
				leftChild = leftChildIndex(index)
				rightChild = rightChildIndex(index)
			} else {
				return
			}
		}

	}
}
func parentIndex(i int) int {
	return (i - 1) / 2
}

func (h *MaxHeap) swapTwoValues(i1, i2 int) {
	h.array[i1], h.array[i2] = h.array[i2], h.array[i1]
}

func leftChildIndex(i int) int {
	return 2*i + 1
}
func rightChildIndex(i int) int {
	return 2*i + 2
}
