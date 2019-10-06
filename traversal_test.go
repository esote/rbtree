package rbtree

import (
	"sort"
	"strconv"
	"testing"
)

func TestFuzzInOrder(t *testing.T) {
	tree := New(compare)
	keys := generateKeys(512)
	for _, key := range keys {
		tree.Insert(key)
	}
	sort.Ints(keys)

	inOrder := tree.InOrder()
	for i := 0; inOrder.HasNext(); i++ {
		if i >= len(keys) {
			t.Fatal("fuzz in order HasNext failed")
		}
		if inOrder.Next() != keys[i] {
			t.Fatal("fuzz in order bad val")
		}
	}
}

func TestFuzzOutOrder(t *testing.T) {
	tree := New(compare)
	keys := generateKeys(512)
	for _, key := range keys {
		tree.Insert(key)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(keys)))

	outOrder := tree.OutOrder()
	for i := 0; outOrder.HasNext(); i++ {
		if i >= len(keys) {
			t.Fatal("fuzz out order HasNext failed")
		}
		if outOrder.Next() != keys[i] {
			t.Fatal("fuzz out order bad val")
		}
	}
}

// Benchmark time taken to traverse the entire red-black tree.
func BenchmarkInOrderGrowth(b *testing.B) {
	top := 6
	if testing.Short() {
		top /= 2
	}

	sizes := []int{256}
	for i := 1; i < top; i++ {
		sizes = append(sizes, sizes[i-1]*2)
	}

	for _, size := range sizes {
		b.Run(strconv.Itoa(size), func(sub *testing.B) {
			tree := New(compare)
			keys := generateKeys(size)
			for _, key := range keys {
				tree.Insert(key)
			}
			sub.ResetTimer()
			for i := 0; i < sub.N; i++ {
				for in := tree.InOrder(); in.HasNext(); in.Next() {
				}
			}
		})
	}
}
