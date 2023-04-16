package DataStructure

import (
	"reflect"
	"testing"
)

func TestInitStackLinkedList(t *testing.T) {
	type args struct {
		listItems []string
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
				listItems: []string{"a", "b", "c", "d"},
			},
			push:    []string{"e", "f"},
			pop:     []string{"f", "e", "d", "c"},
			isEmpty: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stack := InitStackLinkedList(tt.args.listItems)

			for _, item := range tt.push {
				stack.Push(NewListItem(item))
			}

			for _, itam := range tt.pop {
				if got := stack.Pup(); !reflect.DeepEqual(got.Value, itam) {
					t.Errorf("StackLinkedList Pup() = %v, want %v", got, itam)
				}
			}
		})
	}
}
