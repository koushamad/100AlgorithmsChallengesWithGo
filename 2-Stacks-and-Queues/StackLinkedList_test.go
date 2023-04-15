package main

import (
	"testing"
)

func TestInitStackLinkedList(t *testing.T) {
	type args struct {
		nodeData []int
	}
	tests := []struct {
		name    string
		args    args
		push    []int
		pop     []int
		isEmpty bool
	}{
		{
			name: "Case 1",
			args: args{
				nodeData: []int{1, 2, 3},
			},
			push:    []int{4, 5},
			pop:     []int{5, 4, 3, 2, 1},
			isEmpty: true,
		},
		{
			name: "Case 2",
			args: args{
				nodeData: []int{1, 2},
			},
			push:    []int{3, 4},
			pop:     []int{4, 3, 2},
			isEmpty: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := InitStackLinkedList(tt.args.nodeData)
			for _, item := range tt.push {
				s.Push(NewLinkedListNode(item))
			}

			for _, item := range tt.pop {
				if got := s.Pup(); got.data != item {
					t.Errorf("Pup() = %v, want %v", got.data, item)
				}
			}

			if got := s.IsEmpty(); got != tt.isEmpty {
				t.Errorf("IsEmpty() = %v, want %v", got, tt.isEmpty)
			}

		})
	}
}
