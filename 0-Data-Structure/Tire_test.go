package DataStructure

import (
	"testing"
)

func TestInitTrie(t *testing.T) {
	type args struct {
		items []string
	}
	tests := []struct {
		name         string
		args         args
		insert       []string
		searchFind   []string
		searchFailed []string
	}{
		{
			name: "Case 1",
			args: args{
				items: []string{"kousha", "samira"},
			},
			insert:       []string{"hamed", "nousha"},
			searchFind:   []string{"hamed", "samira"},
			searchFailed: []string{"soraya", "katayon"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			trie := InitTrie(tt.args.items)

			for _, item := range tt.insert {
				trie.Insert(item)
			}

			for _, item := range tt.searchFind {
				if got := trie.Search(item); got != true {
					t.Errorf("Trie Search() = %v, want %v", got, true)
				}
			}

			for _, item := range tt.searchFailed {
				if got := trie.Search(item); got != false {
					t.Errorf("Trie Search() = %v, want %v", got, true)
				}
			}
		})
	}
}
