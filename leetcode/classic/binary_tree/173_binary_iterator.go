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
