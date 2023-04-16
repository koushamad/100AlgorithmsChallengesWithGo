package DataStructure

const TableSize = 7

// HashTable structure BC O(1), AC O(m), WC O(n)
type HashTable struct {
	array []*Bucket
}

// Bucket structure
type Bucket struct {
	Node *BucketNode
}

// BucketNode structure
type BucketNode struct {
	key   int
	value string
	next  *BucketNode
}

// NewBucket create a new bucket
func NewBucket() *Bucket {
	return &Bucket{
		Node: nil,
	}
}

// NewBucketNode Create a new node for bucket
func NewBucketNode(key int, value string) *BucketNode {
	return &BucketNode{
		key:   key,
		value: value,
		next:  nil,
	}
}

// NewHashTable create a new empty hash table
func NewHashTable() *HashTable {
	buckets := make([]*Bucket, TableSize)
	bucketsLen := len(buckets)
	for i := 0; i < bucketsLen; i++ {
		buckets[i] = NewBucket()
	}

	return &HashTable{
		array: buckets,
	}
}

func InitHashTable(items map[int]string) *HashTable {
	hashTable := NewHashTable()

	for key, value := range items {
		hashTable.Insert(key, value)
	}

	return hashTable
}

// Insert to Bucket
func (b *Bucket) Insert(key int, value string) {
	newNode := NewBucketNode(key, value)
	currentNode := b.Node

	if currentNode == nil {
		b.Node = newNode
		return
	}

	for currentNode.next != nil {
		currentNode.next = newNode
	}
}

// Search by key from Bucket
func (b *Bucket) Search(key int) (string, bool) {
	currentNode := b.Node

	for currentNode != nil && currentNode.next != nil {
		if currentNode.key == key {
			return currentNode.value, true
		}
	}

	return "", false
}

// Delete by key from Bucket
func (b *Bucket) Delete(key int) {
	currentNode := b.Node
	previousNode := b.Node

	for currentNode.next != nil {
		if currentNode.key == key {
			if previousNode == currentNode {
				b.Node = nil
			} else {
				previousNode.next = currentNode.next
			}
			break
		}
	}
}

// Insert to HashTable
func (h *HashTable) Insert(key int, value string) {
	hashKey := h.HashFunction(key)
	h.array[hashKey].Insert(key, value)
}

// Search in HashTable
func (h *HashTable) Search(key int) (string, bool) {
	hashKey := h.HashFunction(key)
	return h.array[hashKey].Search(key)
}

// Delete in HashTable
func (h *HashTable) Delete(key int) {
	hashKey := h.HashFunction(key)
	h.array[hashKey].Delete(key)
}

// HashFunction create a hash for key
func (h *HashTable) HashFunction(key int) int {
	return key % TableSize
}
