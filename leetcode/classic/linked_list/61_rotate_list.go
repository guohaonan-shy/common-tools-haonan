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
