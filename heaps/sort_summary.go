package heaps

import (
	"fmt"
)

func DoSortSummary() {
	arr := []int{5, 6, 5, 6, 3, 4, 5, 7, 8, 9, 10, 3, 4, 2, 1, 4, 6, 8, 9, 10, 2, 3, 4, 5, 7, 2, 5, 4, 3, 2, 3, 4, 5, 6, 10, 8, 4, 3, 5, 1}
	fmt.Println(SortSummary((arr)))
}
func SortSummary(arr []int) [][]int {
	// Step 1: Count the frequency of each element using a map
	frequency := make(map[int]int)
	for _, num := range arr {
		frequency[num]++
	}
	fmt.Println(frequency)

	// Step 2: Create the 2D array with elements and their frequencies
	var summary [][]int
	for num, freq := range frequency {
		summary = append(summary, []int{num, freq})
	}
	myHeap := &Heap{}
	for i := 0; i < len(summary); i++ {
		myHeap.insert(summary[i])
	}
	for i := 0; i < len(summary); i++ {
		summary[i] = myHeap.extractRoot()
	}

	return summary
}

type Heap struct {
	heap [][]int
}

func (h *Heap) insert(arr []int) {
	h.heap = append(h.heap, arr)
	h.maxHeapifyUp(len(h.heap) - 1)
}
func (h *Heap) maxHeapifyUp(i int) {
	for h.heap[getParentIndex(i)][1] < h.heap[i][1] {
		h.swap(getParentIndex(i), i)
		i = getParentIndex(i)
	}
}
func (h *Heap) maxHeapifyDown(i int) {
	last := len(h.heap) - 1
	for {
		l := getLeftChildIndex(i)
		r := getRightChildIndex(i)
		c := i // Assume current index is the largest

		// Determine which child to compare with
		if l <= last && (h.heap[l][1] > h.heap[c][1] || (h.heap[l][1] == h.heap[c][1] && h.heap[l][0] < h.heap[c][0])) {
			c = l
		}
		if r <= last && (h.heap[r][1] > h.heap[c][1] || (h.heap[r][1] == h.heap[c][1] && h.heap[r][0] < h.heap[c][0])) {
			c = r
		}

		// If the current index is not the largest, swap with the largest child
		if c != i {
			h.swap(i, c)
			i = c // Move to the child index to continue heapifying down
		} else {
			break // If no swap was made, we are done
		}
	}
}

func (h *Heap) extractRoot() []int {
	top := []int{}
	if len(h.heap) != 0 {
		top = h.heap[0]
		last := len(h.heap) - 1
		h.heap[0] = h.heap[last]
		h.heap = h.heap[:last]
		h.maxHeapifyDown(0)
	}
	return top
}
func (h *Heap) swap(i1, i2 int) {
	h.heap[i1], h.heap[i2] = h.heap[i2], h.heap[i1]
}
func getParentIndex(i int) int {
	return (i - 1) / 2
}
func getLeftChildIndex(i int) int {
	return i*2 + 1
}
func getRightChildIndex(i int) int {
	return i*2 + 2
}
