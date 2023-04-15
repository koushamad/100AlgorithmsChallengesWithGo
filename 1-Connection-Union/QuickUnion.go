package ConnectionUnion

// QuickUnion init N; union N ^ 2; Find N
type QuickUnion struct {
	nodes []int
}

func NewQuickUnion(count int) *QuickUnion {
	nodes := make([]int, count)

	for i := 0; i < count; i++ {
		nodes[i] = i
	}

	return &QuickUnion{
		nodes: nodes,
	}
}

func QuickUnionInit(nodes int, connections [][]int) *QuickUnion {
	qu := NewQuickUnion(nodes)

	for _, n1n2 := range connections {
		n1 := n1n2[0]
		n2 := n1n2[1]
		qu.Union(n1, n2)
	}

	return qu
}

func (qu *QuickUnion) root(i int) int {
	for i != qu.nodes[i] {
		i = qu.nodes[i]
	}

	return i
}

func (qu *QuickUnion) IsConnected(n1, n2 int) bool {
	rootN1 := qu.root(n1)
	rootN2 := qu.root(n2)

	return rootN1 == rootN2
}

func (qu *QuickUnion) Union(n1, n2 int) {
	rootN1 := qu.root(n1)
	rootN2 := qu.root(n2)

	qu.nodes[rootN1] = rootN2
}
