package linked_list

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// calculate the length of list
	length := 0
	for cur := head; cur != nil; cur = cur.Next {
		length++
	}
	dummy := &ListNode{
		Val:  -1,
		Next: head,
	}
	pre, cur := dummy, head
	for i := 1; i < length-n+1; i++ {
		pre = cur
		cur = cur.Next
	}
	// cur is definitely not nil
	pre.Next = cur.Next
	return dummy.Next
}
