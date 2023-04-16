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
				list: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			},
			insert:  []int{100, 300, 200, 500, 600, 400, 11, 22, 33, 44, 55, 66, 77, 88, 99},
			extract: []int{600, 500, 400},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			heap := InitMaxHeap(tt.args.list)
			for _, item := range tt.insert {
				heap.Insert(NewHeapItem(item, strconv.Itoa(item)))
			}
			heap.Print()
			for _, item := range tt.extract {
				if got := heap.Extract().order; got != item {
					t.Errorf("HeapExtract() = %v, want %v", got, item)
				}
			}
		})
	}
}
