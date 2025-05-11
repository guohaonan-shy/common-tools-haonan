package binary_tree

func numTrees(n int) int {
	return bstCnt(1, n)
}

func bstCnt(start, end int) int {
	if start == end {
		return 1
	}
	cnt := 0
	for root := start; root <= end; root++ {
		left := 1
		if root > start {
			left = bstCnt(start, root-1)
		}

		right := 1
		if root < end {
			right = bstCnt(root+1, end)
		}
		cnt += left * right
	}
	return cnt
}
