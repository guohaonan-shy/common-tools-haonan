package every_day

func minSteps(n int) int {

	dp := make([]int, n+1)
	dp[1] = 0
	if n == 1 {
		return dp[1]
	}

	dp[2] = 2

	for i := 3; i <= n; i++ {
		dp[i] = dp[1] + i
		for j := 2; j < i/2; j++ {
			if i%j != 0 {
				continue
			}
			dp[i] = min(dp[i], dp[j]+i/j)
		}
	}

	return dp[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
