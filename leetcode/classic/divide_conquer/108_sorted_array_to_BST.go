package divide_conquer

import . "github.com/common-tools-haonan/leetcode/classic/binary_tree"

// -10,-3,0,5,9
func convert(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	if len(nums) == 1 {
		return &TreeNode{
			Val: nums[0],
		}
	}

	rootIdx := len(nums) / 2
	root := &TreeNode{
		Val: nums[rootIdx],
	}

	leftNums, rightNums := nums[:rootIdx], nums[rootIdx+1:]

	left, right := convert(leftNums), convert(rightNums)
	root.Left, root.Right = left, right
	return root
}
