package DataStructure

import (
	"testing"
)

func TestInitNewLinkedList2(t *testing.T) {
	type args struct {
		nodes []*LinkedListTwoNode[int]
	}
	tests := []struct {
		name      string
		args      args
		addToHead []*LinkedListTwoNode[int]
		addToEnd  []*LinkedListTwoNode[int]
		next      []int
		previous  []int
		length    int
	}{
		{
			name: "Case 1",
			args: args{
				nodes: []*LinkedListTwoNode[int]{
					NewLinkedListTwoNode[int](1),
					NewLinkedListTwoNode[int](2),
					NewLinkedListTwoNode[int](3),
				},
			},
			addToHead: []*LinkedListTwoNode[int]{
				NewLinkedListTwoNode[int](4),
				NewLinkedListTwoNode[int](5),
				NewLinkedListTwoNode[int](6),
			},
			addToEnd: []*LinkedListTwoNode[int]{
				NewLinkedListTwoNode[int](0),
				NewLinkedListTwoNode[int](-1),
				NewLinkedListTwoNode[int](-2),
			},
			next:     []int{6, 5, 4},
			previous: []int{-2, -1},
			length:   4,
		},
		{
			name: "Case 2",
			args: args{
				nodes: []*LinkedListTwoNode[int]{
					NewLinkedListTwoNode[int](1),
					NewLinkedListTwoNode[int](2),
					NewLinkedListTwoNode[int](3),
				},
			},
			addToHead: []*LinkedListTwoNode[int]{
				NewLinkedListTwoNode[int](4),
				NewLinkedListTwoNode[int](5),
				NewLinkedListTwoNode[int](6),
			},
			addToEnd: []*LinkedListTwoNode[int]{},
			next:     []int{},
			previous: []int{1, 2, 3, 4, 5, 6},
			length:   0,
		},
		{
			name: "Case 3",
			args: args{
				nodes: []*LinkedListTwoNode[int]{
					NewLinkedListTwoNode[int](1),
					NewLinkedListTwoNode[int](2),
					NewLinkedListTwoNode[int](3),
				},
			},
			addToHead: []*LinkedListTwoNode[int]{},
			addToEnd: []*LinkedListTwoNode[int]{
				NewLinkedListTwoNode[int](0),
				NewLinkedListTwoNode[int](-1),
				NewLinkedListTwoNode[int](-2),
			},
			next:     []int{3, 2, 1, 0, -1, -2},
			previous: []int{},
			length:   0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ll2 := InitNewLinkedListTwo[int](tt.args.nodes)

			for _, node := range tt.addToHead {
				ll2.AddToHead(node)
			}

			for _, node := range tt.addToEnd {
				ll2.AddToEnd(node)
			}

			for _, item := range tt.next {
				if got := ll2.Next(); got.Data != item {
					t.Errorf("Next() = %v, want %v", got.Data, item)
				}
			}

			for _, item := range tt.previous {
				if got := ll2.Previous(); got.Data != item {
					t.Errorf("Previous() = %v, want %v", got.Data, item)
				}
			}

			if ll2.Length != tt.length {
				t.Errorf("Length = %v, want %v", ll2.Length, tt.length)
			}
		})
	}
}
