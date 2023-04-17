package ClosestLocation

import "testing"

func Test_solution(t *testing.T) {
	type args struct {
		address string
		objects [][]int
		names   []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Case 1",
			args: args{
				address: "Cat",
				objects: [][]int{
					{-2, 0},
					{1, 2},
					{2, 1, 2, 4},
					{-3, -1, 4, -1},
				},
				names: []string{
					"Bat building",
					"Cast exhibition",
					"At street",
					"Cat avenue",
				},
			},
			want: "Cat avenue",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solution(tt.args.address, tt.args.objects, tt.args.names); got != tt.want {
				t.Errorf("solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
