package ConvexHull

import (
	"reflect"
	"testing"
)

func Test_outerTrees(t *testing.T) {
	type args struct {
		points []Point
	}
	tests := []struct {
		name string
		args args
		want []Point
	}{
		{
			name: "Case 1",
			args: args{
				Points{
					Point{X: 1, Y: 6},
					Point{X: 4, Y: 1},
					Point{X: 1, Y: 8},
					Point{X: 6, Y: 56},
					Point{X: 1, Y: 1},
					Point{X: 7, Y: 4},
					Point{X: 8, Y: 56},
					Point{X: 13, Y: 12},
					Point{X: 17, Y: 6},
					Point{X: 12, Y: 13},
					Point{X: 71, Y: 12},
					Point{X: 15, Y: 5},
				},
			},
			want: []Point{
				{X: 1, Y: 1},
				{X: 4, Y: 1},
				{X: 71, Y: 12},
				{X: 8, Y: 56},
				{X: 6, Y: 56},
				{X: 1, Y: 8},
				{X: 1, Y: 6},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := outerTrees(tt.args.points); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("outerTrees() = %v, want %v", got, tt.want)
			}
		})
	}
}
