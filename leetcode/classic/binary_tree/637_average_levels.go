package binary_tree

func averageOfLevels(root *TreeNode) []float64 {
	if root == nil {
		return []float64{}
	}

	queue := make([]*struct {
		node  *TreeNode
		depth int
	}, 0)

	dict := make(map[int]*struct {
		sum float64
		cnt int
	}, 0)

	queue = append(queue, &struct {
		node  *TreeNode
		depth int
	}{node: root, depth: 1})

	maxdepth := 1
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		if cur.depth > maxdepth {
			maxdepth = cur.depth
		}

		if info, ok := dict[cur.depth]; ok {
			info.sum += float64(cur.node.Val)
			info.cnt++
		} else {
			dict[cur.depth] = &struct {
				sum float64
				cnt int
			}{sum: float64(cur.node.Val), cnt: 1}
		}

		if cur.node.Left != nil {
			queue = append(queue, &struct {
				node  *TreeNode
				depth int
			}{node: cur.node.Left, depth: cur.depth + 1})
		}

		if cur.node.Right != nil {
			queue = append(queue, &struct {
				node  *TreeNode
				depth int
			}{node: cur.node.Right, depth: cur.depth + 1})
		}
	}
	res := make([]float64, 0)
	for i := 1; i <= maxdepth; i++ {
		res = append(res, dict[i].sum/float64(dict[i].cnt))
	}
	return res
}
