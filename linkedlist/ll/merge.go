package ll

func Merge(l1 *LL, l2 *LL) *LL {
	newList := LL{}
	ptr1 := l1.head
	ptr2 := l2.head

	for ptr1 != nil && ptr2 != nil {
		if ptr1.val < ptr2.val {
			newList.Insert(ptr1.val)
			ptr1 = ptr1.next
		} else {
			newList.Insert(ptr2.val)
			ptr2 = ptr2.next
		}
	}
	for ptr1 != nil {
		newList.Insert(ptr1.val)
		ptr1 = ptr1.next
	}
	for ptr2 != nil {
		newList.Insert(ptr2.val)
		ptr2 = ptr2.next
	}
	return &newList
}
