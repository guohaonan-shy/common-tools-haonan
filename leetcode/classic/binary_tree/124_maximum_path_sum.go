package binary_tree

import "math"

func maxPathSum(root *TreeNode) int {
	res := math.MinInt32
	return max(res, maxPathSumHandle(root, &res))
}

func maxPathSumHandle(root *TreeNode, maxV *int) int {
	if root == nil {
		return 0
	}

	leftPath, rightPath := maxPathSumHandle(root.Left, maxV), maxPathSumHandle(root.Right, maxV)

	maxSub := max(root.Val, root.Val+max(leftPath+rightPath, max(leftPath, rightPath)))

	*maxV = max(*maxV, maxSub)
	return max(max(root.Val+leftPath, root.Val+rightPath), root.Val)
}

//func max(a,b int) int {
//	if a > b {
//		return a
//	}
//	return b
//}
