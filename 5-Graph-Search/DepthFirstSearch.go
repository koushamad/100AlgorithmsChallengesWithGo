package GraphSearch

import (
	DataStructure "github.com/koushamad/100AlgorithmsChallengesWithGo/0-Data-Structure"
	"strconv"
)

// DepthFirstSearch time: O(e) -> O(n^2), space O(n)
type DepthFirstSearch struct {
	graph *DataStructure.GraphList
	stack *DataStructure.StackLinkedList
}

func (d *DepthFirstSearch) HashPath(src, dst *DataStructure.Vertex) bool {
	if src.GetKey() == dst.GetKey() {
		return true
	}

	for _, vertex := range src.GetAdjacent() {
		key := DataStructure.NewListItem(strconv.Itoa(vertex.GetKey()))
		if !d.stack.Contain(key) {
			d.stack.Push(key)
			if d.HashPath(vertex, dst) {
				return true
			}
			d.stack.Pup()
		}
	}

	return false
}
