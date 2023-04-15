package DataStructure

type LinkedListTwoNode[ordered Ordered] struct {
	Data     ordered
	Next     *LinkedListTwoNode[ordered]
	Previous *LinkedListTwoNode[ordered]
}

type LinkedListTwo[ordered Ordered] struct {
	Head   *LinkedListTwoNode[ordered]
	End    *LinkedListTwoNode[ordered]
	Length int
}

func NewLinkedListTwo[ordered Ordered]() *LinkedListTwo[ordered] {
	return &LinkedListTwo[ordered]{
		Head:   nil,
		End:    nil,
		Length: 0,
	}
}

func NewLinkedListTwoNode[ordered Ordered](data ordered) *LinkedListTwoNode[ordered] {
	return &LinkedListTwoNode[ordered]{
		Data:     data,
		Next:     nil,
		Previous: nil,
	}
}

func InitNewLinkedListTwo[ordered Ordered](nodes []*LinkedListTwoNode[ordered]) *LinkedListTwo[ordered] {
	ll2 := NewLinkedListTwo[ordered]()

	for _, node := range nodes {
		ll2.AddToHead(node)
	}

	return ll2
}

func (ll2 *LinkedListTwo[ordered]) AddToHead(n *LinkedListTwoNode[ordered]) {
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

func (ll2 *LinkedListTwo[ordered]) AddToEnd(n *LinkedListTwoNode[ordered]) {
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

func (ll2 *LinkedListTwo[ordered]) Next() *LinkedListTwoNode[ordered] {
	node := ll2.Head
	ll2.Head = ll2.Head.Next
	if ll2.Head != nil {
		ll2.Head.Previous = nil
	}

	ll2.Length--

	return node
}

func (ll2 *LinkedListTwo[ordered]) Previous() *LinkedListTwoNode[ordered] {
	node := ll2.End
	ll2.End = ll2.End.Previous
	if ll2.End != nil {
		ll2.End.Next = nil
	}

	ll2.Length--

	return node
}
