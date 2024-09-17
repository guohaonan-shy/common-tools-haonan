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

func convertToList(node *ListNode) []int {
	cur := node
	res := make([]int, 0)
	for ; cur != nil; cur = cur.Next {
		res = append(res, cur.Val)
	}
	return res
}
