package binary_tree

func rightSideView(root *TreeNode) []int {

	if root == nil {
		return []int{}
	}

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	res := make([]int, 0)

	nextQueue := make([]*TreeNode, 0)
	for len(queue) > 0 {
		for i := 0; i < len(queue); i++ {
			cur := queue[i]
			if cur.Left != nil {
				nextQueue = append(nextQueue, cur.Left)
			}
			if cur.Right != nil {
				nextQueue = append(nextQueue, cur.Right)
			}
		}

		res = append(res, queue[len(queue)-1].Val)

		queue = nextQueue
		nextQueue = []*TreeNode{}
	}

	return res
}

/*
	Let us implement a solution with the depth of current node
*/

type TreeNodeWithDepth struct {
	treeNode *TreeNode
	depth    int
}

func rightSideViewV2(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	dict := make(map[int]*TreeNode, 0) // key is the depth, index is current the rightmost node under this depth
	queue := make([]*TreeNodeWithDepth, 0)
	queue = append(queue, &TreeNodeWithDepth{
		treeNode: root,
		depth:    1,
	})

	maxdepth := 1
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		depth, node := cur.depth, cur.treeNode

		if depth > maxdepth {
			maxdepth = depth
		}

		dict[depth] = node

		if node.Left != nil {
			queue = append(queue, &TreeNodeWithDepth{
				treeNode: node.Left,
				depth:    depth + 1,
			})
		}

		if node.Right != nil {
			queue = append(queue, &TreeNodeWithDepth{
				treeNode: node.Right,
				depth:    depth + 1,
			})
		}
	}
	res := make([]int, 0)
	for i := 1; i <= maxdepth; i++ {
		res = append(res, dict[i].Val)
	}
	return res
}

func rightSideViewV3(root *TreeNode) []int {

	if root == nil {
		return []int{}
	}

	stack := make([]*TreeNodeWithDepth, 0)
	dict := make(map[int]*TreeNode, 0)
	maxdepth := 1

	stack = append(stack, &TreeNodeWithDepth{
		treeNode: root,
		depth:    1,
	})

	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		stack = stack[0 : len(stack)-1]

		if _, ok := dict[cur.depth]; !ok {
			dict[cur.depth] = cur.treeNode
		}

		if cur.depth > maxdepth {
			maxdepth = cur.depth
		}

		if cur.treeNode.Left != nil {
			stack = append(stack, &TreeNodeWithDepth{
				treeNode: cur.treeNode.Left,
				depth:    cur.depth + 1,
			})
		}

		if cur.treeNode.Right != nil {
			stack = append(stack, &TreeNodeWithDepth{
				treeNode: cur.treeNode.Right,
				depth:    cur.depth + 1,
			})
		}

	}
	res := make([]int, 0)
	for i := 1; i <= maxdepth; i++ {
		res = append(res, dict[i].Val)
	}
	return res
}
