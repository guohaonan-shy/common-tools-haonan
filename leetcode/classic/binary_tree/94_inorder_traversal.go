package binary_tree

func inorderTraversal(root *TreeNode) []int {
	stack := make([]*TreeNode, 0)
	res := make([]int, 0)

	cur := root // might be nil

	for cur != nil || len(stack) > 0 {
		if cur != nil {
			stack = append(stack, cur)
			cur = cur.Left // iterate left nodes until it is nil
			continue
		}

		cur = stack[len(stack)-1]
		stack = stack[0 : len(stack)-1]
		res = append(res, cur.Val)

		cur = cur.Right
	}
	return res
}
