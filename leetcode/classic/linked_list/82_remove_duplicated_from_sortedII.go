package linked_list

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dummy := &ListNode{
		Val:  -101,
		Next: head,
	}

	pre := dummy
	slow, fast := head, head.Next

	for fast != nil {
		if slow.Val == fast.Val {
			for fast != nil && fast.Val == slow.Val {
				fast = fast.Next
			}

			pre.Next = fast
			slow = fast
			if fast != nil { // this solution has to guarantee that fast is not equal with slow
				fast = fast.Next
			}
		} else {
			fast = fast.Next
			slow = slow.Next
			pre = pre.Next
		}
	}
	return dummy.Next
}

func deleteDuplicatesV2(head *ListNode) *ListNode {

	// because the linked-list's length is [0, 300]
	if head == nil || head.Next == nil {
		return head
	}

	// length greater than one
	dummy := &ListNode{
		Val:  -1,
		Next: head,
	}
	pre := dummy
	cur, next := head, head.Next

	for next != nil {
		if cur.Val == next.Val {
			// find the first value that is greater than cur.val
			for next != nil && next.Val == cur.Val {
				next = next.Next
			}
			// next might be nil;
			// if next is not nil, cur must not be nil
			pre.Next = next
			cur = next
			if next != nil {
				next = next.Next
			}
		} else {
			cur = cur.Next
			next = next.Next
			pre = pre.Next
		}

		//if cur.Val == next.Val {
		//	next = next.Next
		//} else {
		//	if cur.Next == next {
		//		cur = cur.Next
		//		next = next.Next
		//		pre = pre.Next
		//	} else {
		//		pre.Next = next
		//		cur = next
		//	}
		//}
	}
	//if cur.Next != nil { // cur ~ end are duplicated
	//	pre.Next = nil
	//}
	return dummy.Next
}
