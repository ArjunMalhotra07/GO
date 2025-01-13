package trees

type Stack struct {
	stack []*Node
}

func (root *Node) InOrderTraversal2() []*Node {
	s := &Stack{}
	ans := []*Node{}

	if root == nil {
		return ans
	}
	for root != nil || len(s.stack) != 0 {
		for root != nil {
			s.Push(root)
			root = root.Left
		}
		root = s.Pop()
		ans = append(ans, root)
		root = root.Right
	}
	return ans
}

func (s *Stack) Push(n *Node) {
	s.stack = append(s.stack, n)
}

func (s *Stack) Pop() *Node {
	last := len(s.stack)
	if last != 0 {
		popped := s.stack[last-1]
		if last != 1 {
			s.stack = s.stack[1:]
		} else {
			s.stack = []*Node{}
		}
		return popped

	}
	return nil
}
