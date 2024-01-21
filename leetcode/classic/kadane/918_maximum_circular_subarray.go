package kadane

func maxSubarraySumCircular(nums []int) int {
	maxEndHere, minEndHere := nums[0], nums[0]
	globalMax, globalMin := nums[0], nums[0]
	totalSum := nums[0]

	for i := 1; i < len(nums); i++ {

		maxEndHere = max(maxEndHere+nums[i], nums[i])
		minEndHere = min(minEndHere+nums[i], nums[i])

		globalMax = max(globalMax, maxEndHere)
		globalMin = min(globalMin, minEndHere)

		totalSum += nums[i]
	}

	// 另外一种情况，最大子序列跨环
	// 如果global max 小于零，代表，列表元素无大于等于0的值，即此时全局最小为所有元素之和
	if globalMax < 0 {

	} else {
		globalMax = max(globalMax, totalSum-globalMin)
	}

	return globalMax
}

//func max(a, b int) int {
//	if a < b {
//		return b
//	}
//
//	return a
//}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
