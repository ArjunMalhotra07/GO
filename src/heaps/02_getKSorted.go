package heaps

import (
	"fmt"
	"math"
)

func DoReturnKClosestPoints() {
	// array := [][]int{{3, 3}, {-2, 4}, {5, -1}}
	// array := [][]int{{1, 0}, {10, 0}, {11, 0}, {5, 0}}
	array := [][]int{{1, 3}, {-2, 2}}

	fmt.Println(kClosest(array, 1))

}

type hMapHeap struct {
	array []hMapStruct
}
type hMapStruct struct {
	nums []int
	sum  float64
}

func kClosest(points [][]int, k int) [][]int {
	mapVariable := &hMapHeap{}
	for _, v := range points {
		mapVariable.insert(v)
	}
	ansArray := [][]int{}
	for i := 0; i < k; i++ {
		ansArray = append(ansArray, mapVariable.extractRootElement().nums)
	}
	return ansArray
}

func (h *hMapHeap) insert(temp []int) {
	h.array = append(h.array, hMapStruct{nums: temp, sum: math.Sqrt(math.Pow(float64(temp[0])-float64(0), 2) + math.Pow(float64(temp[1])-float64(0), 2))})
	h.minHeapifyUp(len(h.array) - 1)
}
func (h *hMapHeap) minHeapifyUp(index int) {
	for h.array[getParentIndex(index)].sum > h.array[index].sum {
		h.swap(getParentIndex(index), index)
		index = getParentIndex(index)
	}
}

func (h *hMapHeap) swap(i1, i2 int) {
	h.array[i1], h.array[i2] = h.array[i2], h.array[i1]
}

func (h *hMapHeap) extractRootElement() hMapStruct {
	extractedElement := hMapStruct{}
	if len(h.array) > 0 {
		extractedElement = h.array[0]
		h.array[0] = h.array[len(h.array)-1]
		h.array = h.array[:len(h.array)-1]
		h.minHeapifyDown(0)
	}
	return extractedElement
}

func (h *hMapHeap) minHeapifyDown(index int) {
	if len(h.array) > 0 {

		leftChildIndex := getLeftChildIndex(index)
		rigtChildIndex := getRightChildIndex(index)
		lastIndex := len(h.array) - 1
		childToCompare := 0
		for leftChildIndex <= lastIndex {

			if leftChildIndex == lastIndex {
				childToCompare = leftChildIndex
			} else if h.array[leftChildIndex].sum < h.array[rigtChildIndex].sum {
				childToCompare = leftChildIndex
			} else {
				childToCompare = rigtChildIndex
			}
			if h.array[index].sum > h.array[childToCompare].sum {
				h.swap(index, childToCompare)
				index = childToCompare
				leftChildIndex = getLeftChildIndex(index)
				rigtChildIndex = getRightChildIndex(index)
			} else {
				return
			}
		}

	}
}
