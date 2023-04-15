package main

import "golang.org/x/exp/constraints"

type LinkedListNode struct {
	data interface{}
	next *LinkedListNode
}

type LinkedList struct {
	head   *LinkedListNode
	length int
}

func NewLinkedList() *LinkedList {
	return &LinkedList{
		head:   nil,
		length: 0,
	}
}

func NewLinkedListNode[ordered constraints.Ordered](data ordered) *LinkedListNode {
	return &LinkedListNode{
		data: data,
		next: nil,
	}
}

func InitLinkedList(nodes []*LinkedListNode) *LinkedList {
	ll := NewLinkedList()

	for _, node := range nodes {
		ll.Add(node)
	}

	return ll
}

func (ll *LinkedList) Add(n *LinkedListNode) {
	head := ll.head
	ll.head = n
	ll.head.next = head
	ll.length++
}

func (ll *LinkedList) Delete(index int) {
	head := ll.head
	for i := 1; i < index; i++ {
		head = ll.head.next
	}

	ll.head.next = head.next
	ll.length--
}

func (ll *LinkedList) Next() *LinkedListNode {
	node := ll.head
	ll.head = ll.head.next
	ll.length--

	return node
}

func (ll *LinkedList) Get(index int) *LinkedListNode {
	head := ll.head
	for i := 0; i < index; i++ {
		head = head.next
	}

	return head
}

func (ll *LinkedList) GetAll() []*LinkedListNode {
	llNodes := make([]*LinkedListNode, ll.length)
	head := ll.head

	for i := 0; i < ll.length; i++ {
		llNodes[i] = head
		head = head.next
	}

	return llNodes
}
