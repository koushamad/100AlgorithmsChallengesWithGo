package GasStation

import "testing"

func Test_solution(t *testing.T) {
	type args struct {
		gas  []int
		cost []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Case 1",
			args: args{
				gas:  []int{1, 5, 3, 3, 5, 3, 1, 3, 4, 5},
				cost: []int{5, 2, 2, 8, 2, 4, 2, 5, 1, 2},
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solution(tt.args.gas, tt.args.cost); got != tt.want {
				t.Errorf("solution() = %v, want %v", got, tt.want)
			}
			if got := solution2(tt.args.gas, tt.args.cost); got != tt.want {
				t.Errorf("solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
