package linkedlist

import "go.uber.org/zap"

func Merge(l1 *LinkedList, l2 *LinkedList) *LinkedList {
	ptr1 := l1.head
	ptr2 := l2.head

	if ptr1 == nil && ptr2 == nil {
		Logger.Error("Both the LinkedLists are empty")
		return nil
	}
	newList := LinkedList{}

	for ptr1 != nil && ptr2 != nil {
		if ptr1.val < ptr2.val {
			newList.Insert(ptr1.val)
			Logger.Debug("Node from linkedlist 1 added", zap.Int("Node value", ptr1.val))
			ptr1 = ptr1.next
		} else {
			newList.Insert(ptr2.val)
			Logger.Debug("Node from linkedlist 2 added", zap.Int("Node value", ptr2.val))
			ptr2 = ptr2.next
		}
	}
	for ptr1 != nil {
		newList.Insert(ptr1.val)
		Logger.Debug("Node from linkedlist 1 added", zap.Int("Node value", ptr1.val))
		ptr1 = ptr1.next
	}
	for ptr2 != nil {
		newList.Insert(ptr2.val)
		Logger.Debug("Node from linkedlist 2 added", zap.Int("Node value", ptr2.val))
		ptr2 = ptr2.next
	}
	return &newList
}
