package SymmetricTree

type Tree struct {
	root  int
	Left  *Tree
	Right *Tree
}

func New(value int) *Tree {
	return &Tree{
		root:  value,
		Left:  nil,
		Right: nil,
	}
}

func solution(t *Tree) bool {
	return isSymmetric(t.Left, t.Right)
}

func isSymmetric(t1 *Tree, t2 *Tree) bool {
	result := true

	if t1.root != t2.root {
		return false
	}

	if t1.Left != nil || t2.Right != nil {
		result = result && isSymmetric(t1.Left, t2.Right)

	}

	if t1.Right != nil || t2.Left != nil {
		result = result && isSymmetric(t1.Right, t2.Left)
	}

	return result
}
