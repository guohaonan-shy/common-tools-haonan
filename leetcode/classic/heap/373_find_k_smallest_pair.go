package heap

import "container/heap"

type Pair struct {
	A, B int
}

func (p Pair) pairSum() int {
	return p.A + p.B
}

type PairHeap struct {
	Pairs []Pair
}

func (h PairHeap) Len() int {
	return len(h.Pairs)
}

func (h PairHeap) Less(i, j int) bool {
	return h.Pairs[i].pairSum() < h.Pairs[j].pairSum()
}

func (h PairHeap) Swap(i, j int) {
	h.Pairs[i], h.Pairs[j] = h.Pairs[j], h.Pairs[i]
}

func (h *PairHeap) Push(x any) {
	element := x.(Pair)
	h.Pairs = append(h.Pairs, element)
}

func (h *PairHeap) Pop() any {
	old := h.Pairs
	x := old[h.Len()-1]
	h.Pairs = h.Pairs[:h.Len()-1]
	return x
}

// 暴力解法超时
func kSmallestPairs_General(nums1 []int, nums2 []int, k int) [][]int {
	hp := &PairHeap{
		Pairs: make([]Pair, 0),
	}

	for _, element1 := range nums1 {
		for _, element2 := range nums2 {
			p := Pair{
				A: element1,
				B: element2,
			}

			if hp.Len() >= k && p.pairSum() > hp.Pairs[hp.Len()-1].pairSum() {
				continue
			}
			heap.Push(hp, p)

			if hp.Len() > k {
				hp.Pairs = hp.Pairs[:hp.Len()-1]
			}
		}
	}
	ans := make([][]int, 0, k)
	for ; k > 0; k-- {
		pair := heap.Pop(hp).(Pair)
		ans = append(ans, []int{pair.A, pair.B})
	}
	return ans
}

type PairIdxHeap struct {
	Pairs        []Pair
	Nums1, Nums2 []int
}

func (h PairIdxHeap) Len() int {
	return len(h.Pairs)
}

func (h PairIdxHeap) Less(i, j int) bool {
	a, b := h.Pairs[i], h.Pairs[j]
	return h.Nums1[a.A]+h.Nums2[a.B] < h.Nums1[b.A]+h.Nums2[b.B]
}

func (h PairIdxHeap) Swap(i, j int) {
	h.Pairs[i], h.Pairs[j] = h.Pairs[j], h.Pairs[i]
}

func (h *PairIdxHeap) Push(x any) {
	element := x.(Pair)
	h.Pairs = append(h.Pairs, element)
}

func (h *PairIdxHeap) Pop() any {
	old := h.Pairs
	x := old[h.Len()-1]
	h.Pairs = h.Pairs[:h.Len()-1]
	return x
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {

	pairs := make([]Pair, 0, min_kSmallestPairs(len(nums1), k))
	hq := &PairIdxHeap{
		Pairs: make([]Pair, 0),
		Nums1: nums1,
		Nums2: nums2,
	}

	for i := 0; i < k && i < len(nums1); i++ {
		pairs = append(pairs, Pair{
			A: i,
			B: 0,
		})
	}

	hq.Pairs = pairs

	ans := make([][]int, 0, k)
	for hq.Len() > 0 && k > 0 {
		idx := heap.Pop(hq).(Pair)
		ans = append(ans, []int{nums1[idx.A], nums2[idx.B]})

		if idx.B+1 < len(nums2) {
			heap.Push(hq, Pair{
				A: idx.A,
				B: idx.B + 1,
			})
		}
		k--
	}
	return ans
}

func min_kSmallestPairs(a, b int) int {
	if a < b {
		return a
	}
	return b
}
