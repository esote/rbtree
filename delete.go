package rbtree

// Delete value from the red-black tree. Returns the value deleted or nil if no
// value was found.
func (r *RedBlackTree) Delete(value interface{}) interface{} {
	n := r.searchNode(value)

	if n == nil {
		return nil
	}

	ret := n.value

	if n.left != nil && n.right != nil {
		max := n.left.max()
		n.value = max.value
		n = max
	}

	var child *node

	if n.right == nil {
		child = n.left
	} else {
		child = n.right
	}

	if n.color == black {
		n.color = child.defcolor()
		r.deleteCases(n)
	}

	r.replace(n, child)
	r.size--
	return ret
}

func (r *RedBlackTree) deleteCases(n *node) {
	// case 1
	if n.parent == nil {
		return
	}

	// case 2
	if n.sibling().color == red {
		n.parent.color = red
		n.sibling().color = black
		if n == n.parent.left {
			r.rotateLeft(n.parent)
		} else {
			r.rotateRight(n.parent)
		}
	}

	// case 3
	if n.parent.color == black &&
		n.sibling().defcolor() == black &&
		n.sibling().left.defcolor() == black &&
		n.sibling().right.defcolor() == black {

		n.sibling().color = red
		r.deleteCases(n.parent)
		return
	}

	// case 4
	if n.parent.color == red &&
		n.sibling().color == black &&
		n.sibling().left.defcolor() == black &&
		n.sibling().right.defcolor() == black {

		n.sibling().color = red
		n.parent.color = black
		return
	}

	// case 5
	if n == n.parent.left &&
		n.sibling().color == black &&
		n.sibling().left.defcolor() == red &&
		n.sibling().right.defcolor() == black {

		n.sibling().color = red
		n.sibling().left.color = black
		r.rotateRight(n.sibling())
	} else if n == n.parent.right &&
		n.sibling().color == black &&
		n.sibling().right.defcolor() == red &&
		n.sibling().left.defcolor() == black {

		n.sibling().color = red
		n.sibling().right.color = black
		r.rotateLeft(n.sibling())
	}

	// case 6
	n.sibling().color = n.parent.color
	n.parent.color = black
	if n == n.parent.left {
		n.sibling().right.color = black
		r.rotateLeft(n.parent)
	} else {
		n.sibling().left.color = black
		r.rotateRight(n.parent)
	}
}
