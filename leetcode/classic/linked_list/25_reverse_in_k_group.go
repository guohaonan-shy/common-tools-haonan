package linked_list

func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy := &ListNode{
		Val:  -1,
		Next: head,
	}

	pre, cur := dummy, head
	end := head
	cnt := 1
	for end != nil {

		if cnt == k {
			temp := end.Next
			newHead := reverse(cur, end)
			cnt = 1
			pre.Next = newHead
			pre, cur = cur, temp
			end = temp
		} else {
			end = end.Next
			cnt++
		}
	}
	pre.Next = cur
	return dummy.Next
}

func reverse(head *ListNode, end *ListNode) *ListNode {
	var (
		pre *ListNode
		cur *ListNode = head
	)

	if end == nil {
		return head
	}

	for cur != end {
		temp := cur.Next
		cur.Next = pre
		pre = cur
		cur = temp
	}
	end.Next = pre
	return end
}

func reverseKGroupV2(head *ListNode, k int) *ListNode {

	if head == nil {
		return head
	}

	var (
		cnt  = 1
		pre  *ListNode
		cur  = head
		tail = head
	)

	for ; cnt < k && tail != nil; cnt++ {
		tail = tail.Next
	}

	if tail == nil || cnt < k {
		return head
	}

	newHead := reverseKGroup(tail.Next, k)

	for pre != tail {
		temp := cur.Next
		cur.Next = pre
		pre = cur
		cur = temp
	}
	head.Next = newHead
	return pre
}
