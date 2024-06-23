package graphs

import "fmt"

func PerformPrims() {
	exampleGraphStruct := &ExampleGraphStruct{maxVerticesCount: 7, array: [][]int{{0, 1, 5}, {0, 2, 3}, {1, 2, 4}, {4, 6, 4}, {4, 5, 3}, {5, 6, 4}, {1, 4, 2}, {2, 5, 6}, {1, 3, 6}, {2, 3, 5}, {4, 3, 6}, {5, 3, 6}}}
	matrixGraph := &GraphMatrix{}
	matrixGraph.Nodes = make([][]int, exampleGraphStruct.maxVerticesCount)

	for i := 0; i < exampleGraphStruct.maxVerticesCount; i++ {
		matrixGraph.Nodes[i] = make([]int, exampleGraphStruct.maxVerticesCount)
	}
	for i := 0; i < len(exampleGraphStruct.array); i++ {
		matrixGraph.addEdgeWeightInMatrixGraph(exampleGraphStruct.array[i][0], exampleGraphStruct.array[i][1], exampleGraphStruct.array[i][2])
	}
	// ! Print Graph Info
	matrixGraph.printGraphMatrix()
	fmt.Println(PrimsAlgo(matrixGraph.Nodes))
}
func PrimsAlgo(graph [][]int) int {
	visited := make([]bool, len(graph))
	ans := [][]int{}
	h := &heap{}
	h.addToHeap([]int{0, 0, -1})
	sum := 0
	for len(h.edges) != 0 {
		root := h.extractRoot()
		wt := root[0]
		node := root[1]
		if !visited[node] {
			sum += wt
			ans = append(ans, root)
			for i := 0; i < len(graph[0]); i++ {
				if graph[node][i] != 0 && !visited[i] {
					h.addToHeap([]int{graph[node][i], i, node})
				}
			}
			visited[node] = true
		}
	}
	fmt.Println(ans[1:])
	return sum
}

type heap struct {
	edges [][]int
}

func (h *heap) addToHeap(arr []int) {
	h.edges = append(h.edges, arr)
	h.heapifyUp(len(h.edges) - 1)
}
func (h *heap) heapifyDown(i int) {
	if len(h.edges) > 0 {
		leftChild := getLeftChildIndex(i)
		rightChild := getRightChildIndex(i)
		lastIndex := len(h.edges) - 1
		childToCompare := -1
		for leftChild <= lastIndex {
			//! Until atleast one child remains
			if leftChild == lastIndex {
				childToCompare = leftChild
			} else if h.edges[leftChild][0] < h.edges[rightChild][0] {
				childToCompare = leftChild
			} else {
				childToCompare = rightChild
			}
			//! Swap
			if h.edges[i][0] > h.edges[childToCompare][0] {
				h.swapTwoValues(i, childToCompare)
				i = childToCompare
				leftChild = getLeftChildIndex(i)
				rightChild = getRightChildIndex(i)
			} else {
				return
			}
		}

	}
}
func (h *heap) extractRoot() []int {
	if len(h.edges) != 0 {
		popped := h.edges[0]
		if len(h.edges) != 1 {
			h.edges[0] = h.edges[len(h.edges)-1]
			h.edges = h.edges[:len(h.edges)-1]
			h.heapifyDown(0)
		} else {
			h.edges = [][]int{}
		}
		return popped
	}
	return []int{}
}
func (h *heap) heapifyUp(i int) {
	for h.edges[getParentIndex(i)][0] > h.edges[i][0] {
		h.swapTwoValues(i, getParentIndex(i))
		i = getParentIndex(i)
	}
}
func (h *heap) swapTwoValues(i1, i2 int) {
	h.edges[i1], h.edges[i2] = h.edges[i2], h.edges[i1]
}
