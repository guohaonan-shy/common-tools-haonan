package divide_conquer

import (
	. "github.com/common-tools-haonan/leetcode/classic/linked_list"
)

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	if len(lists) == 1 {
		return lists[0]
	}

	mid := len(lists) / 2
	leftList, rightList := lists[:mid], lists[mid:]

	left, right := mergeKLists(leftList), mergeKLists(rightList)

	dummy := &ListNode{
		Val: 0,
	}
	pre := dummy
	for left != nil && right != nil {
		val := 0
		if left.Val < right.Val {
			val = left.Val
			left = left.Next
		} else {
			val = right.Val
			right = right.Next
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
