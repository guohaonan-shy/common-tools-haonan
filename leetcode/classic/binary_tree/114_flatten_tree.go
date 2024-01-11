package binary_tree

func flatten(root *TreeNode) {
	flattenHandle(root)
}

func flattenHandle(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	leftHead := flattenHandle(root.Left)
	rightHead := flattenHandle(root.Right)

	tail := root
	if leftHead != nil {
		tail.Left, tail.Right = nil, leftHead
		cur := tail.Right
		for ; cur.Right != nil; cur = cur.Right {

		}
		tail = cur
	}

	if rightHead != nil {
		tail.Left, tail.Right = nil, rightHead
	}

	return root
}
