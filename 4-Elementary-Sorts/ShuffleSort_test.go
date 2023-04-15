package ElementarySorts

import "testing"

func TestShuffleSort(t *testing.T) {
	type args[ordered Ordered] struct {
		list []ordered
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
				list: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			},
			want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := tt.args.list
			ShuffleSort[int](list)
			inSameSort := 0
			for i, w := range tt.want {
				if got := list[i]; got == w {
					inSameSort++
				}
			}
			if inSameSort > 5 {
				t.Errorf("SelectionSort() = %v, want %v", inSameSort, 5)
			}
		})
	}
}
