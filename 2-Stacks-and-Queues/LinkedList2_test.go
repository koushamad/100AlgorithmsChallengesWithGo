package main

import (
	"testing"
)

func TestInitNewLinkedList2(t *testing.T) {
	type args struct {
		nodes []*LinkedList2Node
	}
	tests := []struct {
		name      string
		args      args
		addToHead []*LinkedList2Node
		addToEnd  []*LinkedList2Node
		next      []int
		previous  []int
		length    int
	}{
		{
			name: "Case 1",
			args: args{
				nodes: []*LinkedList2Node{
					NewLinkedList2Node(1),
					NewLinkedList2Node(2),
					NewLinkedList2Node(3),
				},
			},
			addToHead: []*LinkedList2Node{
				NewLinkedList2Node(4),
				NewLinkedList2Node(5),
				NewLinkedList2Node(6),
			},
			addToEnd: []*LinkedList2Node{
				NewLinkedList2Node(0),
				NewLinkedList2Node(-1),
				NewLinkedList2Node(-2),
			},
			next:     []int{6, 5, 4},
			previous: []int{-2, -1},
			length:   4,
		},
		{
			name: "Case 2",
			args: args{
				nodes: []*LinkedList2Node{
					NewLinkedList2Node(1),
					NewLinkedList2Node(2),
					NewLinkedList2Node(3),
				},
			},
			addToHead: []*LinkedList2Node{
				NewLinkedList2Node(4),
				NewLinkedList2Node(5),
				NewLinkedList2Node(6),
			},
			addToEnd: []*LinkedList2Node{},
			next:     []int{},
			previous: []int{1, 2, 3, 4, 5, 6},
			length:   0,
		},
		{
			name: "Case 3",
			args: args{
				nodes: []*LinkedList2Node{
					NewLinkedList2Node(1),
					NewLinkedList2Node(2),
					NewLinkedList2Node(3),
				},
			},
			addToHead: []*LinkedList2Node{},
			addToEnd: []*LinkedList2Node{
				NewLinkedList2Node(0),
				NewLinkedList2Node(-1),
				NewLinkedList2Node(-2),
			},
			next:     []int{3, 2, 1, 0, -1, -2},
			previous: []int{},
			length:   0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ll2 := InitNewLinkedList2(tt.args.nodes)

			for _, node := range tt.addToHead {
				ll2.AddToHead(node)
			}

			for _, node := range tt.addToEnd {
				ll2.AddToEnd(node)
			}

			for _, item := range tt.next {
				if got := ll2.Next(); got.data != item {
					t.Errorf("Next() = %v, want %v", got.data, item)
				}
			}

			for _, item := range tt.previous {
				if got := ll2.Previous(); got.data != item {
					t.Errorf("Previous() = %v, want %v", got.data, item)
				}
			}

			if ll2.length != tt.length {
				t.Errorf("length = %v, want %v", ll2.length, tt.length)
			}
		})
	}
}
