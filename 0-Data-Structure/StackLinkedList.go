package DataStructure

type StackLinkedList struct {
	linkList *LinkedList
}

func NewStackLinkedList() *StackLinkedList {
	return &StackLinkedList{
		linkList: NewLinkedList(),
	}
}

func InitStackLinkedList(listItems []string) *StackLinkedList {
	s := NewStackLinkedList()
	for _, item := range listItems {
		s.linkList.Add(NewLinkedListNode(NewListItem(item)))
	}

	return s
}

func (s *StackLinkedList) IsEmpty() bool {
	return s.linkList.IsEmpty()
}

func (s *StackLinkedList) Push(item *ListItem) {
	s.linkList.Add(NewLinkedListNode(item))
}

func (s *StackLinkedList) Pup() *ListItem {
	return s.linkList.Next().Data
}
