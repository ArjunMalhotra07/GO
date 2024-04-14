package graphs

import "fmt"

func PerformDijkstra() {
	// ! Initialisation of Graph Vars
	matrixGraph := &GraphMatrix{}
	//! From, To, Weight
	// exampleGraphStruct := &ExampleGraphStruct{array: [][]int{{1, 2, 4}, {1, 3, 4}, {2, 3, 2}, {3, 4, 3}, {3, 5, 1}, {3, 6, 6}, {4, 6, 2}, {5, 6, 3}}, maxVerticesCount: 6} //! Weighted Graphs
	exampleGraphStruct := &ExampleGraphStruct{array: [][]int{{1, 2, 7}, {1, 3, 9}, {1, 6, 14}, {3, 6, 2}, {3, 2, 10}, {2, 4, 15}, {3, 4, 11}, {5, 4, 6}, {6, 5, 9}}, maxVerticesCount: 6}
	matrixGraph.Nodes = make([][]int, exampleGraphStruct.maxVerticesCount+1)

	for i := 0; i < exampleGraphStruct.maxVerticesCount+1; i++ {
		matrixGraph.Nodes[i] = make([]int, exampleGraphStruct.maxVerticesCount+1)
	}
	for i := 0; i < len(exampleGraphStruct.array); i++ {
		matrixGraph.addEdgeWeightInMatrixGraph(exampleGraphStruct.array[i][0], exampleGraphStruct.array[i][1], exampleGraphStruct.array[i][2])
	}
	// ! Print Graph Info
	matrixGraph.printGraphMatrix()
	// ! Finding Minimum Distance of each node from a source Vertex
	source := 1
	fmt.Println(FindMinDistMatrix(*matrixGraph, source))
}

func FindMinDistMatrix(g GraphMatrix, source int) []int {
	ans := make([]int, len(g.Nodes))
	for i := 1; i < len(ans); i++ {
		ans[i] = 343434343434
	}
	ans[source] = 0
	h := &Heap{}
	if len(g.Nodes) == 0 {
		return ans
	}
	//! Adding {Distance, Node} in Min Heap
	h.addInHeap([]int{0, source})
	for len(h.heap) != 0 {
		root := h.getTopElementFromHeap()
		for node := 1; node < len(g.Nodes); node++ {
			//! Looping through each edge and getting its weight
			cellValue := g.Nodes[root[1]][node]
			dist := cellValue + root[0]

			if cellValue != 0 && dist < ans[node] {
				ans[node] = dist
				h.addInHeap([]int{dist, node})
			}
		}
	}
	return ans
}

type Heap struct {
	heap [][]int
}

func (h *Heap) addInHeap(arr []int) {
	h.heap = append(h.heap, arr)
	h.heapifyUp(len(h.heap) - 1)
}
func (h *Heap) heapifyUp(index int) {
	for h.heap[getParentIndex(index)][0] > h.heap[index][0] {
		h.swapTwoValues(getParentIndex(index), index)
		index = getParentIndex(index)
	}
}
func (h *Heap) getTopElementFromHeap() []int {
	extracted := []int{}
	if len(h.heap) > 0 {
		extracted = h.heap[0]
		lastIndex := len(h.heap) - 1
		h.heap[0] = h.heap[lastIndex]
		h.heap = h.heap[:lastIndex]
		h.heapifyDown(0)
	}
	return extracted
}

func (h *Heap) heapifyDown(index int) {
	if len(h.heap) > 0 {
		leftChild := getLeftChildIndex(index)
		rightChild := getRightChildIndex(index)
		lastIndex := len(h.heap) - 1
		childToCompare := -1
		for leftChild <= lastIndex {
			//! Until atleast one child remains
			if leftChild == lastIndex {
				childToCompare = leftChild
			} else if h.heap[leftChild][0] < h.heap[rightChild][0] {
				childToCompare = leftChild
			} else {
				childToCompare = rightChild
			}

			//! Swap
			if h.heap[index][0] > h.heap[childToCompare][0] {
				h.swapTwoValues(index, childToCompare)
				index = childToCompare
				leftChild = getLeftChildIndex(index)
				rightChild = getRightChildIndex(index)
			} else {
				return
			}
		}

	}
}
func (h *Heap) swapTwoValues(i1, i2 int) {
	h.heap[i1], h.heap[i2] = h.heap[i2], h.heap[i1]
}
func getParentIndex(i int) int {
	return (i - 1) / 2
}
func getLeftChildIndex(i int) int {
	return (2 * i) + 1
}
func getRightChildIndex(i int) int {
	return (2 * i) + 2
}
