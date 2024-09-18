package linked_list

func partition(head *ListNode, x int) *ListNode {
	low, high := &ListNode{Val: -101}, &ListNode{Val: -101}
	dummyLow, dummyHigh := low, high
	for cur := head; cur != nil; cur = cur.Next {
		if cur.Val < x {
			low.Next = &ListNode{
				Val: cur.Val,
			}
			low = low.Next
		} else {
			high.Next = &ListNode{
				Val: cur.Val,
			}
			high = high.Next
		}
	}

	low.Next = dummyHigh.Next
	return dummyLow.Next
}
