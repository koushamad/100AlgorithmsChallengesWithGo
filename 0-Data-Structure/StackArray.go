package DataStructure

// StackArray type
type StackArray struct {
	Items    []*ListItem
	pointer  int
	capacity int
}

// NewStackArray New StackArray crate empty stack
func NewStackArray(capacity int) *StackArray {
	return &StackArray{
		Items:    make([]*ListItem, capacity),
		pointer:  -1,
		capacity: capacity,
	}
}

// InitStackArray create a new stack and add items
func InitStackArray(items []*ListItem) *StackArray {
	stack := NewStackArray(len(items))

	for _, item := range items {
		stack.Push(item)
	}

	return stack
}

// IsEmpty if stack is empty it will return true
func (s *StackArray) IsEmpty() bool {
	return s.pointer == -1
}

// Push item to stack
func (s *StackArray) Push(item *ListItem) {
	s.pointer++

	if s.pointer == s.capacity {
		s.resize(s.capacity * 2)
	}

	s.Items[s.pointer] = item
}

// Pop item from stack
func (s *StackArray) Pop() *ListItem {
	if s.pointer < 0 {
		return NewListItem("")
	}

	if s.pointer > 1 && s.pointer < s.capacity/3 {
		s.resize(s.capacity / 2)
	}

	item := s.Items[s.pointer]
	s.Items[s.pointer] = nil
	s.pointer--
	return item
}

// resize the stack size
func (s *StackArray) resize(size int) {
	items := make([]*ListItem, size)

	for i := 0; i < s.pointer; i++ {
		items[i] = s.Items[i]
	}

	s.Items = items
	s.capacity = size
}
