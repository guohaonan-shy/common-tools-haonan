package binary_tree

/*
	postorder: left=>right=>root
	preorder: root=>left=>right   ====>  root=>right=>left ====> reverse output: left=>right=>root
*/

func postorderTraversal(root *TreeNode) []int {
	stack := make([]*TreeNode, 0)
	res := make([]int, 0)
	cur := root

	for cur != nil || len(stack) > 0 {
		if cur != nil {
			res = append(res, cur.Val)
			stack = append(stack, cur)
			cur = cur.Right
			continue
		}

		cur = stack[len(stack)-1]
		stack = stack[0 : len(stack)-1]

		cur = cur.Left

	}

	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res)-1-i] = res[len(res)-1-i], res[i]
	}

	return res
}
