package linkedlist

import (
	"fmt"

	"go.uber.org/zap"
)

var Logger *zap.Logger

type Node struct {
	val  int
	next *Node
}

type LinkedList struct {
	head   *Node
	length int
}

func (l *LinkedList) Insert(data int) {
	if l.head == nil {
		l.head = &Node{val: data, next: nil}
		l.length += 1
		Logger.Info("Head node created", zap.Int("Head value", data))
		return
	} else {
		ptr := l.head
		for i := 0; i < l.length; i += 1 {
			if ptr.next == nil {
				ptr.next = &Node{val: data, next: nil}
				l.length += 1
				Logger.Info("Node inserted at tail", zap.Int("Node value", data), zap.Int("LinkedList length", l.length))
				break
			}
			ptr = ptr.next
		}
	}
}

func (l *LinkedList) Print() {
	ptr := l.head
	if ptr == nil {
		Logger.Error("The linkedlist is empty")
		return
	}
	for i := 0; i < l.length; i += 1 {
		fmt.Println(ptr.val)
		ptr = ptr.next
	}
}
