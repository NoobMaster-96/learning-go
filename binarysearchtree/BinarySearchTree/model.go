package binarysearchtree

import "fmt"

type Node struct {
	val   int
	left  *Node
	right *Node
}

type BST struct {
	root *Node
}

//insert nodes into bst
func (t *BST) Insert(data int) *BST {
	if t.root == nil {
		t.root = &Node{val: data, left: nil, right: nil}
	} else {
		t.root.insert(data)
	}
	return t
}

func (n *Node) insert(data int) {
	if n == nil {
		return
	} else if data <= n.val {
		if n.left == nil {
			n.left = &Node{val: data, left: nil, right: nil}
		} else {
			n.left.insert(data)
		}
	} else {
		if n.right == nil {
			n.right = &Node{val: data, left: nil, right: nil}
		} else {
			n.right.insert(data)
		}
	}
}

//inorder traversal of bst
func (t *BST) Inorder() {
	if t.root == nil {
		return
	} else {
		inorder(t.root)
	}
}

func inorder(n *Node) {
	if n != nil {
		inorder(n.left)
		fmt.Println(n.val)
		inorder(n.right)
	}
}
