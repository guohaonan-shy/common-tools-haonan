package binary_tree

type BSTIterator struct {
	Values []int
	Cur    int
}

func Constructor(root *TreeNode) BSTIterator {
	return BSTIterator{
		Values: inOrder(root),
		Cur:    -1,
	}
}

func inOrder(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	res := make([]int, 0)
	res = append(inOrder(root.Left), root.Val)
	res = append(res, inOrder(root.Right)...)
	return res
}

func (this *BSTIterator) Next() int {
	this.Cur++
	return this.Values[this.Cur]
}

func (this *BSTIterator) HasNext() bool {
	if this.Cur+1 < len(this.Values) && this.Cur+1 >= 0 {
		return true
	}
	return false
}

type BSTIteratorV2 struct {
	inOrder []int
	cur     int
}

func NewBSTIteratorV2(root *TreeNode) BSTIteratorV2 {
	return BSTIteratorV2{
		inOrder: inOrderV2(root),
		cur:     -1,
	}
}

func inOrderV2(root *TreeNode) []int {

	res := make([]int, 0)
	stack := make([]*TreeNode, 0)
	// storage nodes that we are iterating their left side but still need to iterate the right side
	// therefore, when we pop out from the stack, we have switched the current pointer to cur.right
	cur := root
	/*
		why should we check the cur is nil:
		1. process the corner case in which the root is nil
		2. as we iterate through the left path, if left node is nil, we should rollback to the parent of cur nil node
	*/
	for cur != nil || len(stack) > 0 {
		if cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
			continue
		}
		cur = stack[len(stack)-1]
		stack = stack[0 : len(stack)-1]
		res = append(res, cur.Val)
		cur = cur.Right
	}
	return res
}

func (iter *BSTIteratorV2) HasNext() bool {
	next := iter.cur + 1
	if next < len(iter.inOrder) {
		return true
	}
	return false
}

func (iter *BSTIteratorV2) Next() int {
	iter.cur++
	val := iter.inOrder[iter.cur]
	return val
}
