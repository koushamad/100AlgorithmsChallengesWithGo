package main

type QueueLinkedList struct {
	LinkedList *LinkedList2
}

func NewQueueLinkedList() *QueueLinkedList {
	return &QueueLinkedList{
		LinkedList: NewLinkedList2(),
	}
}

func InitQueueLinkedList(nodes []*LinkedList2Node) *QueueLinkedList {
	q := NewQueueLinkedList()
	for _, node := range nodes {
		q.Enqueue(node)
	}

	return q
}

func (q *QueueLinkedList) Enqueue(node *LinkedList2Node) {
	q.LinkedList.AddToHead(node)
}

func (q *QueueLinkedList) Dequeue() *LinkedList2Node {
	return q.LinkedList.Previous()
}
