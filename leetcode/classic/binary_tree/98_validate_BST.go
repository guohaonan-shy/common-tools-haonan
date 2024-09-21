package binary_tree

import "math"

func isValidBST(root *TreeNode) bool {
	stack := make([]*TreeNode, 0)
	cur := root

	pre := math.MinInt64
	for cur != nil || len(stack) > 0 {
		if cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
			continue
		}

		cur = stack[len(stack)-1]
		stack = stack[0 : len(stack)-1]

		if cur.Val <= pre {
			return false
		}

		pre = cur.Val
		cur = cur.Right
	}
	return true
}
