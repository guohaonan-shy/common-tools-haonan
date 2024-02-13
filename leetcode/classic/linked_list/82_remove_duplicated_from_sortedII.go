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
			if fast != nil {
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
