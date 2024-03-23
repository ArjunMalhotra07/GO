package graphs

import "fmt"

func PerformDijkstra() {
	// ! Initialisation of Graph Vars
	listGraph := &Graph{}
	matrixGraph := &GraphMatrix{}

	exampleGraphStruct := &ExampleGraphStruct{array: [][]int{{1, 2, 4}, {1, 3, 4}, {2, 3, 2}, {3, 4, 3}, {3, 5, 1}, {3, 6, 6}, {4, 6, 2}, {5, 6, 3}}, maxVerticesCount: 6} //! Weighted Graphs

	matrixGraph.Nodes = make([][]int, exampleGraphStruct.maxVerticesCount+1)
	for i := 1; i <= exampleGraphStruct.maxVerticesCount; i++ {
		listGraph.addVertexMethod(i)
	}
	for i := 0; i < exampleGraphStruct.maxVerticesCount+1; i++ {
		matrixGraph.Nodes[i] = make([]int, exampleGraphStruct.maxVerticesCount+1)
	}
	for i := 0; i < len(exampleGraphStruct.array); i++ {
		listGraph.addEdge(exampleGraphStruct.array[i][0], exampleGraphStruct.array[i][1])
		matrixGraph.addEdgeInMatrixGraph(exampleGraphStruct.array[i][0], exampleGraphStruct.array[i][1])
		matrixGraph.addEdgeWeightInMatrixGraph(exampleGraphStruct.array[i][0], exampleGraphStruct.array[i][1], exampleGraphStruct.array[i][2])
	}
	// ! Print Graph Info
	listGraph.printGraphAdjacencyList()
	matrixGraph.printGraphMatrix()
	// ! BFS and DFS of same Graph
	BreadthFirstSearch(listGraph)
	DepthFirstSearch(listGraph)
	fmt.Println(FindMinDistMatrix(*matrixGraph, 3))
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
	//! Distance, Node
	h.addInHeap([]int{0, source})
	for len(h.heap) != 0 {
		root := h.getTopElementFromHeap()
		for node := 0; node < len(g.Nodes); node++ {
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
