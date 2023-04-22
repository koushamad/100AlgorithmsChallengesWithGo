package GenerateParentheses

import (
	"reflect"
	"testing"
)

func Test_solution(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Case 0",
			args: args{
				n: 0,
			},
			want: []string{},
		},
		{
			name: "Case 1",
			args: args{
				n: 1,
			},
			want: []string{"()"},
		},
		{
			name: "Case 2",
			args: args{
				n: 2,
			},
			want: []string{"()()", "(())"},
		},
		{
			name: "Case 3",
			args: args{
				n: 3,
			},
			want: []string{"()()()", "()(())", "(()))()", "(())())", "(()())"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solution(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
