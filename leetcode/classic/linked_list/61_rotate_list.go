package linked_list

func rotateRight(head *ListNode, k int) *ListNode {
	if k == 0 || head == nil || head.Next == nil {
		return head
	}

	cur := head
	length := 1
	for cur.Next != nil {
		length++
		cur = cur.Next
	}

	move := length - k%length // the steps that need to move the prev of new head
	cur.Next = head           // link from tail with head
	for ; move > 0; move-- {
		cur = cur.Next
	}

	newHead := cur.Next
	cur.Next = nil
	return newHead
}

func rotateRightV2(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	length := 0
	tail := new(ListNode)
	tail.Next = head
	for cur := head; cur != nil; cur = cur.Next {
		length++
		tail = tail.Next
	}

	k = k % length

	// however, here is a smarter solution, we connect the tail with the head to make it as a ring;
	// at last, we just cut the connection

	tail.Next = head
	cur := head
	for i := 1; i < length-k; i++ {
		cur = cur.Next
	}
	newHead := cur.Next
	cur.Next = nil
	return newHead

	//if k == 0 {
	//	return head
	//}

	//dummy := &ListNode{
	//	Val:  -101,
	//	Next: head,
	//}
	//pre, cur := dummy, head
	//for i := 1; i <= length-k; i++ {
	//	pre = pre.Next
	//	cur = cur.Next
	//}
	//
	//// k == 0, cur == head
	//pre.Next = nil   // cut
	//dummy.Next = cur // connect
	//tail.Next = head
	//return dummy.Next
}
