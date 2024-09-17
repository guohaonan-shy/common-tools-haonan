package linked_list

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummy := &ListNode{
		Val:  -501,
		Next: head,
	}
	pre, cur := dummy, head
	leftNode, rightNode := &ListNode{}, &ListNode{}
	preLeft := &ListNode{}
	// find preLeft, left and right
	// [5,3]
	for i := 1; i < right; i++ {
		cur = cur.Next
		pre = pre.Next
		if i == left-1 { // => same as i < left
			leftNode = cur
			preLeft = pre
		}
	}
	rightNode = cur
	//
	cur = leftNode.Next
	pre = leftNode
	rightNext := rightNode.Next
	for cur != rightNext {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	preLeft.Next = rightNode
	leftNode.Next = rightNext
	return dummy.Next
}
