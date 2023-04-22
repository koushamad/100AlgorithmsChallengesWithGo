package GenerateParentheses

type Stack struct {
	array    []string
	pointer  int
	capacity int
}

func NewStack() *Stack {
	return &Stack{
		array:    make([]string, 0),
		pointer:  -1,
		capacity: 0,
	}
}

func (s *Stack) Push(val string) {
	s.pointer++
	if s.pointer < s.capacity {
		s.array[s.pointer] = val
	} else {
		s.array = append(s.array, val)
	}
}

func (s *Stack) Pop() string {
	if s.pointer == -1 {
		return ""
	}
	val := s.array[s.pointer]
	s.pointer--
	return val
}

func (s Stack) ToString() string {
	result := ""

	for s.pointer != -1 {
		result += s.Pop()
	}

	return result
}
