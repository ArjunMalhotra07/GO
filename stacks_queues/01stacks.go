package stacksqueues

import "fmt"

type Stack struct {
	items []int
}

func (s *Stack) Push(n int) {
	s.items = append(s.items, n)
}

func (s *Stack) Pop() int {
	lastIndex := len(s.items) - 1
	if lastIndex >= 0 {
		removedItem := s.items[len(s.items)-1]
		s.items = s.items[:len(s.items)-1]
		return removedItem
	} else {
		fmt.Println("Empty Stack")
	}
	return -1
}

func DoStackOpeartions() {
	myStack := &Stack{}
	myStack.Push(100)
	myStack.Push(200)
	myStack.Push(500)
	fmt.Println("Popped", myStack.Pop())
	fmt.Println("Popped", myStack.Pop())
	fmt.Println("Popped", myStack.Pop())
	fmt.Println("Popped", myStack.Pop())

	fmt.Println(myStack.items)
}
