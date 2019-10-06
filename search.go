package rbtree

// Search searches for a value in the tree. Returns the value found otherwise
// nil.
func (r *RedBlackTree) Search(value interface{}) interface{} {
	ret := r.searchNode(value)

	if ret == nil {
		return nil
	}

	return ret.value
}

func (r *RedBlackTree) searchNode(value interface{}) *node {
	n := r.root
	for n != nil {
		cmp := r.Compare(value, n.value)
		if cmp == 0 {
			break
		}
		if cmp < 0 {
			n = n.left
		} else {
			n = n.right
		}
	}
	return n
}
