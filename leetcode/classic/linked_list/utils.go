package linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func buildLinkedList(vals []int) *ListNode {
	dummy := &ListNode{}
	pre := dummy
	for _, val := range vals {
		cur := &ListNode{
			Val: val,
		}

		pre.Next = cur
		pre = cur
	}
	return dummy.Next
}
