package SelectedElements

import "testing"

func Test_solution(t *testing.T) {
	type args struct {
		arr []int
		p   int
		q   int
		r   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Case 1",
			args: args{
				arr: []int{3, 1, 0, 5, 1, 6, 5, -1, -100},
				p:   1,
				q:   1,
				r:   1,
			},
			want: -96,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := QuickSortSolution(tt.args.arr, tt.args.p, tt.args.q, tt.args.r); got != tt.want {
				t.Errorf("solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
