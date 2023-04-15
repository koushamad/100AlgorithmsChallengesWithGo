package DataStructure

type StackLinkedList struct {
	linkList *LinkedList
}

func NewStackLinkedList() *StackLinkedList {
	return &StackLinkedList{
		linkList: NewLinkedList(),
	}
}

func InitStackLinkedList(nodeData []string) *StackLinkedList {
	s := NewStackLinkedList()
	for _, data := range nodeData {
		s.linkList.Add(NewLinkedListNode(data))
	}

	return s
}

func (s *StackLinkedList) IsEmpty() bool {
	return s.linkList.IsEmpty()
}

func (s *StackLinkedList) Push(node *LinkedListNode) {
	s.linkList.Add(node)
}

func (s *StackLinkedList) Pup() *LinkedListNode {
	return s.linkList.Next()
}