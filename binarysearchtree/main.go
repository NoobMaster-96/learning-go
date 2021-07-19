package main

import (
	bst "binarysearchtree/bst"
	"fmt"
)

func main() {
	fmt.Println("Binary Search Tree")
	tree := &bst.BST{}
	tree.Insert(20)
	tree.Insert(15)
	tree.Insert(10)
	tree.Insert(28)
	tree.Insert(27)
	tree.Inorder()
}
