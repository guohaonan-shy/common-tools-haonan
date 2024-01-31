package every_day

import (
	"math"
	"sort"
)

// 排序+双指针
func threeSumClosest(nums []int, target int) int {

	sort.Ints(nums)
	minVal := math.MaxInt32
	for i := 0; i < len(nums)-2; i++ {
		left, right := i+1, len(nums)-1

		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if abs(sum-target) < abs(minVal-target) {
				minVal = sum
			}
			if sum > target {
				right--
			} else if nums[i]+nums[left]+nums[right] < target {
				left++
			} else {
				return target
			}
		}
	}

	return minVal
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func threeSum(nums []int) int {
	total := 0
	for _, val := range nums {
		total += val
	}
	return total
}
