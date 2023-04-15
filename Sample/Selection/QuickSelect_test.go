package Selection

import "testing"

func TestFindKthLargest(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Case Max",
			args: args{
				nums: []int{2, 4, 6, 2, 1, 67, 834, 2134, 12, 1, 4, 66, 1, 7, 1, 3},
				k:    0,
			},
			want: 2134,
		},
		{
			name: "Case Mid",
			args: args{
				nums: []int{2, 4, 6, 2, 1, 67, 834, 2134, 12, 1, 4, 66, 1, 7, 1, 3},
				k:    7,
			},
			want: 4,
		},
		{
			name: "Case Min",
			args: args{
				nums: []int{2, 4, 6, 2, 1, 67, 834, 2134, 12, 1, 4, 66, 1, 7, 1, 3},
				k:    15,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindKthLargest(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("FindKthLargest() = %v, want %v", got, tt.want)
			}
		})
	}
}
