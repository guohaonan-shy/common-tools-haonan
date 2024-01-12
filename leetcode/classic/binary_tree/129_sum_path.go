package binary_tree

func sumNumbers(root *TreeNode) int {
	var total int = 0
	findPath(root, make([]int, 0, 100), &total)
	return total
}

func findPath(root *TreeNode, prev []int, total *int) {

	if root.Left == nil && root.Right == nil {
		number := 0
		for _, val := range prev {
			number = number*10 + val
		}
		number = number*10 + root.Val
		*total += number
		return
	}

	if root.Left != nil {
		findPath(root.Left, append(prev, root.Val), total)
	}

	if root.Right != nil {
		findPath(root.Right, append(prev, root.Val), total)
	}
}
