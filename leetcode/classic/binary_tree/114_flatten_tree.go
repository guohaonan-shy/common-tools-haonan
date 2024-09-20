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

func flattenV2(root *TreeNode) {
	flattenProcessor(root)
}

func flattenProcessor(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	left := flattenProcessor(root.Left)
	right := flattenProcessor(root.Right)

	root.Left = nil
	root.Right = left
	pre := root
	for cur := root.Right; cur != nil; cur = cur.Right {
		pre = pre.Right
	}
	pre.Right = right
	return root
}
