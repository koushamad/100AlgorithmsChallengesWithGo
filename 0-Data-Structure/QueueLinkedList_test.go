package DataStructure

import (
	"testing"
)

func TestInitQueueLinkedList(t *testing.T) {
	type args struct {
		nodes []*ListItem
	}
	tests := []struct {
		name    string
		args    args
		enqueue []string
		dequeue []string
		isEmpty bool
	}{
		{
			name: "Case 1",
			args: args{
				nodes: []*ListItem{
					NewListItem("a"),
					NewListItem("b"),
					NewListItem("c"),
					NewListItem("d"),
				},
			},
			enqueue: []string{"e", "f"},
			dequeue: []string{"a", "b", "c", "d", "e", "f"},
			isEmpty: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			queue := InitQueueLinkedList(tt.args.nodes)

			for _, item := range tt.enqueue {
				queue.Enqueue(NewListItem(item))
			}

			for _, item := range tt.dequeue {
				if got := queue.Dequeue(); got.Value != item {
					t.Errorf("QueueLinkedList Dequeue() = %v, want %v", got.Value, item)
				}
			}

			if got := queue.IsEmpty(); got != tt.isEmpty {
				t.Errorf("QueueLinkedList IsEmpty() = %v, want %v", got, tt.isEmpty)
			}
		})
	}
}
