package StacksAndQueues

import (
	"github.com/koushamad/100AlgorithmsChallengesWithGo/2-Data-Structure"
)

type StackLinkedList[ordered Ordered] struct {
	linkList *DataStructure.LinkedList[ordered]
}

func NewStackLinkedList[ordered Ordered]() *StackLinkedList[ordered] {
	return &StackLinkedList[ordered]{
		linkList: DataStructure.NewLinkedList[ordered](),
	}
}

func InitStackLinkedList[ordered Ordered](nodeData []ordered) *StackLinkedList[ordered] {
	s := NewStackLinkedList[ordered]()
	for _, data := range nodeData {
		s.linkList.Add(DataStructure.NewLinkedListNode(data))
	}

	return s
}

func (s *StackLinkedList[ordered]) IsEmpty() bool {
	return s.linkList.Length == 0
}

func (s *StackLinkedList[ordered]) Push(node *DataStructure.LinkedListNode[ordered]) {
	s.linkList.Add(node)
}

func (s *StackLinkedList[ordered]) Pup() *DataStructure.LinkedListNode[ordered] {
	return s.linkList.Next()
}
