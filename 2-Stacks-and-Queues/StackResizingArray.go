package main

type StackResizingArray struct {
	items  []int
	length int
}

func NewStackResizingArray(capacity int) *StackResizingArray {
	return &StackResizingArray{
		items:  make([]int, capacity),
		length: 0,
	}
}

func InitStackResizingArray(capacity int, items []int) *StackResizingArray {
	s := NewStackResizingArray(capacity)

	for _, item := range items {
		s.Push(item)
	}

	return s
}

func (s *StackResizingArray) Push(item int) {
	if s.length == len(s.items) {
		s.Resize(s.length * 2)
	}

	s.items[s.length] = item
	s.length++
}

func (s *StackResizingArray) Pop() int {
	item := s.items[s.length-1]
	s.length--

	if len(s.items)/4 == s.length {
		s.Resize(len(s.items) / 2)
	}

	return item
}

func (s *StackResizingArray) Resize(capacity int) {
	sNew := NewStackResizingArray(capacity)
	for i := 0; i < s.length; i++ {
		sNew.items[i] = s.items[i]
		sNew.length++
	}

	s.items = sNew.items
	s.length = sNew.length
}
