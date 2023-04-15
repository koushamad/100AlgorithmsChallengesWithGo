package main

// WeightedQuickUnion init N; union Log N ^ 2; Find Log N
type WeightedQuickUnion struct {
	nodes []int
	sizes []int
}

func NewWeightedQuickUnion(count int) *WeightedQuickUnion {
	n := make([]int, count)
	s := make([]int, count)

	for i := 0; i < count; i++ {
		n[i] = i
		n[i] = 0
	}

	return &WeightedQuickUnion{
		nodes: n,
		sizes: s,
	}
}

func WeightedQuickUnionInit(nodes int, connections [][]int) *WeightedQuickUnion {
	qf := NewWeightedQuickUnion(nodes)
	for _, n1n2 := range connections {
		n1 := n1n2[0]
		n2 := n1n2[1]
		qf.Union(n1, n2)
	}

	return qf
}

func (wqu *WeightedQuickUnion) root(i int) int {
	for i != wqu.nodes[i] {
		i = wqu.nodes[i]
	}

	return i
}

func (wqu *WeightedQuickUnion) IsConnected(n1, n2 int) bool {
	rootN1 := wqu.root(n1)
	rootN2 := wqu.root(n2)

	return rootN1 == rootN2
}

func (wqu *WeightedQuickUnion) Union(n1, n2 int) {
	rootN1 := wqu.root(n1)
	rootN2 := wqu.root(n2)
	sizeN1 := wqu.sizes[rootN1]
	sizeN2 := wqu.sizes[rootN2]

	if sizeN1 >= sizeN2 {
		wqu.nodes[rootN2] = rootN1
		wqu.sizes[rootN1]++
	} else {
		wqu.nodes[rootN1] = rootN2
		wqu.sizes[rootN2]++
	}
}
