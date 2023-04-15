package StacksAndQueues

import (
	"github.com/koushamad/100AlgorithmsChallengesWithGo/2-Data-Structure"
	"testing"
)

func TestInitQueueLinkedList(t *testing.T) {
	type args struct {
		nodes []*DataStructure.LinkedListTwoNode[int]
	}
	tests := []struct {
		name    string
		args    args
		enqueue []*DataStructure.LinkedListTwoNode[int]
		dequeue []*DataStructure.LinkedListTwoNode[int]
	}{
		{
			name: "Case 1",
			args: args{
				[]*DataStructure.LinkedListTwoNode[int]{
					DataStructure.NewLinkedListTwoNode[int](1),
					DataStructure.NewLinkedListTwoNode[int](2),
					DataStructure.NewLinkedListTwoNode[int](3),
				},
			},
			enqueue: []*DataStructure.LinkedListTwoNode[int]{
				DataStructure.NewLinkedListTwoNode[int](4),
				DataStructure.NewLinkedListTwoNode[int](5),
				DataStructure.NewLinkedListTwoNode[int](6),
			},
			dequeue: []*DataStructure.LinkedListTwoNode[int]{
				DataStructure.NewLinkedListTwoNode[int](1),
				DataStructure.NewLinkedListTwoNode[int](2),
				DataStructure.NewLinkedListTwoNode[int](3),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := InitQueueLinkedList[int](tt.args.nodes)

			for _, node := range tt.enqueue {
				q.Enqueue(node)
			}

			for _, node := range tt.dequeue {
				if got := q.Dequeue(); node.Data != got.Data {
					t.Errorf("Dequeue() = %v, want %v", got.Data, node.Data)
				}
			}
		})
	}
}
