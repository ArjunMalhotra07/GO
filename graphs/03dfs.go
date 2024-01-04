package graphs

import "fmt"

func DepthFirstSearch(graph *Graph) {
	if len(graph.Vertices) == 0 {
		return
	}
	visitedArray := make([]int, len(graph.Vertices)+1)
	myStack := &Stack{}
	myStack.Push(graph.Vertices[0])
	ansArray := []int{}
	for len(myStack.vertices) != 0 {
		poppedVertex := myStack.Pop()
		if visitedArray[poppedVertex.Key] != 1 {
			visitedArray[poppedVertex.Key] = 1
			ansArray = append(ansArray, poppedVertex.Key)
			for _, v := range poppedVertex.Neighbours {
				if visitedArray[v.Key] != 1 {
					myStack.Push(v)
				}
			}
		}
	}
	fmt.Println("Depth First Search:", ansArray)
}

type Stack struct {
	vertices []*Vertex
}

func (s *Stack) Push(vertex *Vertex) {
	s.vertices = append(s.vertices, vertex)
}

func (s *Stack) Pop() *Vertex {
	length := len(s.vertices)
	if length >= 0 {
		removedvertex := s.vertices[length-1]
		s.vertices = s.vertices[:length-1]
		return removedvertex
	} else {
		fmt.Println("Empty Stack")
	}
	return nil
}
