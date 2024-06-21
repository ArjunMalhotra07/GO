package graphs

import "fmt"

func PerformBFS() {
	exampleGraphStruct := &ExampleGraphStruct{array: [][]int{{1, 2}, {1, 6}, {2, 3}, {2, 4}, {6, 7}, {6, 8}, {4, 5}, {7, 5}}, maxVerticesCount: 8}
	listGraph := &Graph{}
	listGraph.makeListGraph(*exampleGraphStruct)
	listGraph.printGraphAdjacencyList()
	BreadthFirstSearch(listGraph)
}
func BreadthFirstSearch(graph *Graph) {
	if len(graph.Vertices) == 0 {
		fmt.Println("Empty graph")
		return
	}

	visitedArray := make([]int, len(graph.Vertices)+1)
	myQueue := &Queue{}
	myQueue.Enqueue(graph.Vertices[0])
	ansArray := []int{}
	for len(myQueue.vertices) != 0 {
		node := myQueue.DeQueue()
		if visitedArray[node.Key] != 1 {
			visitedArray[node.Key] = 1
			ansArray = append(ansArray, node.Key)

			for _, v := range node.Neighbours {
				if visitedArray[v.Key] != 1 {
					myQueue.Enqueue(v)
				}
			}
		}
	}
	fmt.Println("BFS:", ansArray)
}

type Queue struct {
	vertices []*Vertex
}

func (q *Queue) Enqueue(vertex *Vertex) {
	q.vertices = append(q.vertices, vertex)
}

func (q *Queue) DeQueue() *Vertex {
	length := len(q.vertices)
	if length != 0 {
		removedVertex := q.vertices[0]
		if length != 1 {
			q.vertices = q.vertices[1:]
		} else {
			q.vertices = []*Vertex{}
		}
		return removedVertex
	} else {
		fmt.Println("Empty Queue")
	}
	return nil
}
