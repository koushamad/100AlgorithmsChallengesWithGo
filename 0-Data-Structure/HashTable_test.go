package DataStructure

import (
	"testing"
)

func TestInitHashTable(t *testing.T) {
	type args struct {
		items map[int]string
	}
	tests := []struct {
		name          string
		args          args
		insert        map[int]string
		searchSuccess map[int]string
		delete        []int
		searchFailed  []int
	}{
		{
			name: "Case 1",
			args: args{
				items: map[int]string{1: "1", 3: "3", 2: "2"},
			},
			insert:        map[int]string{4: "4", 5: "5"},
			searchSuccess: map[int]string{1: "1", 5: "5"},
			delete:        []int{1, 4},
			searchFailed:  []int{1, 4, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hashTable := InitHashTable(tt.args.items)

			for key, value := range tt.insert {
				hashTable.Insert(key, value)
			}

			for key, value := range tt.searchSuccess {
				if got, ok := hashTable.Search(key); got != value && ok == true {
					t.Errorf("HashTable Search() = %v, want %v", got, value)
				}
			}

			for _, key := range tt.delete {
				hashTable.Delete(key)
			}

			for _, key := range tt.searchFailed {
				if got, ok := hashTable.Search(key); got != "" && ok == false {
					t.Errorf("HashTable Search() = %v, want %v", got, false)
				}
			}
		})
	}
}
