package SymmetricTree

import (
	"testing"
)

func Test_solution(t *testing.T) {
	type args struct {
		t *Tree
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Case Symmetric",
			args: args{
				t: func() *Tree {
					t := New(4)
					t.Left = New(5)
					t.Right = New(5)
					t.Left.Left = New(2)
					t.Left.Right = New(8)
					t.Right.Right = New(2)
					t.Right.Left = New(8)
					t.Left.Right.Left = New(1)
					t.Right.Left.Right = New(1)
					t.Left.Left.Left = New(9)
					t.Left.Left.Right = New(7)
					t.Right.Right.Right = New(9)
					t.Right.Right.Left = New(7)

					return t
				}(),
			},
			want: true,
		},
		{
			name: "Case Symmetric",
			args: args{
				t: func() *Tree {
					t := New(4)
					t.Left = New(5)
					t.Right = New(5)
					t.Left.Left = New(2)
					t.Left.Right = New(8)
					t.Right.Right = New(2)
					t.Right.Left = New(8)
					t.Left.Right.Left = New(1)
					t.Right.Left.Right = New(1)
					t.Left.Left.Left = New(9)
					t.Left.Left.Right = New(7)
					t.Right.Right.Right = New(9)
					t.Right.Right.Left = New(2)

					return t
				}(),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solution(tt.args.t); got != tt.want {
				t.Errorf("solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
