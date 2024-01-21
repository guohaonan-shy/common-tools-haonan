package kadane

func maxSubArray(nums []int) int {
	res := nums[0]
	maxEndHere := nums[0]

	for i := 1; i < len(nums); i++ {
		maxEndHere = max(maxEndHere+nums[i], nums[i])
		res = max(res, maxEndHere)
	}
	return res
}

func max(a, b int) int {
	if a < b {
		return b
	}

	return a
}
