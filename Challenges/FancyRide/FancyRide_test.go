package FancyRide

import "testing"

func Test_solution(t *testing.T) {
	type args struct {
		l     int
		fares []float64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Case 1",
			args: args{
				l:     30,
				fares: []float64{0.3, 0.5, 0.7, 1, 1.3},
			},
			want: "UberXI",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solution(tt.args.l, tt.args.fares); got != tt.want {
				t.Errorf("solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
