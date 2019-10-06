package rbtree

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"testing"
	"time"
)

func compare(a, b interface{}) int {
	if a.(int) == b.(int) {
		return 0
	}
	if a.(int) < b.(int) {
		return -1
	}
	return 1
}

func TestMain(m *testing.M) {
	seed := time.Now().UnixNano()
	rand.Seed(seed)

	if !flag.Parsed() {
		flag.Parse()
	}
	if testing.Verbose() {
		fmt.Printf("seeded to %d\n", seed)
	}

	os.Exit(m.Run())
}

func TestFuzz(t *testing.T) {
	var (
		count = 10
		more  = 30
		max   = 200
	)
	sizes := []int{1, 2, 3, 4, 5, 1e5}

	if testing.Short() {
		count /= 5
		more /= 3
		max /= 2
	}

	// Additional random sizes
	for i := 0; i < more; i++ {
		sizes = append(sizes, rand.Intn(max)+1)
	}

	for _, size := range sizes {
		for i := 0; i < count; i++ {
			tree := New(compare)
			keys := generateKeys(size)

			for _, key := range keys {
				if !tree.Insert(key) {
					// keys contained duplicates
					continue
				}
			}

			// Size
			if tree.Size() != size {
				t.Fatalf("fuzz %d:%d size incorrect", size, i)
			}

			// Search
			for _, key := range keys {
				if tree.Search(key) != key {
					t.Fatalf("fuzz %d:%d search bad val",
						size, i)
				}
			}

			// Delete
			for _, key := range keys {
				if tree.Delete(key) != key {
					t.Fatalf("fuzz %d:%d delete bad val",
						size, i)
				}
			}
		}
	}
}

func TestDeleteNonexistent(t *testing.T) {
	tree := New(compare)

	if tree.Delete(0) != nil {
		t.Fatal("deleted nonexistent key")
	}
}

func TestInsertDuplicate(t *testing.T) {
	tree := New(compare)

	if !tree.Insert(0) {
		t.Fatal("first insert failed")
	}
	if tree.Insert(0) {
		t.Fatal("second insert succeeded")
	}
}

func TestSize(t *testing.T) {
	tree := New(compare)

	const size int = 100
	keys := generateKeys(size)

	for _, key := range keys {
		tree.Insert(key)
	}

	if tree.Size() != size {
		t.Fatalf("expected %d got %d", size, tree.Size())
	}

	for _, key := range keys {
		tree.Delete(key)
	}

	if tree.Size() != 0 {
		t.Fatalf("expected %d got %d", 0, tree.Size())
	}
}

func TestMinMax(t *testing.T) {
	tree := New(func(a, b interface{}) int {
		if a.(string) == b.(string) {
			return 0
		}
		if a.(string) < b.(string) {
			return -1
		}
		return 1
	})
	tree.Insert("c")
	tree.Insert("e")
	tree.Insert("d")
	tree.Insert("a")
	tree.Insert("b")

	if tree.Max() != "e" {
		t.Fatal("wrong max")
	}

	if tree.Min() != "a" {
		t.Fatal("wrong min")
	}
}

// Benchmark time taken to do one insert+delete with a pre-existing red-black
// tree.
func BenchmarkInsertDeleteGrowth(b *testing.B) {
	top := 6
	if testing.Short() {
		top /= 2
	}

	sizes := []int{256}
	for i := 1; i < top; i++ {
		sizes = append(sizes, sizes[i-1]*4)
	}

	for _, size := range sizes {
		b.Run(strconv.Itoa(size), func(sub *testing.B) {
			tree := New(compare)
			keys := generateKeys(size)
			for i := 0; i < size-1; i++ {
				tree.Insert(keys[i])
			}
			sub.ResetTimer()
			for i := 0; i < sub.N; i++ {
				tree.Insert(keys[size-1])
				tree.Delete(keys[size-1])
			}
		})
	}
}

func generateKeys(n int) []int {
	ret := make([]int, n)

	for i := range ret {
		ret[i] = rand.Int()

		// Half are negative.
		if rand.Intn(2) == 0 {
			ret[i] = -ret[i]
		}
	}

	return ret
}
