package binary_tree

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	cnt := 1
	handle(root, &cnt)
	return cnt
}

func handle(root *TreeNode, total *int) {
	if root == nil {
		return
	}

	if root.Right != nil {
		*total += 2
		handle(root.Right, total)
		handle(root.Left, total)
	} else {
		if root.Left != nil {
			*total += 1
		}
	}

	return
}

func countNodesV2(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftCnts := countNodesV2(root.Left)
	rightCnts := countNodesV2(root.Right)

	return 1 + leftCnts + rightCnts

}
