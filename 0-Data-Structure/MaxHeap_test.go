package DataStructure

import (
	"strconv"
	"testing"
)

func TestInitMaxHeap(t *testing.T) {
	type args struct {
		list []int
	}
	tests := []struct {
		name    string
		args    args
		insert  []int
		extract []int
	}{
		{
			name: "Case 1",
			args: args{
				list: []int{4, 1, 23, 6, 88, 95, 2, 4, 56, 74, 12, 67, 8, 9, 99, 22, 12, 56},
			},
			insert:  []int{1, 100, 300, 400, 500, 600, 1, 3, 4, 5, 6, 7, 200},
			extract: []int{600, 500, 400, 300, 200, 100},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			heap := InitMaxHeap(tt.args.list)

			for _, item := range tt.insert {
				heap.Insert(NewHeapItem(item, strconv.Itoa(item)))
			}

			for _, item := range tt.extract {
				if got := heap.Extract().order; got != item {
					t.Errorf("HeapExtract() = %v, want %v", got, item)
				}
			}
		})
	}
}
