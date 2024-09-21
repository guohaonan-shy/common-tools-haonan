package binary_tree

import "math"

var diff int

func getMinimumDifference(root *TreeNode) int {
	diff = math.MaxInt32
	diffProcessor(root)
	return diff
}

func diffProcessor(root *TreeNode) {

	/*
		because this is a BST, all left nodes are less than root, and all right nodes are greater than root
		the minimum diff is among abs(root - max in left) and abs(root - min in right)
	*/

	maxInLeft, minInRight := getMax(root.Left), getMin(root.Right)
	diff = min(diff, min(abs(root.Val-maxInLeft), abs(root.Val-minInRight)))
	if root.Left != nil {
		diffProcessor(root.Left)
	}
	if root.Right != nil {
		diffProcessor(root.Right)
	}

	return
}

func getMax(root *TreeNode) int {
	if root == nil {
		return math.MinInt32
	}

	// the rightmost node
	cur := root
	for ; cur.Right != nil; cur = cur.Right {

	}
	return cur.Val
}

func getMin(root *TreeNode) int {
	if root == nil {
		return math.MaxInt32
	}

	// the leftmost node
	cur := root
	for ; cur.Left != nil; cur = cur.Left {

	}
	return cur.Val
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(val int) int {
	if val < 0 {
		return -val
	}
	return val
}
