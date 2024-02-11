package linked_list

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	slow, fast := head, head
	var preEnd *ListNode
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		preEnd = slow
		slow = slow.Next
	}

	preEnd.Next = nil
	head1 := sortList(head)
	head2 := sortList(slow)

	var prev = &ListNode{}
	dummy := prev
	for head1 != nil && head2 != nil {
		if head1.Val < head2.Val {
			prev.Next = &ListNode{
				Val:  head1.Val,
				Next: nil,
			}
			head1 = head1.Next
		} else {
			prev.Next = &ListNode{
				Val:  head2.Val,
				Next: nil,
			}
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
