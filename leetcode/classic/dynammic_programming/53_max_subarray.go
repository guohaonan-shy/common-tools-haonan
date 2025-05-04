package dynammic_programming

func maxSubArray(nums []int) int {
	maxVal := 0
	dp := make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		pre := 0
		if i > 0 {
			pre = dp[i-1]
		}
		/*
			该题目的子问题是求解“以索引i为终点的子数组的最大和”，即dp[i]:
				dp[i] = max(dp[i-1]+nums[i], nums[i])
			dp[i-1]即以i-1为终点的子数组最大和 <= 0时，无论当前元素nums[i]是任何值，它的以i为终点(包括当前dp[i-1]所表示的子数组)的子数组和一定是小于仅包含nums[i]
			反之，dp[i-1] > 0时，dp[i]必须要包含dp[i-1]，否则值不是最大值
		*/
		dp[i] = max(pre+nums[i], nums[i])

		maxVal = max(maxVal, dp[i])
	}
	return maxVal
}
