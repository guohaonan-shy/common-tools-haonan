package heap

import (
	"container/heap"
	"sort"
)

type Ipo struct {
	Nums [][2]int
}

func (ipo Ipo) Len() int {
	return len(ipo.Nums)
}

func (ipo Ipo) Less(i, j int) bool {
	nums := ipo.Nums
	return nums[i][1] > nums[j][1]
}

func (ipo Ipo) Swap(i, j int) {
	ipo.Nums[i], ipo.Nums[j] = ipo.Nums[j], ipo.Nums[i]
}

func (ipo *Ipo) Push(x any) {
	nums := ipo.Nums
	element := x.([2]int)
	nums = append(nums, element)
	ipo.Nums = nums
}

func (ipo *Ipo) Pop() any {
	nums := ipo.Nums
	x := nums[len(nums)-1]
	ipo.Nums = nums[:len(nums)-1]
	return x
}

func findMaximizedCapital(k int, w int, profits []int, capital []int) int {

	list := make([][2]int, 0, len(profits))
	for i := 0; i < len(profits); i++ {
		list = append(list, [2]int{capital[i], profits[i]})
	}
	sort.Slice(list, func(i, j int) bool {
		return list[i][0] < list[j][0]
	})

	hp := &Ipo{
		Nums: make([][2]int, 0),
	}
	cur := 0
	for ; k > 0; k-- {
		for cur < len(list) && list[cur][0] <= w {
			heap.Push(hp, [2]int{list[cur][0], list[cur][1]})
			cur++
		}

		if hp.Len() == 0 {
			break
		}

		project := heap.Pop(hp)
		w += project.([2]int)[1]
	}
	return w
}
