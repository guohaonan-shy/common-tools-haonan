package slide_window

import "container/heap"

// 可以通过，但是时间复杂度较高
func maxSlidingWindow(nums []int, k int) []int {
	if k > len(nums) {
		return []int{}
	}

	res := make([]int, 0)
	maxIdx := -1
	for left := 0; left <= len(nums)-k; left++ {
		if maxIdx < left {
			maxIdx = left + maxInSlice(nums[left:left+k])
		} else {
			right := left + k - 1
			if nums[maxIdx] < nums[right] {
				maxIdx = right
			}
		}
		res = append(res, nums[maxIdx])
	}
	return res
}

func maxInSlice(nums []int) int {
	maxVal := 0
	for i, val := range nums {
		if val > nums[maxVal] {
			maxVal = i
		}
	}
	return maxVal
}

type Node struct {
	Val, Idx int // val用来对堆内元素进行排序，idx存下来，用于判断是否还在窗口
}

type MaxHeap struct {
	Nums []Node
}

func (h MaxHeap) Len() int {
	return len(h.Nums)
}

func (h MaxHeap) Less(i, j int) bool {
	return h.Nums[i].Val > h.Nums[j].Val
}

func (h MaxHeap) Swap(i, j int) {
	h.Nums[i], h.Nums[j] = h.Nums[j], h.Nums[i]
}

func (h *MaxHeap) Push(ele any) {
	node := ele.(Node)
	h.Nums = append(h.Nums, node)
}

func (h *MaxHeap) Pop() any {
	res := h.Nums[h.Len()-1]
	h.Nums = h.Nums[:h.Len()-1]
	return res
}

func maxSlidingWindow_heap(nums []int, k int) []int {

	if k > len(nums) {
		return []int{}
	}

	res := make([]int, 0)

	initH := &MaxHeap{
		Nums: make([]Node, 0),
	}
	for i := 0; i < k; i++ {
		initH.Nums = append(initH.Nums, Node{
			Val: nums[i],
			Idx: i,
		})
	}
	heap.Init(initH)
	res = append(res, initH.Nums[0].Val)
	for i := 1; i <= len(nums)-k; i++ {

		heap.Push(initH, Node{
			Val: nums[i+k-1],
			Idx: i + k - 1,
		})

		maxNode := initH.Nums[0]
		for maxNode.Idx < i {
			_ = heap.Pop(initH).(Node)
			maxNode = initH.Nums[0]
		}

		res = append(res, maxNode.Val)

	}
	return res
}

func maxSlidingWindow_queue(nums []int, k int) []int {

	queue := make([]int, 0) // 存元素索引，按照索引对应值的大小，左边的值大于右边

	var push func(idx int)
	push = func(idx int) {
		for len(queue) > 0 && nums[queue[len(queue)-1]] < nums[idx] { // queue最多存储情况是nums为单调递减，那么queue内存的是第一大值的索引，第二大值的索引 一直到 第k大值的索引
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, idx)
	}

	for i := 0; i < k; i++ {
		push(i)
		queue = append(queue, i)
	}

	res := []int{nums[queue[0]]}
	for idx := k; idx < len(nums); idx++ {
		push(idx)

		maxIdx := queue[0]
		for maxIdx < idx-k+1 {
			queue = queue[1:]
			maxIdx = queue[0]
		}

		res = append(res, nums[maxIdx])
	}

	return res
}
