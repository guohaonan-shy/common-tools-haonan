package divide_conquer

import (
	. "github.com/common-tools-haonan/leetcode/classic/linked_list"
)

func sortList(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	// length greater than or equal with 2
	// corner case: length == 2; if we set fast = head, middle point is often tail node => loop => OOM
	slow, fast := head, head.Next

	for ; fast != nil && fast.Next != nil; fast = fast.Next.Next {
		slow = slow.Next
	}
	// break the connectivity
	temp := slow.Next
	slow.Next = nil

	left, right := sortList(head), sortList(temp)

	dummy := &ListNode{
		Val: 0,
	}
	pre := dummy
	for left != nil && right != nil {
		val := 0
		if right.Val < left.Val {
			val = right.Val
			right = right.Next
		} else {
			val = left.Val
			left = left.Next
		}

		cur := &ListNode{
			Val: val,
		}
		pre.Next = cur
		pre = cur
	}

	if left != nil {
		pre.Next = left
	}

	if right != nil {
		pre.Next = right
	}

	return dummy.Next
}
