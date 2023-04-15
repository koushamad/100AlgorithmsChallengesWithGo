package main

import (
	"testing"
)

func TestWeightedQuickUnionInit(t *testing.T) {
	type args struct {
		nodes       int
		connections [][]int
	}
	tests := []struct {
		name    string
		args    args
		success [][]int
		failed  [][]int
	}{
		{
			name: "case 1",
			args: args{
				nodes:       10,
				connections: [][]int{{1, 2}, {4, 5}, {6, 4}, {3, 6}},
			},
			success: [][]int{{1, 2}, {3, 5}},
			failed:  [][]int{{1, 6}, {1, 0}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wqu := WeightedQuickUnionInit(tt.args.nodes, tt.args.connections)
			for _, n := range tt.success {
				n1 := n[0]
				n2 := n[1]
				c := wqu.IsConnected(n1, n2)
				if c != true {
					t.Errorf("WeightedQuickUnionInit = %v, %d,%d  want True", wqu.nodes, n1, n2)
				}
			}

			for _, n := range tt.failed {
				n1 := n[0]
				n2 := n[1]
				c := wqu.IsConnected(n1, n2)
				if c == true {
					t.Errorf("WeightedQuickUnionInit = %v, %d,%d  want False", wqu.nodes, n1, n2)
				}
			}
		})
	}
}
