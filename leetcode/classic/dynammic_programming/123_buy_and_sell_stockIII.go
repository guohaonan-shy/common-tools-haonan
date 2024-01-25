package dynammic_programming

func maxProfit(prices []int) int {

	dp := make([][5]int, len(prices)+1)

	for i := 1; i <= len(prices); i++ {
		for j := 1; j < 5 && j <= i; j++ {
			// sell stock or buy stock
			money := dp[i-1][j-1]
			if j%2 == 1 { // buy
				money -= prices[i-1]
			} else {
				money += prices[i-1]
			}

			if j == i { // 这块是个特殊case，dp[i][i] 即第i天的第i个状态只能由前一天的上一个状态转移
				dp[i][j] = money
			} else {
				dp[i][j] = max(dp[i-1][j], money)
			}
		}
	}

	profits := 0
	for i := 0; i < 5; i++ {
		profits = max(profits, dp[len(prices)][i])
	}
	return profits
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
