package main

type StackArray struct {
	nodes  []int
	length int
}

func NewStackArray(capacity int) *StackArray {
	return &StackArray{
		nodes:  make([]int, capacity),
		length: 0,
	}
}

func InitStackArray(capacity int, data []int) *StackArray {
	s := NewStackArray(capacity)
	for _, item := range data {
		s.Push(item)
	}

	return s
}

func (s *StackArray) IsEmpty() bool {
	return s.length == 0
}

func (s *StackArray) IsFull() bool {
	return s.length == len(s.nodes)
}

func (s *StackArray) Push(item int) {
	s.nodes[s.length] = item
	s.length++
}

func (s *StackArray) Pup() int {
	item := s.nodes[s.length-1]
	s.length--

	return item
}
