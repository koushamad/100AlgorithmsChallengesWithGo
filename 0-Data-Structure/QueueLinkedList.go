package DataStructure

type QueueLinkedList struct {
	LinkedList *LinkedListTwo
}

func NewQueueLinkedList() *QueueLinkedList {
	return &QueueLinkedList{
		LinkedList: NewLinkedListTwo(),
	}
}

func InitQueueLinkedList(items []*ListItem) *QueueLinkedList {
	q := NewQueueLinkedList()
	for _, item := range items {
		q.Enqueue(item)
	}

	return q
}

func (q *QueueLinkedList) IsEmpty() bool {
	return q.LinkedList.IsEmpty()
}

func (q *QueueLinkedList) Enqueue(item *ListItem) {
	node := NewLinkedListTwoNode(item)
	q.LinkedList.AddToHead(node)
}

func (q *QueueLinkedList) Dequeue() *ListItem {
	node := q.LinkedList.Previous()
	return node.Data
}
