package linked_list

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{
		Val: -101,
	}

	pre := dummy

	for list1 != nil || list2 != nil {
		v1, v2 := 101, 101
		if list1 != nil {
			v1 = list1.Val
		}

		if list2 != nil {
			v2 = list2.Val
		}

		minmum := 101
		if v1 < v2 {
			if list1 != nil {
				list1 = list1.Next
			}
			minmum = v1
		} else {
			if list2 != nil {
				list2 = list2.Next
			}
			minmum = v2
		}

		cur := &ListNode{
			Val: minmum,
		}

		pre.Next = cur
		pre = cur
	}
	return dummy.Next
}
