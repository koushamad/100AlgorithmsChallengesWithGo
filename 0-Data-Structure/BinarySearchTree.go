package DataStructure

type BinaryTreeNode struct {
	Item   *ListItem
	Left   *BinaryTreeNode
	Right  *BinaryTreeNode
	Parent *BinaryTreeNode
}

// NewBinaryTreeNode create a new child node
func NewBinaryTreeNode(item *ListItem) *BinaryTreeNode {
	return &BinaryTreeNode{
		Item:   item,
		Left:   nil,
		Right:  nil,
		Parent: nil,
	}
}

// InitBinaryTreeNode create and add nodes
func InitBinaryTreeNode(items []*ListItem) *BinaryTreeNode {
	var tree *BinaryTreeNode
	lenItems := len(items)

	if lenItems <= 0 {
		return tree
	} else {
		tree = NewBinaryTreeNode(items[0])
	}

	for i := 1; i < lenItems; i++ {
		tree.Insert(items[i])
	}

	return tree
}

// Insert a ListItem
func (t *BinaryTreeNode) Insert(item *ListItem) {
	node := NewBinaryTreeNode(item)
	t.insert(node)
}

// Delete item from tree
func (t *BinaryTreeNode) Delete(item *ListItem) {
	node := t.Search(item)
	if node != nil {
		t.delete(node)
	}
}

func (t *BinaryTreeNode) Update(node *BinaryTreeNode, item *ListItem) {
	t.delete(node)
	node = NewBinaryTreeNode(item)
	t.insert(node)
}

func (t *BinaryTreeNode) delete(node *BinaryTreeNode) {
	leftNode := node.Left
	rightNode := node.Right
	ParentNode := node.Parent

	if ParentNode.Left != nil && ParentNode.Left.Item.Value == node.Item.Value {
		node.Parent.Left = nil
	} else {
		node.Parent.Right = nil
	}

	if leftNode != nil {
		t.insert(leftNode)
	}
	if rightNode != nil {
		t.insert(rightNode)
	}
}

func (t *BinaryTreeNode) insert(node *BinaryTreeNode) {
	if t.Item.Value < node.Item.Value {
		if t.Right == nil {
			t.Right = node
		} else {
			node.Parent = t.Right
			t.Right.insert(node)
		}
	} else if t.Item.Value > node.Item.Value {
		if t.Left == nil {
			t.Left = node
		} else {
			node.Parent = t.Left
			t.Left.insert(node)
		}
	}
}

// Search a ListItem
func (t *BinaryTreeNode) Search(item *ListItem) *BinaryTreeNode {
	if t.Item.Value < item.Value {
		if t.Right == nil {
			return nil
		}

		return t.Right.Search(item)
	} else if t.Item.Value > item.Value {
		if t.Left == nil {
			return nil
		}

		return t.Left.Search(item)
	} else {
		return t
	}
}
