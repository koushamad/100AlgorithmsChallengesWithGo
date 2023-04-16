package DataStructure

import (
	"reflect"
	"testing"
)

func TestInitLinkedList(t *testing.T) {
	type args struct {
		nodes []*LinkedListNode
	}
	tests := []struct {
		name    string
		args    args
		add     []string
		next    []string
		isEmpty bool
	}{
		{
			name: "Case 1",
			args: args{
				[]*LinkedListNode{
					NewLinkedListNode(NewListItem("kousha")),
					NewLinkedListNode(NewListItem("samira")),
					NewLinkedListNode(NewListItem("hamed")),
					NewLinkedListNode(NewListItem("nousha")),
				},
			},
			add:     []string{"maman", "baba"},
			next:    []string{"baba", "maman", "nousha"},
			isEmpty: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ll := InitLinkedList(tt.args.nodes)

			for _, item := range tt.add {
				ll.Add(NewLinkedListNode(NewListItem(item)))
			}

			for _, item := range tt.next {
				if got := ll.Next(); got.Data.Value != item {
					t.Errorf("LinkedList Next() = %v, want %v", got, item)
				}
			}

			if got := ll.IsEmpty(); !reflect.DeepEqual(got, tt.isEmpty) {
				t.Errorf("LinkedList IsEmpty() = %v, want %v", got, tt.isEmpty)
			}
		})
	}
}
