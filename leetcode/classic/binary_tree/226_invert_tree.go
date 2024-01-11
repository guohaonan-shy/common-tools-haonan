package binary_tree

func invertTree(root *TreeNode) *TreeNode {

	if root == nil {
		return nil
	}

	invertedLeft, invertedRight := invertTree(root.Left), invertTree(root.Right)

	root.Left, root.Right = invertedRight, invertedLeft
	return root
}
