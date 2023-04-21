package KthLargestElement

import "testing"

func Test_solution(t *testing.T) {
	type args struct {
		arr []int
		k   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Case 1",
			args: args{
				arr: []int{5, 32, 6, 2, 17, 12, 53, 2, 11, 55, 55, 2, 3, 4, 5, 6, 7, 8},
				k:   4,
			},
			want: 32,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solution(tt.args.arr, tt.args.k); got != tt.want {
				t.Errorf("solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
