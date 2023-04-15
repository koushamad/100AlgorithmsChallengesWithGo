package main

import (
	"testing"
)

func TestInitStackResizingArray(t *testing.T) {
	type args struct {
		capacity int
		items    []int
	}
	tests := []struct {
		name   string
		args   args
		push   []int
		pop    []int
		length int
	}{
		{
			name: "Case 1",
			args: args{
				capacity: 3,
				items:    []int{1, 2},
			},
			push:   []int{3, 4, 5, 6, 7},
			pop:    []int{7, 6, 5, 4, 3, 2},
			length: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := InitStackResizingArray(tt.args.capacity, tt.args.items)

			for _, item := range tt.push {
				s.Push(item)
			}

			for _, item := range tt.pop {
				if got := s.Pop(); got != item {
					t.Errorf("Pop() = %v, want %v", got, item)
				}
			}

			if got := len(s.items); got != tt.length {
				t.Errorf("length = %v, want %v", got, tt.length)
			}
		})
	}
}
