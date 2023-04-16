package DataStructure

type LinkedListTwoNode struct {
	Data     *ListItem
	Next     *LinkedListTwoNode
	Previous *LinkedListTwoNode
}

type LinkedListTwo struct {
	Head   *LinkedListTwoNode
	End    *LinkedListTwoNode
	Length int
}

func NewLinkedListTwo() *LinkedListTwo {
	return &LinkedListTwo{
		Head:   nil,
		End:    nil,
		Length: 0,
	}
}

func NewLinkedListTwoNode(data *ListItem) *LinkedListTwoNode {
	return &LinkedListTwoNode{
		Data:     data,
		Next:     nil,
		Previous: nil,
	}
}

func InitNewLinkedListTwo(nodes []*LinkedListTwoNode) *LinkedListTwo {
	ll2 := NewLinkedListTwo()

	for _, node := range nodes {
		ll2.AddToHead(node)
	}

	return ll2
}

func (ll2 *LinkedListTwo) IsEmpty() bool {
	return ll2.Length == 0
}

func (ll2 *LinkedListTwo) AddToHead(n *LinkedListTwoNode) {
	head := ll2.Head
	if head != nil {
		head.Previous = n
	}
	ll2.Head = n
	ll2.Head.Next = head

	if ll2.Length == 0 {
		end := ll2.End
		if head != nil {
			end.Next = n
		}
		ll2.End = n
		ll2.End.Previous = end
	}
	ll2.Length++
}

func (ll2 *LinkedListTwo) AddToEnd(n *LinkedListTwoNode) {
	end := ll2.End
	if end != nil {
		end.Next = n
	}
	ll2.End = n
	ll2.End.Previous = end

	if ll2.Length == 0 {
		head := ll2.Head
		if head != nil {
			head.Previous = n
		}
		ll2.Head = n
		ll2.Head.Next = head
	}
	ll2.Length++
}

func (ll2 *LinkedListTwo) Next() *LinkedListTwoNode {
	node := ll2.Head
	ll2.Head = ll2.Head.Next
	if ll2.Head != nil {
		ll2.Head.Previous = nil
	}

	ll2.Length--

	return node
}

func (ll2 *LinkedListTwo) Previous() *LinkedListTwoNode {
	node := ll2.End
	ll2.End = ll2.End.Previous
	if ll2.End != nil {
		ll2.End.Next = nil
	}

	ll2.Length--

	return node
}
