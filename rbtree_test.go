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

const value = "hi"

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

	const (
		fuzzNoErr = iota
		fuzzSize
		fuzzDelete
	)

	for _, size := range sizes {
		for i := 0; i < count; i++ {
			tree := New()
			keys := generateKeys(size)

			for _, key := range keys {
				tree.Insert(key, value+strconv.Itoa(key))
			}

			var fuzzState = fuzzNoErr

			// Size
			if tree.Size() != size {
				fuzzState = fuzzSize
				goto checkDups
			}

			// Search
			for _, key := range keys {
				if search := tree.Search(key); search == nil {
					t.Fatalf("fuzz %d:%d no key", size, i)
				} else if search != value+strconv.Itoa(key) {
					t.Fatalf("fizz %d:%d bad val", size, i)
				}
			}

			// Delete
			for _, key := range keys {
				if deleted := tree.Delete(key); deleted == nil {
					fuzzState = fuzzDelete
					goto checkDups
				} else if deleted != value+strconv.Itoa(key) {
					t.Fatalf("fuzz %d:%d bad val", size, i)
				}
			}

		checkDups:
			if fuzzState != fuzzNoErr {
				if hasDuplicates(t, keys) {
					continue
				}

				switch fuzzState {
				case fuzzSize:
					t.Fatalf("fuzz %d:%d size incorrect",
						size, i)
				case fuzzDelete:
					t.Fatalf("fuzz %d:%d delete incorrect",
						size, i)
				}
			}
		}
	}
}

func TestDeleteNonexistent(t *testing.T) {
	tree := New()

	if tree.Delete(0) != nil {
		t.Fatal("deleted nonexistent key")
	}
}

func TestInsertDuplicate(t *testing.T) {
	tree := New()

	if !tree.Insert(0, value) {
		t.Fatal("first insert failed")
	}
	if tree.Insert(0, value) {
		t.Fatal("second insert succeeded")
	}
}

func TestSize(t *testing.T) {
	tree := New()

	const size int = 100
	keys := generateKeys(size)

	for _, key := range keys {
		tree.Insert(key, value)
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
	tree := New()
	tree.Insert(0, "c")
	tree.Insert(1, "d")
	tree.Insert(-1, "b")
	tree.Insert(2, "e")
	tree.Insert(-2, "a")

	if tree.Max() != "e" {
		t.Fatal("wrong max")
	}

	if tree.Min() != "a" {
		t.Fatal("wrong min")
	}
}

// Benchmark time taken to do one insert+delete with a pre-existing red-black
// tree.
func BenchmarkGrowth(b *testing.B) {
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
			tree := New()
			keys := generateKeys(size)
			for i := 0; i < size-1; i++ {
				tree.Insert(keys[i], nil)
			}
			sub.ResetTimer()
			for i := 0; i < sub.N; i++ {
				tree.Insert(keys[size-1], nil)
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

// Check if array has duplicates. This could be done with a map, but for large
// lists it will allocate too much memory. Calling this function is very rare
// since the key range is so large.
func hasDuplicates(t *testing.T, keys []int) bool {
	for i := range keys {
		for j := range keys {
			if keys[i] == keys[j] && i != j {
				t.Logf("keys len %d had duplicates", len(keys))
				return true
			}
		}
	}

	return false
}
