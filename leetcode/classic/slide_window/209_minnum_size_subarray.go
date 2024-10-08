package slide_window

import (
	"math"
)

func standard_minSubArrayLen(target int, nums []int) int {
	minSize, windowSum := len(nums), 0
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
	if left == 0 {
		return 0
	}
	return minSize
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

/*
pay more attention to the movement of pointer, we have to guarantee the index element is valid.
*/
func minSubArrayLenV2(target int, nums []int) int {
	minRes := len(nums) + 1
	window := nums[0]
	left, right := 0, 0 // 'left' is left side of the window, and 'right' is the right side of the window

	for left <= right && right < len(nums) {
		if window < target {
			if right+1 < len(nums) {
				right++
				window += nums[right] // keep right+1 in the right scope
			} else {
				// even though we scale down the left side of window, the sum of elements in window will be still lower than target
				break
			}
		} else {
			minRes = min(minRes, right-left+1)
			window -= nums[left]
			left++
		}
	}
	if minRes == len(nums)+1 {
		minRes = 0
	}
	return minRes
}

/*
after sorting out the loop logic, we have to control one varation and then iterate based on the current value => no need to consider corner casr
*/
func minSubArrayV2_Refractor(target int, nums []int) int {
	window := 0
	left, right := 0, 0
	minLength := len(nums)

	for ; right < len(nums); right++ {
		window += nums[right]
		for ; left <= right && window >= target; left++ {
			window -= nums[left]
			minLength = min(minLength, right-left+1)
		}
	}

	if left == 0 {
		minLength = 0
	}
	return minLength
}
