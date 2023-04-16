package DataStructure

type QueueArray struct {
	items    []*ListItem
	pointer  int
	capacity int
}

// NewQueueArray create a new queue
func NewQueueArray(capacity int) *QueueArray {
	return &QueueArray{
		items:    make([]*ListItem, capacity),
		pointer:  -1,
		capacity: capacity,
	}
}

// InitQueueArray create and push items to queue
func InitQueueArray(items []*ListItem) *QueueArray {
	queue := NewQueueArray(len(items))

	for _, item := range items {
		queue.Enqueue(item)
	}

	return queue
}

// Enqueue add an item to queue
func (q *QueueArray) Enqueue(item *ListItem) {
	q.pointer++

	if q.pointer == q.capacity {
		q.resize(q.capacity * 2)
	}

	q.items[q.pointer] = item
}

// Dequeue get an item from queue
func (q *QueueArray) Dequeue() *ListItem {
	if q.pointer == -1 {
		return NewListItem("")
	}

	item := q.items[q.pointer]

	if q.pointer < (q.capacity / 3) {
		q.resize(q.capacity / 2)
	}

	q.pointer--
	return item
}

// IsEmpty check queue is empty
func (q *QueueArray) IsEmpty() bool {
	return q.pointer <= 0
}

// resize
func (q *QueueArray) resize(size int) {
	items := make([]*ListItem, size)

	for i := 0; i < q.pointer; i++ {
		items[i] = q.items[i]
	}

	q.items = items
	q.capacity = size
}
