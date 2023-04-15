package ElementarySorts

import "testing"

func TestShellSort(t *testing.T) {
	type args[ordered Ordered] struct {
		list     []ordered
		callback OrderedCallback[ordered]
	}
	type testCase[ordered Ordered] struct {
		name string
		args args[ordered]
		want []ordered
	}
	tests := []testCase[int]{
		{
			name: "Case 1",
			args: args[int]{
				list: []int{1, 34, 3, 23, 6, 9, 22, 1, 5, 6, 9, 77, 5, 4, 3},
				callback: func(a, b int) bool {
					return a < b
				},
			},
			want: []int{1, 1, 3, 3, 4, 5, 5, 6, 6, 9, 9, 22, 23, 34, 77},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := tt.args.list
			ShellSort[int](list, tt.args.callback)
			for i, w := range tt.want {
				if got := list[i]; got != w {
					t.Errorf("SelectionSort() = %v, want %v", got, w)
				}
			}
		})
	}
}
