package binary_tree

func generateTrees(n int) []*TreeNode {
	return genBSTByInterval(1, n)
}

func genBSTByInterval(start, end int) []*TreeNode {

	if start == end {
		return []*TreeNode{{
			Val: start,
		}}
	}

	res := []*TreeNode{}
	for root := start; root <= end; root++ {

		left := []*TreeNode{nil}
		if root > start {
			left = genBSTByInterval(start, root-1)
		}

		right := make([]*TreeNode, 0)
		if root < end {
			right = []*TreeNode{nil}
		}

		for i := range left {
			for j := range right {
				res = append(res, &TreeNode{
					Val:   root,
					Left:  left[i],
					Right: right[j],
				})
			}
		}
	}
	return res
}
