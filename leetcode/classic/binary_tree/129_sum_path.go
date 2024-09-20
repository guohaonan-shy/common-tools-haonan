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

var (
	total int
)

func sumNumbersV2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	prefix := make([]int, 0)
	sumProcessor(root.Left, append(prefix, root.Val))
	sumProcessor(root.Right, append(prefix, root.Val))
	return total
}

func sumProcessor(root *TreeNode, prefix []int) {
	if root.Left == nil && root.Right == nil {
		// leaf node
		number := 0
		prefix = append(prefix, root.Val)
		for i := 0; i < len(prefix); i++ {
			number = number*10 + prefix[i]
		}
		//number = number*10 + root.Val
		total += number
		return
	}

	if root.Left != nil {
		sumProcessor(root.Left, append(prefix, root.Val))
	}

	if root.Right != nil {
		sumProcessor(root.Right, append(prefix, root.Val))
	}
	return
}
