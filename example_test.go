package rbtree_test

import (
	"fmt"

	"github.com/esote/rbtree"
)

func ExampleRedBlackTree_InOrder() {
	cmp := func(a, b interface{}) int {
		if a.(uint8) == b.(uint8) {
			return 0
		}
		if a.(uint8) < b.(uint8) {
			return -1
		}
		return 1
	}

	tree := rbtree.New(cmp)

	str := "alphabet"

	// Inserting into the red-black tree will remove duplicate letters.
	for i := range str {
		tree.Insert(str[i])
	}

	for order := tree.InOrder(); order.HasNext(); {
		fmt.Printf("%c", order.Next())
	}
	// Output:
	// abehlpt
}
