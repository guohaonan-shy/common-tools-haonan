package binary_tree

// 判定条件：
// 1.当p,q两个节点分别在某个节点的左右时，即该节点为最低的祖先节点
// 2. 当p,q两个节点在某个节点的左侧或者右侧时，则目标节点，一定是p或者q
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val == p.Val || root.Val == q.Val {
		return root
	}

	left, right := lowestCommonAncestor(root.Left, p, q), lowestCommonAncestor(root.Right, p, q)

	if left != nil && right != nil {
		return root
	}

	if left == nil {
		return right
	}
	return left
}
