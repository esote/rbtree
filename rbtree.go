// Package rbtree provides a red-black tree implementation.
package rbtree

const (
	red   = false
	black = true
)

// RedBlackTree is a red-black tree.
type RedBlackTree struct {
	root *node
	size int
}

// New constructs an empty red-black tree.
func New() *RedBlackTree {
	return new(RedBlackTree)
}

// Size returns the size of the red-black tree in constant time.
func (r *RedBlackTree) Size() int {
	return r.size
}

// Max returns the maximum value in the tree.
func (r *RedBlackTree) Max() interface{} {
	if r.root == nil {
		return nil
	}

	return r.root.max().value
}

// Min returns the minimum value in the tree.
func (r *RedBlackTree) Min() interface{} {
	if r.root == nil {
		return nil
	}

	return r.root.min().value
}

type node struct {
	key                 int
	value               interface{}
	left, right, parent *node
	color               bool
}

func (n *node) sibling() *node {
	if n == n.parent.left {
		return n.parent.right
	}

	return n.parent.left
}

func (n *node) max() *node {
	for n.right != nil {
		n = n.right
	}
	return n
}

func (n *node) min() *node {
	for n.left != nil {
		n = n.left
	}
	return n
}

func (n *node) defcolor() bool {
	if n == nil {
		return black
	}

	return n.color
}

func (r *RedBlackTree) replace(src, dest *node) {
	if src.parent == nil {
		r.root = dest
	} else {
		if src == src.parent.left {
			src.parent.left = dest
		} else {
			src.parent.right = dest
		}
	}
	if dest != nil {
		dest.parent = src.parent
	}
}

func (r *RedBlackTree) rotateLeft(src *node) {
	dest := src.right
	r.replace(src, dest)
	src.right = dest.left
	if dest.left != nil {
		dest.left.parent = src
	}
	dest.left = src
	src.parent = dest
}

func (r *RedBlackTree) rotateRight(src *node) {
	dest := src.left
	r.replace(src, dest)
	src.left = dest.right
	if dest.right != nil {
		dest.right.parent = src
	}
	dest.right = src
	src.parent = dest
}
