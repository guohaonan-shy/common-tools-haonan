package classic

import "sort"

// assume the majority element always exists
// Could you solve the problem in linear time and in O(1) space?
// 1. hashmap, time complexity and space complexity is O(N)
func majorityElement(nums []int) int {
	mapping := make(map[int]int, len(nums))
	for _, value := range nums {
		if cnt, ok := mapping[value]; ok {
			mapping[value] = cnt + 1
		} else {
			mapping[value] = 1
		}
	}

	var res int
	for value, cnt := range mapping {
		if cnt > len(nums)/2 {
			res = value
			break
		}
	}
	return res
}

func majorityElement_sort(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]
}
