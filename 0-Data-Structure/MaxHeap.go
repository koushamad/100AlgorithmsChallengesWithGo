package DataStructure

import "strconv"

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
		pointer:  0,
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
	h.pointer--

	if h.pointer < (h.capacity / 3) {
		h.resize(h.capacity / 2)
	}

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
