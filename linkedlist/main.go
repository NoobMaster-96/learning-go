package main

import (
	"fmt"
	ll "linkedlist/linkedlist"

	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewDevelopment()
	defer logger.Sync()
	ll.Logger = logger
	fmt.Println("LinkedList")
	//initialising the first linkedlist
	l1 := &ll.LinkedList{}
	l1.Insert(1)
	l1.Insert(3)
	l1.Insert(7)
	l1.Insert(8)
	l1.Insert(11)
	l1.Print()
	//initialising the second linkedlist
	l2 := &ll.LinkedList{}
	l2.Insert(2)
	l2.Insert(5)
	l2.Insert(9)
	l2.Print()
	//merging bothe the lists
	l3 := ll.Merge(l1, l2)
	l3.Print()
}
