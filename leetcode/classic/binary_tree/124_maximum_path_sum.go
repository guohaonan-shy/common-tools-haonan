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

var maxSum int

func maxPathSumV2(root *TreeNode) int {
	maxSum = math.MinInt32
	if root == nil {
		return 0
	}

	_ = maxPathProcessorV2(root)
	return maxSum
}

func maxPathProcessorV2(root *TreeNode) int {
	if root == nil {
		return 0
	}

	maxPath := root.Val

	leftSum := maxPathProcessorV2(root.Left)
	rightSum := maxPathProcessorV2(root.Right)

	if leftSum <= 0 && rightSum <= 0 {
		maxSum = max(maxSum, maxPath)
		// maxPath = maxPath
	} else if leftSum > 0 && rightSum > 0 {
		maxSum = max(maxPath+leftSum+rightSum, maxSum)
		maxPath += max(leftSum, rightSum)
	} else {
		maxSum = max(maxPath+max(leftSum, rightSum), maxSum)
		maxPath += max(leftSum, rightSum)
	}

	return maxPath
}
