package binary_search

import "math"

func findPeakElement(nums []int) int {
	return binaryFindPeak(nums, 0, len(nums)-1)
}

func binaryFindPeak(nums []int, left, right int) int {
	if len(nums) == 1 {
		return 0
	}

	// 注意一开始mid就是数组边界
	//if len(nums) == 2 {
	//	if nums[0] > nums[1] {
	//		return 0
	//	} else {
	//		return 1
	//	}
	//}

	mid := (left + right + 1) / 2

	var leftVal, rightVal int
	if mid == 0 {
		leftVal = math.MinInt64
	} else {
		leftVal = nums[mid-1]
	}

	if mid == len(nums)-1 {
		rightVal = math.MinInt64
	} else {
		rightVal = nums[mid+1]
	}

	if leftVal < nums[mid] && nums[mid] > rightVal {
		return mid
	}

	if leftVal > nums[mid] && nums[mid] > rightVal {
		return binaryFindPeak(nums, 0, mid-1)
	} else {
		return binaryFindPeak(nums, mid+1, right)
	}

}
