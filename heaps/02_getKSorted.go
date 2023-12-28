package heaps

import (
	"fmt"
	"math"
)

func DoReturnKClosestPoints() {
	// array := [][]int{{3, 3}, {-2, 4}, {5, -1}}
	// array := [][]int{{1, 0}, {10, 0}, {11, 0}, {5, 0}}
	array := [][]int{{1, 3}, {-2, 2}}

	fmt.Println(kClosest(array, 2))

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
	h.array = append(h.array, hMapStruct{nums: temp, sum: getDistance(temp)})
	h.minHeapifyUp(len(h.array) - 1)
}
func (h *hMapHeap) minHeapifyUp(index int) {

	for h.array[getparentIndex(index)].sum > h.array[index].sum {
		h.swap(getparentIndex(index), index)
		index = getparentIndex(index)
	}
}

func getDistance(tempArray []int) float64 {
	distX := math.Pow(float64(tempArray[0])-float64(0), 2)
	distY := math.Pow(float64(tempArray[1])-float64(0), 2)
	total := math.Sqrt(distX + distY)
	return total
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

func getparentIndex(index int) int {
	return (index - 1) / 2
}
func getLeftChildIndex(index int) int {
	return 2*index + 1
}
func getRightChildIndex(index int) int {
	return 2*index + 2
}
