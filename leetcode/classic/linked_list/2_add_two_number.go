package linked_list

func addTwoNumber(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{
		Val: -1,
	}
	pre := dummy

	step := 0
	for l1 != nil || l2 != nil || step == 1 {
		v1, v2 := 0, 0
		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		}

		sum := (v1 + v2 + step) % 10
		step = (v1 + v2 + step) / 10

		cur := &ListNode{
			Val: sum,
		}

		pre.Next = cur
		pre = cur
	}
	return dummy.Next
}
