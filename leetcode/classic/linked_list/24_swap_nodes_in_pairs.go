package linked_list

func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	if head.Next == nil {
		return head
	}

	// at least the linked list has two nodes

	remain := swapPairs(head.Next.Next)

	dummy := &ListNode{}
	temp := head.Next

	head.Next.Next = head
	head.Next = remain
	dummy.Next = temp
	return dummy.Next
}
