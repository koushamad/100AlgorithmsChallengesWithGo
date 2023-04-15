package main

import (
	"testing"
)

func TestInitStackArray(t *testing.T) {
	type args struct {
		capacity int
		data     []int
	}
	tests := []struct {
		name    string
		args    args
		push    []int
		isFull  bool
		pop     []int
		isEmpty bool
	}{
		{
			name: "Case 1",
			args: args{
				capacity: 5,
				data:     []int{1, 2, 3},
			},
			push:    []int{4, 5},
			isFull:  true,
			pop:     []int{5, 4, 3, 2, 1},
			isEmpty: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := InitStackArray(tt.args.capacity, tt.args.data)

			for _, item := range tt.push {
				s.Push(item)
			}

			if got := s.IsFull(); got != tt.isFull {
				t.Errorf("IsFull() = %v, want %v", got, tt.isFull)
			}

			for _, item := range tt.pop {
				if got := s.Pup(); got != item {
					t.Errorf("IsFull() = %v, want %v", got, item)
				}
			}

			if got := s.IsEmpty(); got != tt.isEmpty {
				t.Errorf("IsEmpty() = %v, want %v", got, tt.isEmpty)
			}

		})
	}
}
