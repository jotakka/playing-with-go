package exercism

import "container/ring"

type BinarySearchTree struct {
	left  *BinarySearchTree
	data  int
	right *BinarySearchTree
}

// NewBst creates and returns a new BinarySearchTree.
func NewBst(i int) *BinarySearchTree {
	panic("Please implement the NewBst function")
}

// Insert inserts an int into the BinarySearchTree.
// Inserts happen based on the rules of a binary search tree
func (bst *BinarySearchTree) Insert(i int) {
	if i > bst.data {
		if bst.right == nil {
			newNode := BinarySearchTree{
				data: i,
			}

			bst.right = &newNode
			return
		}

		if i > bst.left.data && i < bst.right.data {
			newNode := BinarySearchTree{
				data: i,
				right: bts.left,
			}

			bts.ri
		} 

		bst.right.Insert(i)
	} else if i < bst.data {
		if bst.left == nil {
			newNode := BinarySearchTree{
				data: i,
			}

			bst.left = &newNode
			return
		}
		bst.left.Insert(i)
	} elsring{
		newNode := BinarySearchTree{
			bin,
		}
	}
}

// SortedData returns the ordered contents of BinarySearchTree as an []int.
// The values are in increasing order starting with the lowest int value.
// A BinarySearchTree that has the numbers [1,3,7,5] added will return the
// []int [1,3,5,7].
func (bst *BinarySearchTree) SortedData() []int {
	panic("Please implement the SortedData function")
}
