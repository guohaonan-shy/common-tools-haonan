package binary_tree

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	queue := make([]*struct {
		tree  *TreeNode
		depth int
	}, 0)
	queue = append(queue, &struct {
		tree  *TreeNode
		depth int
	}{tree: root, depth: 1})

	res := make([][]int, 0)
	maxdepth := 1
	level := make([]int, 0)

	for len(queue) > 0 {
		cur, depth := queue[0].tree, queue[0].depth
		queue = queue[1:]

		if depth > maxdepth {
			maxdepth = depth
			res = append(res, level)
			level = []int{}
		}

		level = append(level, cur.Val)
		if cur.Left != nil {
			queue = append(queue, &struct {
				tree  *TreeNode
				depth int
			}{tree: cur.Left, depth: depth + 1})
		}

		if cur.Right != nil {
			queue = append(queue, &struct {
				tree  *TreeNode
				depth int
			}{tree: cur.Right, depth: depth + 1})
		}
	}

	if len(level) > 0 {
		res = append(res, level)
	}
	return res
}
