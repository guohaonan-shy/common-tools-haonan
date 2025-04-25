package binary_tree

func PreorderTraversal(root *TreeNode) []int {
	stack := make([]*TreeNode, 0)
	cur := root
	/*
		'cur' 指的是当前需要进行完整前序遍历的节点
		'stack' 存的则是需要完成右子树的二叉树节点
	*/
	res := make([]int, 0)
	/*
		继续遍历的条件:
		1. 当前存在仍需完整前序遍历的节点
		or
		2. 存在“已完成左子树的遍历” 但仍需遍历右子树的节点
	*/
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
