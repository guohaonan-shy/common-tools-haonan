package binary_tree

import (
	"strconv"
	"strings"
)

// 迎合leetcode，val都是整数；后续有空可将这部分抽离到data_structure部分（泛型）
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func treeify(tree string) *TreeNode {
	tree = strings.TrimLeft(tree, "[")
	tree = strings.TrimRight(tree, "]")
	elements := strings.Split(tree, ",")
	return madeup(0, elements)
}

func madeup(root int, elements []string) *TreeNode {

	if root >= len(elements) {
		return nil
	}

	str := strings.TrimSpace(elements[root])

	if str == "null" {
		return nil
	}

	val, _ := strconv.Atoi(str)

	left, right := 2*root+1, 2*root+2
	rootNode := &TreeNode{
		Val:   val,
		Left:  madeup(left, elements),
		Right: madeup(right, elements),
	}
	return rootNode
}
