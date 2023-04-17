package ConnectedComponentsCount

import "testing"

func Test_solution(t *testing.T) {
	type args struct {
		vertexes [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Case 1",
			args: args{
				vertexes: [][]int{
					{1},
					{1},
					{},
					{5},
					{5},
					{3, 4, 6, 7},
					{5},
					{5},
				},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solution(tt.args.vertexes); got != tt.want {
				t.Errorf("solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
