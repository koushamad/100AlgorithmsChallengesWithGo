package main

import (
	"testing"
)

func TestInitQueueLinkedList(t *testing.T) {
	type args struct {
		nodes []*LinkedList2Node
	}
	tests := []struct {
		name    string
		args    args
		enqueue []*LinkedList2Node
		dequeue []*LinkedList2Node
	}{
		{
			name: "Case 1",
			args: args{
				[]*LinkedList2Node{
					NewLinkedList2Node(1),
					NewLinkedList2Node(2),
					NewLinkedList2Node(3),
				},
			},
			enqueue: []*LinkedList2Node{
				NewLinkedList2Node(4),
				NewLinkedList2Node(5),
				NewLinkedList2Node(6),
			},
			dequeue: []*LinkedList2Node{
				NewLinkedList2Node(1),
				NewLinkedList2Node(2),
				NewLinkedList2Node(3),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := InitQueueLinkedList(tt.args.nodes)

			for _, node := range tt.enqueue {
				q.Enqueue(node)
			}

			for _, node := range tt.dequeue {
				if got := q.Dequeue(); node.data != got.data {
					t.Errorf("Dequeue() = %v, want %v", got.data, node.data)
				}
			}
		})
	}
}
