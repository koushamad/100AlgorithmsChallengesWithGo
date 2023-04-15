package DataStructure

type QueueLinkedList struct {
	LinkedList *LinkedListTwo
}

func NewQueueLinkedList() *QueueLinkedList {
	return &QueueLinkedList{
		LinkedList: NewLinkedListTwo(),
	}
}

func InitQueueLinkedList(nodes []*LinkedListTwoNode) *QueueLinkedList {
	q := NewQueueLinkedList()
	for _, node := range nodes {
		q.Enqueue(node)
	}

	return q
}

func (q *QueueLinkedList) IsEmpty() bool {
	return q.LinkedList.IsEmpty()
}

func (q *QueueLinkedList) Enqueue(node *LinkedListTwoNode) {
	q.LinkedList.AddToHead(node)
}

func (q *QueueLinkedList) Dequeue() *LinkedListTwoNode {
	return q.LinkedList.Previous()
}
