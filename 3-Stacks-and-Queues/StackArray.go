package StacksAndQueues

type StackArray[ordered Ordered] struct {
	nodes  []ordered
	length int
}

func NewStackArray[ordered Ordered](capacity int) *StackArray[ordered] {
	return &StackArray[ordered]{
		nodes:  make([]ordered, capacity),
		length: 0,
	}
}

func InitStackArray[ordered Ordered](capacity int, data []ordered) *StackArray[ordered] {
	s := NewStackArray[ordered](capacity)

	for _, item := range data {
		s.Push(item)
	}

	return s
}

func (s *StackArray[ordered]) IsEmpty() bool {
	return s.length == 0
}

func (s *StackArray[ordered]) IsFull() bool {
	return s.length == len(s.nodes)
}

func (s *StackArray[ordered]) Push(item ordered) {
	s.nodes[s.length] = item
	s.length++
}

func (s *StackArray[ordered]) Pup() ordered {
	item := s.nodes[s.length-1]
	s.length--

	return item
}
