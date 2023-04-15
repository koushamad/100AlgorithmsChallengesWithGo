package DataStructure

type LinkedListNode[ordered Ordered] struct {
	Data ordered
	Next *LinkedListNode[ordered]
}

type LinkedList[ordered Ordered] struct {
	Head   *LinkedListNode[ordered]
	Length int
}

func NewLinkedList[ordered Ordered]() *LinkedList[ordered] {
	return &LinkedList[ordered]{
		Head:   nil,
		Length: 0,
	}
}

func NewLinkedListNode[ordered Ordered](data ordered) *LinkedListNode[ordered] {
	return &LinkedListNode[ordered]{
		Data: data,
		Next: nil,
	}
}

func InitLinkedList[ordered Ordered](nodes []*LinkedListNode[ordered]) *LinkedList[ordered] {
	ll := NewLinkedList[ordered]()

	for _, node := range nodes {
		ll.Add(node)
	}

	return ll
}

func (ll *LinkedList[ordered]) Add(n *LinkedListNode[ordered]) {
	head := ll.Head
	ll.Head = n
	ll.Head.Next = head
	ll.Length++
}

func (ll *LinkedList[ordered]) Delete(index int) {
	head := ll.Head
	for i := 1; i < index; i++ {
		head = ll.Head.Next
	}

	ll.Head.Next = head.Next
	ll.Length--
}

func (ll *LinkedList[ordered]) Next() *LinkedListNode[ordered] {
	node := ll.Head
	ll.Head = ll.Head.Next
	ll.Length--

	return node
}

func (ll *LinkedList[ordered]) Get(index int) *LinkedListNode[ordered] {
	head := ll.Head
	for i := 0; i < index; i++ {
		head = head.Next
	}

	return head
}

func (ll *LinkedList[ordered]) GetAll() []*LinkedListNode[ordered] {
	llNodes := make([]*LinkedListNode[ordered], ll.Length)
	head := ll.Head

	for i := 0; i < ll.Length; i++ {
		llNodes[i] = head
		head = head.Next
	}

	return llNodes
}
