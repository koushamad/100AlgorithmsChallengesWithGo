package ConnectionUnion

// QuickFind init N; union N*N; Find 1
type QuickFind struct {
	nodes []int
}

func newQuickFind(count int) *QuickFind {
	nodes := make([]int, count)
	for i := 0; i < count; i++ {
		nodes[i] = i
	}

	return &QuickFind{
		nodes: nodes,
	}
}

func QuickFindInit(nodes int, connections [][]int) *QuickFind {
	qf := newQuickFind(nodes)
	for _, n1n2 := range connections {
		n1 := n1n2[0]
		n2 := n1n2[1]
		qf.Union(n1, n2)
	}

	return qf
}

func (qf *QuickFind) IsConnected(n1, n2 int) bool {
	n1Value := qf.nodes[n1]
	n2Value := qf.nodes[n2]

	return n1Value == n2Value
}

func (qf *QuickFind) Union(n1, n2 int) {

	n1Value := qf.nodes[n1]
	n2Value := qf.nodes[n2]

	if n1Value != n2Value {
		lenNode := len(qf.nodes)

		for i := 0; i < lenNode; i++ {
			if qf.nodes[i] == n2Value {
				qf.nodes[i] = n1Value
			}
		}
	}
}
