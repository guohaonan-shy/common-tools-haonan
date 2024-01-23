package heap

import (
	"container/heap"
	"sort"
)

type IntHeap struct {
	Nums sort.IntSlice
}

func (h IntHeap) Len() int {
	return len(h.Nums)
}

func (h IntHeap) Less(i, j int) bool {
	return h.Nums[i] > h.Nums[j]
}

func (h IntHeap) Swap(i, j int) {
	h.Nums[i], h.Nums[j] = h.Nums[j], h.Nums[i]
}

func (h *IntHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	h.Nums = append(h.Nums, x.(int))
}

func (h *IntHeap) Pop() any {
	old := h.Nums
	n := len(old)
	x := old[n-1]
	h.Nums = old[0 : n-1]
	return x
}

func findKthLargest(nums []int, k int) int {
	h := &IntHeap{
		Nums: nums,
	}
	heap.Init(h)
	var ans any
	for i := k; i > 0; i-- {
		ans = heap.Pop(h)
	}
	return ans.(int)
}
