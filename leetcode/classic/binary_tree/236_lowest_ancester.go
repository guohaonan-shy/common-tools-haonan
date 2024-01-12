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

func lowestCommonAncestor_Parent(root, p, q *TreeNode) *TreeNode {

	parent := make(map[*TreeNode]*TreeNode, 0)
	var dfs func(*TreeNode)
	dfs = func(cur *TreeNode) {
		if cur == nil {
			return
		}

		if cur.Left != nil {
			parent[cur.Left] = cur
			dfs(cur.Left)
		}

		if cur.Right != nil {
			parent[cur.Right] = cur
			dfs(cur.Right)
		}
		return
	}

	dfs(root)

	visited := make(map[*TreeNode]bool, 0)
	for p != nil {
		visited[p] = true
		p = parent[p]
	}

	for q != nil {
		if visited[q] {
			return q
		}
		q = parent[q]
	}
	return root
}
