package DataStructure

type LinkedListNode struct {
	Data *ListItem
	Next *LinkedListNode
}

type LinkedList struct {
	Head   *LinkedListNode
	Length int
}

func NewLinkedList() *LinkedList {
	return &LinkedList{
		Head:   nil,
		Length: 0,
	}
}

func NewLinkedListNode(data *ListItem) *LinkedListNode {
	return &LinkedListNode{
		Data: data,
		Next: nil,
	}
}

func InitLinkedList(nodes []*LinkedListNode) *LinkedList {
	ll := NewLinkedList()

	for _, node := range nodes {
		ll.Add(node)
	}

	return ll
}

func (ll *LinkedList) IsEmpty() bool {
	return ll.Length == 0
}

func (ll *LinkedList) Add(n *LinkedListNode) {
	head := ll.Head
	ll.Head = n
	ll.Head.Next = head
	ll.Length++
}

func (ll *LinkedList) Next() *LinkedListNode {
	node := ll.Head
	ll.Head = ll.Head.Next
	ll.Length--

	return node
}
