package binary_tree

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	queue := make([]*TreeNode, 0)
	nextQueue := make([]*TreeNode, 0)
	res := make([][]int, 0)
	isReverse := false
	queue = append(queue, root)

	for len(queue) > 0 {
		level := []int{}
		for i := 0; i < len(queue); i++ {
			cur := queue[i]
			level = append(level, cur.Val)
			if cur.Left != nil {
				nextQueue = append(nextQueue, cur.Left)
			}
			if cur.Right != nil {
				nextQueue = append(nextQueue, cur.Right)
			}
		}

		if isReverse {
			for i := 0; i < len(level)/2; i++ {
				level[i], level[len(level)-1-i] = level[len(level)-1-i], level[i]
			}
		}

		isReverse = !isReverse
		res = append(res, level)
		queue = nextQueue
		nextQueue = []*TreeNode{}
	}
	return res
}
