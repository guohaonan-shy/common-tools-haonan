package binary_tree

func kthSmallest(root *TreeNode, k int) int {

	var inOrderF func(*TreeNode)
	res := make([]int, 0)

	inOrderF = func(node *TreeNode) {
		if node == nil {
			return
		}

		if node.Left != nil {
			inOrderF(node.Left)
		}

		res = append(res, node.Val)

		if node.Right != nil {
			inOrderF(node.Right)
		}

		return
	}

	inOrderF(root)

	return res[k-1]
}

func kthSmallest_stack(root *TreeNode, k int) int {

	stack := make([]*TreeNode, 0)
	for {
		// 先遍历左
		for root != nil { // 假如root为空，即root.par.right == nil，则获取下一个遍历节点
			stack = append(stack, root)
			root = root.Left
		}

		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// 此时该节点无左子树
		if k == 1 {
			return node.Val
		}

		k--
		root = root.Right // 按照中序遍历，下一个节点是node.right，如果node.right == nil，则为node.parent
	}
}
