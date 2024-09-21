package binary_tree

func preorderTraversal(root *TreeNode) []int {
	stack := make([]*TreeNode, 0)
	cur := root
	res := make([]int, 0)

	for cur != nil || len(stack) > 0 {
		if cur != nil {
			res = append(res, cur.Val)
			stack = append(stack, cur)
			cur = cur.Left
			continue
		}

		cur = stack[len(stack)-1]
		stack = stack[0 : len(stack)-1]
		cur = cur.Right
	}
	return res
}
