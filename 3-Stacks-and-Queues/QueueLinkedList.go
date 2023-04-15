package StacksAndQueues

import (
	"github.com/koushamad/100AlgorithmsChallengesWithGo/2-Data-Structure"
)

type QueueLinkedList[ordered Ordered] struct {
	LinkedList *DataStructure.LinkedListTwo[ordered]
}

func NewQueueLinkedList[ordered Ordered]() *QueueLinkedList[ordered] {
	return &QueueLinkedList[ordered]{
		LinkedList: DataStructure.NewLinkedListTwo[ordered](),
	}
}

func InitQueueLinkedList[ordered Ordered](nodes []*DataStructure.LinkedListTwoNode[ordered]) *QueueLinkedList[ordered] {
	q := NewQueueLinkedList[ordered]()
	for _, node := range nodes {
		q.Enqueue(node)
	}

	return q
}

func (q *QueueLinkedList[ordered]) Enqueue(node *DataStructure.LinkedListTwoNode[ordered]) {
	q.LinkedList.AddToHead(node)
}

func (q *QueueLinkedList[ordered]) Dequeue() *DataStructure.LinkedListTwoNode[ordered] {
	return q.LinkedList.Previous()
}
