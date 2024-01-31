package array_string

import "sort"

func nextPermutation(nums []int) {

	if len(nums) == 1 {
		return
	}

	index := 0
	for i := len(nums) - 1; i >= 0; i-- {
		if i-1 >= 0 && nums[i] > nums[i-1] {
			index = i
			break
		}
	}

	//changed := prev
	if index == 0 {
		for l, r := 0, len(nums)-1; l < r; {
			nums[l], nums[r] = nums[r], nums[l]
			l++
			r--
		}
		return
	}
	change := 0
	for i := index; i < len(nums); i++ {
		if nums[i] > nums[index-1] {
			change = i
		} else {
			break
		}
	}

	nums[index-1], nums[change] = nums[change], nums[index-1]

	sort.Ints(nums[index:])
	return
}
