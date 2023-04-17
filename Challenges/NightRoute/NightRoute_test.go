package NightRoute

import "testing"

func Test_solution(t *testing.T) {
	type args struct {
		city [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Case 1",
			args: args{
				city: [][]int{
					{-1, 5, 2, 15},
					{2, -1, 0, 3},
					{1, -1, -1, 9},
					{0, 0, 0, -1},
				},
			},
			want: 8,
		},
		{
			name: "Case 2",
			args: args{
				city: [][]int{
					{-1, 5, 20},
					{21, -1, 10},
					{-1, 1, -1},
				},
			},
			want: 15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solution(tt.args.city); got != tt.want {
				t.Errorf("solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
