package PerfectCity

import "testing"

func Test_solution(t *testing.T) {
	type args struct {
		dep []float64
		des []float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Case 1",
			args: args{
				dep: []float64{0.4, 1},
				des: []float64{0.9, 3},
			},
			want: 2.7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solution(tt.args.dep, tt.args.des); got != tt.want {
				t.Errorf("solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
