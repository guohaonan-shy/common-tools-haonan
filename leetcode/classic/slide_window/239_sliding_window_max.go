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
