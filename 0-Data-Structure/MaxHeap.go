package DataStructure

import (
	"fmt"
	"github.com/xlab/treeprint"
	"strconv"
)

type HeapItem struct {
	order int
	data  string
}

// MaxHeap MaxHeap NLog
type MaxHeap struct {
	array    []HeapItem
	pointer  int
	capacity int
}

func InitMaxHeap(list []int) *MaxHeap {
	heap := NewMaxHeap(len(list))

	for _, item := range list {
		heap.Insert(NewHeapItem(item, strconv.Itoa(item)))
	}

	return heap
}

func NewHeapItem(index int, data string) HeapItem {
	return HeapItem{
		order: index,
		data:  data,
	}
}

func NewMaxHeap(capacity int) *MaxHeap {
	return &MaxHeap{
		array:    make([]HeapItem, capacity),
		pointer:  -1,
		capacity: capacity,
	}
}

func (h *MaxHeap) Insert(item HeapItem) {
	h.pointer++
	if h.pointer == h.capacity {
		h.resize(h.capacity * 2)
	}

	h.array[h.pointer] = item
	h.maxHeapIfyUp(h.pointer)
}

func (h *MaxHeap) Extract() *HeapItem {
	if h.pointer == 0 {
		return nil
	}

	extracted := h.array[0]
	h.array[0] = h.array[h.pointer]

	if h.pointer < (h.capacity / 3) {
		h.resize(h.capacity / 2)
	}

	h.pointer--
	h.maxHeapIfyDown(0)

	return &extracted
}

// maxHeapIfyUp
func (h *MaxHeap) maxHeapIfyUp(index int) {
	for h.array[h.parent(index)].order < h.array[index].order {
		h.swap(h.parent(index), index)
		index = h.parent(index)
	}
}

// maxHeapIfyDown
func (h *MaxHeap) maxHeapIfyDown(index int) {
	var childToCompare int
	lastIndex := h.array[h.pointer].order
	leftIndex, rightIndex := h.left(index), h.right(index)

	for leftIndex <= lastIndex {

		// Check HeapItem Active Capacity
		if leftIndex > h.pointer && rightIndex > h.pointer {
			break
		} else if leftIndex > h.pointer {
			leftIndex = rightIndex
		} else if rightIndex > h.pointer {
			rightIndex = leftIndex
		}

		// Check With Child Ws Bigger
		if leftIndex == lastIndex {
			childToCompare = leftIndex
		} else if h.array[leftIndex].order > h.array[rightIndex].order {
			childToCompare = leftIndex
		} else {
			childToCompare = rightIndex
		}

		// Swap Index With Child
		if h.array[index].order < h.array[childToCompare].order {
			h.swap(index, childToCompare)
			index = childToCompare
			leftIndex, rightIndex = h.left(index), h.right(index)
		} else {
			break
		}
	}
}

// resize
func (h *MaxHeap) resize(size int) {
	array := make([]HeapItem, size)

	for i := 0; i < h.pointer; i++ {
		array[i] = h.array[i]
	}

	h.array = array
	h.capacity = size
}

// swap
func (h *MaxHeap) swap(index1, index2 int) {
	h.array[index1], h.array[index2] = h.array[index2], h.array[index1]
}

// parent
func (h *MaxHeap) parent(index int) int {
	return (index - 1) / 2
}

// left
func (h *MaxHeap) left(index int) int {
	return 2*index + 1
}

// right
func (h *MaxHeap) right(index int) int {
	return 2*index + 2
}

func (h *MaxHeap) Print() {
	tree := treeprint.New()

	h.toTree(tree, 0)

	fmt.Println(tree.String())
}

func (h *MaxHeap) toTree(tree treeprint.Tree, index int) {
	if index > h.pointer {
		return
	}

	value, left, right := h.getParentsChildValues(index)

	if index == 0 {
		if left != "" || right != "" {
			tree = tree.AddBranch(value)
		} else {
			tree = tree.AddNode(value)
		}
	}

	if left != "" {
		leftTree := tree
		if h.left(h.left(index)) > h.pointer {
			leftTree = leftTree.AddNode(left)
		} else {
			leftTree = tree.AddBranch(left)
		}
		h.toTree(leftTree, h.left(index))
	}
	if right != "" {
		rightTree := tree
		if h.left(h.right(index)) > h.pointer {
			rightTree = rightTree.AddNode(right)
		} else {
			rightTree = tree.AddBranch(right)
		}
		h.toTree(rightTree, h.right(index))
	}
}

func (h *MaxHeap) getParentsChildValues(index int) (string, string, string) {
	value, left, right := "", "", ""
	if index <= h.pointer {
		value = h.array[index].data
	}
	if h.left(index) <= h.pointer {
		left = h.array[h.left(index)].data
	}
	if h.right(index) <= h.pointer {
		right = h.array[h.right(index)].data
	}

	return value, left, right
}
