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
	buildHeap := []int{10, 20, 30, 50, 0}
	for _, v := range buildHeap {
		m.Insert(v)
		fmt.Println(m)
	}
}

type MaxHeap struct {
	array []int
}

func (h *MaxHeap) Insert(key int) {
	h.array = append(h.array, key)
	h.maxHeapify(len(h.array) - 1)
}
func (h *MaxHeap) maxHeapify(i int) {
	for h.array[parentIndex(i)] < h.array[i] {
		h.swapTwoValues(parentIndex(i), i)
		i = parentIndex(i)
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
