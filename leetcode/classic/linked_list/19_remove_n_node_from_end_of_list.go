package linked_list

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// calculate the length of list
	length := 0
	for cur := head; cur != nil; cur = cur.Next {
		length++
	}
	pre, cur := (*ListNode)(nil), head
	for i := 1; i < length-n+1; i++ {
		pre = cur
		cur = cur.Next
	}
	if pre == nil {
		temp := head.Next
		head.Next = nil
		return temp
	}
	pre.Next = cur.Next
	return head
}
