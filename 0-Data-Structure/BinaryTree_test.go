package DataStructure

import (
	"testing"
)

func TestInitBinaryTreeNode(t *testing.T) {
	type args struct {
		items []*ListItem
	}
	tests := []struct {
		name         string
		args         args
		insert       []string
		update       map[string]string
		searchFind   []string
		delete       []string
		searchFailed []string
	}{
		{
			name: "Case 1",
			args: args{
				items: []*ListItem{
					NewListItem("100"),
					NewListItem("52"),
					NewListItem("203"),
					NewListItem("19"),
					NewListItem("76"),
					NewListItem("150"),
					NewListItem("310"),
				},
			},
			insert:       []string{"7", "24", "88", "276"},
			update:       map[string]string{"7": "65", "24": "21"},
			searchFind:   []string{"150", "19", "276", "76", "21", "65"},
			delete:       []string{"276", "76", "88"},
			searchFailed: []string{"1", "2", "88", "276", "76", "7", "24"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := InitBinaryTreeNode(tt.args.items)

			for _, item := range tt.insert {
				tree.Insert(NewListItem(item))
			}

			for nodeItem, item := range tt.update {
				node := tree.Search(NewListItem(nodeItem))
				tree.Update(node, NewListItem(item))
			}

			for _, item := range tt.searchFind {
				if got := tree.Search(NewListItem(item)); got == nil {
					t.Errorf("BinaryTreeNode Search() = %v, want %v", got.Item.Value, item)
				}
			}

			for _, item := range tt.delete {
				tree.Delete(NewListItem(item))
			}

			for _, item := range tt.searchFailed {
				if got := tree.Search(NewListItem(item)); got != nil {
					t.Errorf("BinaryTreeNode Search() = %v, want %v", got.Item.Value, item)
				}
			}
		})
	}
}
