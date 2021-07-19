package main

import (
	"fmt"
	ll "linkedlist/ll"
)

func main() {
	fmt.Println("LinkedList")
	l1 := &ll.LL{}
	l1.Insert(1)
	l1.Insert(3)
	l1.Insert(7)
	l1.Insert(8)
	l1.Insert(11)
	l1.Print()
	l2 := &ll.LL{}
	l2.Insert(2)
	l2.Insert(5)
	l2.Insert(9)
	l2.Print()
}
