package rbtree

// Traversal defines an interface to traverse the red-black tree.
type Traversal interface {
	// HasNext returns whether there are any more values to be returned.
	HasNext() bool

	// Next retrieves the next value in the tree.
	Next() interface{}
}

type inOrder struct {
	next *node
}

type outOrder struct {
	next *node
}

// InOrder gives an in-order traversal (RNL). The values will be sorted in
// ascending order.
//
// Modifications to the tree during iteration will invalidate the traversal.
func (r *RedBlackTree) InOrder() Traversal {
	if r.root == nil {
		return &inOrder{next: nil}
	}

	return &inOrder{next: r.root.min()}
}

// OutOrder gives an out-order traversal (LNR). The values will be sorted in
// descending order.
//
// Modifications to the tree during iteration will invalidate the traversal.
func (r *RedBlackTree) OutOrder() Traversal {
	if r.root == nil {
		return &outOrder{next: nil}
	}

	return &outOrder{next: r.root.max()}
}

func (t *inOrder) HasNext() bool {
	return t.next != nil
}

func (t *inOrder) Next() (ret interface{}) {
	if !t.HasNext() {
		return nil
	}

	ret = t.next.value

	if t.next.right != nil {
		t.next = t.next.right.min()
		return
	}

	for t.next.parent != nil && t.next.parent.left != t.next {
		t.next = t.next.parent
	}
	t.next = t.next.parent
	return
}

func (t *outOrder) HasNext() bool {
	return t.next != nil
}

func (t *outOrder) Next() (ret interface{}) {
	if !t.HasNext() {
		return nil
	}

	ret = t.next.value

	if t.next.left != nil {
		t.next = t.next.left.max()
		return
	}

	for t.next.parent != nil && t.next.parent.right != t.next {
		t.next = t.next.parent
	}

	t.next = t.next.parent
	return
}
