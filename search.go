package rbtree

// Search searches for a key in the tree. Returns the value found otherwise nil.
func (r *RedBlackTree) Search(key int) interface{} {
	ret := r.searchNode(key)

	if ret == nil {
		return nil
	}

	return ret.value
}

func (r *RedBlackTree) searchNode(key int) *node {
	n := r.root
	for n != nil {
		if n.key == key {
			break
		} else if key < n.key {
			n = n.left
		} else {
			n = n.right
		}
	}
	return n
}
