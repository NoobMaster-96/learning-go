package binarysearchtree

import (
	"fmt"

	"go.uber.org/zap"
)

type Node struct {
	val   int
	left  *Node
	right *Node
}

type BST struct {
	root *Node
}

//insert nodes into bst
func (bst *BST) Insert(data int) *BST {
	if bst.root == nil {
		bst.root = &Node{val: data, left: nil, right: nil}
		zap.L().Info("Root node created", zap.Int("Root value", data))
	} else {
		zap.L().Debug("Helper Function called for inserting nodes into Binary Search Tree")
		bst.root.insert(data)
	}
	return bst
}

func (n *Node) insert(data int) {
	if n == nil {
		return
	} else if data <= n.val {
		if n.left == nil {
			n.left = &Node{val: data, left: nil, right: nil}
			zap.L().Info("Left child node created", zap.Int("Parent value", n.val), zap.Int("Child value", data))
		} else {
			n.left.insert(data)
		}
	} else {
		if n.right == nil {
			n.right = &Node{val: data, left: nil, right: nil}
			zap.L().Info("Right child node created", zap.Int("Parent value", n.val), zap.Int("Child value", data))
		} else {
			n.right.insert(data)
		}
	}
}

//inorder traversal of bst
func (bst *BST) Inorder() {
	if bst.root == nil {
		zap.L().Error("No nodes present in the Binary Search Tree")
		return
	} else {
		inorder(bst.root)
	}
}

func inorder(n *Node) {
	if n != nil {
		inorder(n.left)
		fmt.Println(n.val)
		inorder(n.right)
	}
}
