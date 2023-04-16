package DataStructure

import (
	"reflect"
	"testing"
)

func TestInitStackArray(t *testing.T) {
	type args struct {
		items []*ListItem
	}
	tests := []struct {
		name    string
		args    args
		push    []string
		pop     []string
		isEmpty bool
	}{
		{
			name: "Case 1",
			args: args{
				items: []*ListItem{
					NewListItem("1"),
					NewListItem("2"),
				},
			},
			push:    []string{"3", "4"},
			pop:     []string{"4", "3", "2", "1", ""},
			isEmpty: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stack := InitStackArray(tt.args.items)

			for _, item := range tt.push {
				stack.Push(NewListItem(item))
			}

			for _, item := range tt.pop {
				if got := stack.Pop(); got.Value != item {
					t.Errorf("StackArray Pop() = %v, want %v", got, item)
				}
			}

			if got := stack.IsEmpty(); !reflect.DeepEqual(got, tt.isEmpty) {
				t.Errorf("StackArray IsEmpty() = %v, want %v", got, tt.isEmpty)
			}
		})
	}
}
