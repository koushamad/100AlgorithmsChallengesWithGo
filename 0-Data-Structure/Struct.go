package DataStructure

type ListItem struct {
	Value string
}

func NewListItem(item string) *ListItem {
	return &ListItem{item}
}
