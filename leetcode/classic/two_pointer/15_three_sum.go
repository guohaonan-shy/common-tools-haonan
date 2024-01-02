package two_pointer

import "sort"

// 三层循环的时间复杂度为O(N^3)，排序+双指针的复杂度则为O(N^2 + NlogN)，即为N^2
func threeSum(nums []int) [][]int {
	sort.Sort(sort.IntSlice(nums))
	res := make([][]int, 0)
	for i := 0; i < len(nums)-2; i++ {
		target := -nums[i]
		left, right := i+1, len(nums)-1
		for left < right {
			if nums[left]+nums[right] > target {
				for right-1 >= 0 && nums[right-1] == nums[right] {
					right--
				}
				right--
			} else if nums[left]+nums[right] < target {
				for left+1 < len(nums) && nums[left+1] == nums[left] {
					left++
				}
				left++
			} else {
				res = append(res, []int{nums[i], nums[left], nums[right]})
				for left+1 < len(nums) && nums[left+1] == nums[left] {
					left++
				}
				left++
			}
		}
		for i+1 < len(nums)-2 && nums[i] == nums[i+1] {
			i++
		}
	}
	return res
}
