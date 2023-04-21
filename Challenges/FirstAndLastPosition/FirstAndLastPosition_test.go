package FirstAndLastPosition

import (
	"reflect"
	"testing"
)

func Test_solution(t *testing.T) {
	type args struct {
		arr    []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "case 1",
			args: args{
				arr:    []int{2, 4, 5, 5, 5, 5, 5, 7, 9, 9},
				target: 5,
			},
			want: []int{2, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solution(tt.args.arr, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
