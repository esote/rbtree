package rbtree

import (
	"sort"
	"strconv"
	"testing"
)

func TestFuzzInOrder(t *testing.T) {
	tree := New()
	keys := generateKeys(512)
	for _, key := range keys {
		tree.Insert(key, strconv.Itoa(key))
	}
	sort.Ints(keys)

	inOrder := tree.InOrder()
	for i := 0; inOrder.HasNext(); i++ {
		if i >= len(keys) {
			t.Fatal("fuzz in order HasNext failed")
		}
		if inOrder.Next() != strconv.Itoa(keys[i]) {
			t.Fatal("fuzz in order bad val")
		}
	}
}

func TestFuzzOutOrder(t *testing.T) {
	tree := New()
	keys := generateKeys(512)
	for _, key := range keys {
		tree.Insert(key, strconv.Itoa(key))
	}
	sort.Sort(sort.Reverse(sort.IntSlice(keys)))

	outOrder := tree.OutOrder()
	for i := 0; outOrder.HasNext(); i++ {
		if i >= len(keys) {
			t.Fatal("fuzz out order HasNext failed")
		}
		if outOrder.Next() != strconv.Itoa(keys[i]) {
			t.Fatal("fuzz out order bad val")
		}
	}
}
