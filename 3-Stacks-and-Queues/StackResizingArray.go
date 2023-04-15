package StacksAndQueues

type StackResizingArray[ordered Ordered] struct {
	items  []ordered
	length int
}

func NewStackResizingArray[ordered Ordered](capacity int) *StackResizingArray[ordered] {
	return &StackResizingArray[ordered]{
		items:  make([]ordered, capacity),
		length: 0,
	}
}

func InitStackResizingArray[ordered Ordered](capacity int, items []ordered) *StackResizingArray[ordered] {
	s := NewStackResizingArray[ordered](capacity)

	for _, item := range items {
		s.Push(item)
	}

	return s
}

func (s *StackResizingArray[ordered]) Push(item ordered) {
	if s.length == len(s.items) {
		s.Resize(s.length * 2)
	}

	s.items[s.length] = item
	s.length++
}

func (s *StackResizingArray[ordered]) Pop() ordered {
	item := s.items[s.length-1]
	s.length--

	if len(s.items)/4 == s.length {
		s.Resize(len(s.items) / 2)
	}

	return item
}

func (s *StackResizingArray[ordered]) Resize(capacity int) {
	sNew := NewStackResizingArray[ordered](capacity)
	for i := 0; i < s.length; i++ {
		sNew.items[i] = s.items[i]
		sNew.length++
	}

	s.items = sNew.items
	s.length = sNew.length
}
