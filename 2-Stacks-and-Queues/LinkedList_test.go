package main

import (
	"testing"
)

func TestInitLinkedList(t *testing.T) {
	type args struct {
		nodes []*LinkedListNode
	}
	type gets struct {
		index int
		data  int
	}
	var tests = []struct {
		name   string
		args   args
		add    []*LinkedListNode
		delete []int
		next   []int
		get    []gets
		getAll []int
	}{
		{
			name: "Case 1",
			args: args{
				nodes: []*LinkedListNode{
					NewLinkedListNode(1),
					NewLinkedListNode(2),
					NewLinkedListNode(3),
					NewLinkedListNode(4),
					NewLinkedListNode(5),
				},
			},
			add: []*LinkedListNode{
				NewLinkedListNode(6),
				NewLinkedListNode(7),
			},
			delete: []int{6, 5},
			next:   []int{7, 4},
			get: []gets{
				{
					index: 0,
					data:  3,
				},
				{
					index: 1,
					data:  2,
				},
				{
					index: 2,
					data:  1,
				},
			},
			getAll: []int{3, 2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ll := InitLinkedList(tt.args.nodes)

			for _, data := range tt.add {
				ll.Add(data)
			}

			for _, i := range tt.delete {
				ll.Delete(i)
			}

			for _, data := range tt.next {
				node := ll.Next()
				if node.data != data {
					t.Errorf("Delete() = %v, want %v", node.data, data)
				}
			}

			for _, item := range tt.get {
				node := ll.Get(item.index)
				if node.data != item.data {
					t.Errorf("Get() = %v, want %v", node.data, item.data)
				}
			}

			allNodes := ll.GetAll()

			for i, item := range tt.getAll {
				if allNodes[i].data != item {
					t.Errorf("GetAll() = %v, want %v", allNodes[i].data, item)
				}
			}
		})
	}
}
