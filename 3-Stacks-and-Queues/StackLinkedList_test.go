package StacksAndQueues

import (
	"github.com/koushamad/100AlgorithmsChallengesWithGo/2-Data-Structure"
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
			s := InitStackLinkedList[int](tt.args.nodeData)
			for _, item := range tt.push {
				s.Push(DataStructure.NewLinkedListNode[int](item))
			}

			for _, item := range tt.pop {
				if got := s.Pup(); got.Data != item {
					t.Errorf("Pup() = %v, want %v", got.Data, item)
				}
			}

			if got := s.IsEmpty(); got != tt.isEmpty {
				t.Errorf("IsEmpty() = %v, want %v", got, tt.isEmpty)
			}

		})
	}
}
