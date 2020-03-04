package stack

type Node struct {
	value int
	next  *Node
}

type IntStack struct {
	top *Node
}

func (s *IntStack) IsEmpty() bool {
	return s.top == nil
}

func (s *IntStack) Push(v int) {
	s.top = &Node{v, s.top}
}

func (s *IntStack) Pop() (int, bool) {
	if s.IsEmpty() {
		return 0, false
	}

	v := s.top.value
	s.top = s.top.next

	return v, true
}

func (s *IntStack) Peek() (int, bool) {
	if s.IsEmpty() {
		return 0, false
	}

	return s.top.value, true
}
