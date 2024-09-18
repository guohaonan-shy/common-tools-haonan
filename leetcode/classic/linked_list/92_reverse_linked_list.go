package linked_list

/*
As for reverse, we need to know the pre node of left and the next node of right.
In addition, we have to use the pre node of left and left node to start reverse
*/
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	preLeft, l := (*ListNode)(nil), head
	for i := 1; i < left; i++ {
		preLeft = l
		l = l.Next
	}
	r, rightNext := head, head.Next
	// because right <= length(linked-list), rightNext is always non-nil
	for i := 1; i < right; i++ {
		r = r.Next
		rightNext = rightNext.Next
	}
	// do the reverse
	pre, cur := (*ListNode)(nil), l
	for cur != rightNext {
		temp := cur.Next // Can cur be nil ? - only rightNext is nil, cur is nil.
		cur.Next = pre
		pre = cur
		cur = temp
	}

	if preLeft == nil {
		head = r
		l.Next = rightNext
	} else {
		preLeft.Next = r
		l.Next = rightNext
	}
	return head
}
