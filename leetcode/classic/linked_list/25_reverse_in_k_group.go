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
	end := head
	for i := 1; i < k && end != nil; i++ {
		end = end.Next
	}

	if end == nil {
		return head
	}

	newHead := end.Next // store the next iteration head before reverse, or we will stay in loop
	pre, cur := (*ListNode)(nil), head

	for pre != end {
		temp := cur.Next
		cur.Next = pre
		pre = cur
		cur = temp
	}

	next := reverseKGroupV2(newHead, k)
	head.Next = next
	return end
}
