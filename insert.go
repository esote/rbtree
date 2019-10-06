package rbtree

// Insert value into the red-black tree. Returns whether a new value was
// inserted.
func (r *RedBlackTree) Insert(value interface{}) bool {
	inserted := &node{value: value}

	if r.root == nil {
		r.root = inserted
	} else {
		n := r.root
		for {
			cmp := r.Compare(value, n.value)
			if cmp == 0 {
				return false
			}
			if cmp < 0 {
				if n.left == nil {
					n.left = inserted
					break
				}
				n = n.left
			} else {
				if n.right == nil {
					n.right = inserted
					break
				}
				n = n.right
			}

		}
		inserted.parent = n
	}

	r.insertCases(inserted)
	r.size++
	return true
}

func (r *RedBlackTree) insertCases(n *node) {
	// case 1
	if n.parent == nil {
		n.color = black
		return
	}

	// case 2
	if n.parent.color == black {
		return
	}

	// case 3
	if n.parent.sibling().defcolor() == red {
		n.parent.color = black
		n.parent.sibling().color = black
		n.parent.parent.color = red
		r.insertCases(n.parent.parent)
		return
	}

	// case 4
	if n == n.parent.right && n.parent == n.parent.parent.left {
		r.rotateLeft(n.parent)
		n = n.left
	} else if n == n.parent.left && n.parent == n.parent.parent.right {
		r.rotateRight(n.parent)
		n = n.right
	}

	// case 5
	n.parent.color = black
	n.parent.parent.color = red
	if n == n.parent.left && n.parent == n.parent.parent.left {
		r.rotateRight(n.parent.parent)
	} else {
		r.rotateLeft(n.parent.parent)
	}
}
