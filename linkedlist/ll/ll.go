package ll

import "fmt"

type Node struct {
	val  int
	next *Node
}

type LL struct {
	head   *Node
	length int
}

func (l *LL) Insert(data int) {
	if l.head == nil {
		l.head = &Node{val: data, next: nil}
		l.length += 1
		return
	} else {
		ptr := l.head
		for i := 0; i < l.length; i += 1 {
			if ptr.next == nil {
				ptr.next = &Node{val: data, next: nil}
				l.length += 1
				break
			}
			ptr = ptr.next
		}
	}
}

func (l *LL) Print() {
	ptr := l.head
	for i := 0; i < l.length; i += 1 {
		fmt.Println(ptr.val)
		ptr = ptr.next
	}
}
