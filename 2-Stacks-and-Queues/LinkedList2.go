package main

import "golang.org/x/exp/constraints"

type LinkedList2Node struct {
	data     interface{}
	next     *LinkedList2Node
	previous *LinkedList2Node
}

type LinkedList2 struct {
	head   *LinkedList2Node
	end    *LinkedList2Node
	length int
}

func NewLinkedList2() *LinkedList2 {
	return &LinkedList2{
		head:   nil,
		end:    nil,
		length: 0,
	}
}

func NewLinkedList2Node[ordered constraints.Ordered](data ordered) *LinkedList2Node {
	return &LinkedList2Node{
		data:     data,
		next:     nil,
		previous: nil,
	}
}

func InitNewLinkedList2(nodes []*LinkedList2Node) *LinkedList2 {
	ll2 := NewLinkedList2()

	for _, node := range nodes {
		ll2.AddToHead(node)
	}

	return ll2
}

func (ll2 *LinkedList2) AddToHead(n *LinkedList2Node) {
	head := ll2.head
	if head != nil {
		head.previous = n
	}
	ll2.head = n
	ll2.head.next = head

	if ll2.length == 0 {
		end := ll2.end
		if head != nil {
			end.next = n
		}
		ll2.end = n
		ll2.end.previous = end
	}
	ll2.length++
}

func (ll2 *LinkedList2) AddToEnd(n *LinkedList2Node) {
	end := ll2.end
	if end != nil {
		end.next = n
	}
	ll2.end = n
	ll2.end.previous = end

	if ll2.length == 0 {
		head := ll2.head
		if head != nil {
			head.previous = n
		}
		ll2.head = n
		ll2.head.next = head
	}
	ll2.length++
}

func (ll2 *LinkedList2) Next() *LinkedList2Node {
	node := ll2.head
	ll2.head = ll2.head.next
	if ll2.head != nil {
		ll2.head.previous = nil
	}

	ll2.length--

	return node
}

func (ll2 *LinkedList2) Previous() *LinkedList2Node {
	node := ll2.end
	ll2.end = ll2.end.previous
	if ll2.end != nil {
		ll2.end.next = nil
	}

	ll2.length--

	return node
}
