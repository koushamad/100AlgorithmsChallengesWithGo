package ParkingSpot

import "testing"

func Test_solution(t *testing.T) {
	type args struct {
		carDimensions []int
		parkingLot    [][]int
		luckySpot     []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Case 1",
			args: args{
				carDimensions: []int{3, 2},
				parkingLot: [][]int{
					{1, 0, 1, 0, 1, 0},
					{0, 0, 0, 0, 1, 0},
					{0, 0, 0, 0, 0, 1},
					{1, 0, 1, 1, 1, 1},
				},
				luckySpot: []int{1, 1, 2, 3},
			},
			want: true,
		},
		{
			name: "Case 2",
			args: args{
				carDimensions: []int{3, 2},
				parkingLot: [][]int{
					{1, 0, 1, 0, 1, 0},
					{1, 0, 0, 0, 1, 0},
					{0, 0, 0, 0, 0, 1},
					{1, 0, 0, 0, 1, 1},
				},
				luckySpot: []int{1, 1, 2, 3},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solution(tt.args.carDimensions, tt.args.parkingLot, tt.args.luckySpot); got != tt.want {
				t.Errorf("solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
