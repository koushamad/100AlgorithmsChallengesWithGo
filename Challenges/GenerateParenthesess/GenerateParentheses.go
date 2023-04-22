package GenerateParentheses

func solution(n int) []string {
	return sol(n*2, 0)
}

func sol(n, o int) []string {
	stack := NewStack()
	result := make([]string, 0)

	if n == o && n != 0 {
		return []string{")"}
	}

	if o > 0 {
		stack.Push(")")
		r1 := sol(n-1, o-1)
		for _, r := range r1 {
			result = append(result, stack.ToString()+r)
		}
	}

	if n > 0 {
		stack.Push("(")
		r1 := sol(n-1, o+1)
		for _, r := range r1 {
			result = append(result, stack.ToString()+r)
		}
	}

	return result
}
