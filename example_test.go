package rbtree_test

import (
	"fmt"

	"github.com/esote/rbtree"
)

func ExampleRedBlackTree_InOrder() {
	tree := rbtree.New()

	str := "alphabet"

	// Inserting into the red-black tree will remove duplicate letters.
	for _, c := range str {
		tree.Insert(int(c), c)
	}

	for order := tree.InOrder(); order.HasNext(); {
		fmt.Printf("%c", order.Next())
	}
	// Output:
	// abehlpt
}
