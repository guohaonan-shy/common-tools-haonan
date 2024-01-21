package linked_list

func mergeKLists(lists []*ListNode) *ListNode {
	return splitAndMerge(lists, 0, len(lists)-1)
}

func splitAndMerge(lists []*ListNode, left, right int) *ListNode {
	if left == right {
		return lists[left]
	}
	mid := (left + right + 1) / 2
	head1, head2 := splitAndMerge(lists, left, mid-1), splitAndMerge(lists, mid, right)

	// merge
	var prev = &ListNode{}
	dummy := prev
	for head1 != nil && head2 != nil {
		if head1.Val < head2.Val {
			prev.Next = head1
			head1 = head1.Next
		} else {
			prev.Next = head2
			head2 = head2.Next
		}
		prev = prev.Next
	}

	if head1 != nil {
		prev.Next = head1
	}

	if head2 != nil {
		prev.Next = head2
	}

	return dummy.Next
}
