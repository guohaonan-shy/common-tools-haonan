package binary_tree

func buildTree(preorder []int, inorder []int) *TreeNode {

	if len(preorder) == 0 {
		return nil
	}

	rootVal := preorder[0]

	idx := indexByValue(inorder, rootVal)
	leftInorder, rightInorder := inorder[:idx], inorder[idx+1:]

	leftCnt := len(leftInorder)
	leftPreorder, rightPreorder := preorder[1:leftCnt+1], preorder[leftCnt+1:]

	return &TreeNode{
		Val:   rootVal,
		Left:  buildTree(leftPreorder, leftInorder),
		Right: buildTree(rightPreorder, rightInorder),
	}
}

// list内元素均为唯一
func indexByValue(list []int, val int) int {
	for i := range list {
		if list[i] == val {
			return i
		}
	}
	return -1
}

func buildTreeV2(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	rootVal := preorder[0]

	// find the index of root in inorder array
	idx := -1
	for i, value := range inorder {
		if value == rootVal {
			idx = i
			break
		}
	}

	leftCnt := idx

	left := buildTreeV2(preorder[1:1+leftCnt], inorder[0:idx])
	right := buildTreeV2(preorder[1+leftCnt:], inorder[idx+1:])

	return &TreeNode{
		Val:   rootVal,
		Left:  left,
		Right: right,
	}
}
