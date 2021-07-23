package main

import (
	bst "binarysearchtree/BinarySearchTree"
	"fmt"

	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewDevelopment()
	defer logger.Sync()
	bst.Logger = logger
	fmt.Println("Binary Search Tree")
	tree := &bst.BST{}
	//calling inorder on empty tree
	tree.Inorder()
	//adding nodes into the bst
	tree.Insert(20)
	tree.Insert(15)
	tree.Insert(10)
	tree.Insert(28)
	tree.Insert(27)
	//inorder traversal of the bst
	tree.Inorder()
}
