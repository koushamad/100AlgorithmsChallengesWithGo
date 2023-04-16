package DataStructure

import (
	"testing"
)

func TestInitQueueArray(t *testing.T) {
	type args struct {
		items []*ListItem
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
				items: []*ListItem{
					NewListItem("a"),
					NewListItem("b"),
					NewListItem("c"),
					NewListItem("d"),
				},
			},
			enqueue: []string{"e", "f", "g", "h", "i", "j", "k", "l", "m", "n"},
			dequeue: []string{"n", "m", "l", "k", "j", "i", "h", "g", "f", "e", "d", "c", "b", "a", ""},
			isEmpty: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			queue := InitQueueArray(tt.args.items)

			for _, item := range tt.enqueue {
				queue.Enqueue(NewListItem(item))
			}

			for _, item := range tt.dequeue {
				if got := queue.Dequeue(); got.Value != item {
					t.Errorf("QueueArray Dequeue() = %v, want %v", got.Value, item)
				}
			}

			if got := queue.IsEmpty(); got != tt.isEmpty {
				t.Errorf("QueueArray IsEmpty() = %v, want %v", got, tt.isEmpty)
			}
		})
	}
}
