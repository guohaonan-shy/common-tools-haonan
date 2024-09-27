package dynammic_programming

func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			if i == coin {
				dp[i] = 1
				continue
			}
			if i > coin && dp[i-coin] > 0 && (dp[i] == 0 || dp[i] > dp[i-coin]+1) {
				dp[i] = dp[i-coin] + 1
			}
		}
	}

	if dp[amount] == 0 {
		dp[amount] = -1
	}
	return dp[amount]
}
