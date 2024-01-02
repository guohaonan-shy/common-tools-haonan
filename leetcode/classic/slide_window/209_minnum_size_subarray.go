package slide_window

import (
	"math"
)

func standard_minSubArrayLen(target int, nums []int) int {
	minSize, windowSum := math.MaxInt32, 0
	left, right := 0, 0
	for right < len(nums) {
		windowSum += nums[right]
		for windowSum >= target { // 题目限制，目标值是大于等于1的，即left在这个循环条件最大满足right+1，此时windowSum是0
			minSize = min(minSize, right-left+1)
			windowSum -= nums[left]
			left++
		}
		right++
	}
	if minSize != math.MaxInt32 {
		return minSize
	}
	return 0
}

func minSubArrayLen(target int, nums []int) int {
	minSize, windowSum := math.MaxInt32, nums[0]
	left, right := 0, 0
	for left <= right && left < len(nums) && right < len(nums) {
		if windowSum < target {
			if right+1 < len(nums) {
				right++
				windowSum += nums[right]
			} else {
				break
			}
		} else {
			minSize = min(minSize, right-left+1)
			windowSum -= nums[left]
			left++
		}
	}
	if minSize != math.MaxInt32 {
		return minSize
	}
	return 0
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
