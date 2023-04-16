package DataStructure

import (
	"testing"
)

func TestInitNewLinkedListTwo(t *testing.T) {
	type args struct {
		nodes []*LinkedListTwoNode
	}
	tests := []struct {
		name      string
		args      args
		addToHead []string
		addToEnd  []string
		next      []string
		previous  []string
		isEmpty   bool
	}{
		{
			name: "Case 1",
			args: args{
				[]*LinkedListTwoNode{
					NewLinkedListTwoNode(NewListItem("a")),
					NewLinkedListTwoNode(NewListItem("b")),
					NewLinkedListTwoNode(NewListItem("c")),
					NewLinkedListTwoNode(NewListItem("d")),
				},
			},
			addToHead: []string{"e", "f"},
			addToEnd:  []string{"g", "h", "j"},
			next:      []string{"f", "e", "d"},
			previous:  []string{"j", "h", "g"},
			isEmpty:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ll := InitNewLinkedListTwo(tt.args.nodes)

			for _, item := range tt.addToHead {
				ll.AddToHead(NewLinkedListTwoNode(NewListItem(item)))
			}

			for _, item := range tt.addToEnd {
				ll.AddToEnd(NewLinkedListTwoNode(NewListItem(item)))
			}

			for _, item := range tt.next {
				if got := ll.Next().Data.Value; got != item {
					t.Errorf("LinkedListTwo Next() = %v, want %v", got, item)
				}
			}

			for _, item := range tt.previous {
				if got := ll.Previous().Data.Value; got != item {
					t.Errorf("LinkedListTwo Previous() = %v, want %v", got, item)
				}
			}
		})
	}
}
