package dynammic_programming

func maxProfit_general(k int, prices []int) int {

	statusNum := k*2 + 1

	dp := make([][]int, len(prices)+1)
	for i := range dp {
		dp[i] = make([]int, statusNum)
	}

	for i := 1; i <= len(prices); i++ {
		for j := 1; j < statusNum && j <= i; j++ {
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
	for i := 0; i < statusNum; i++ {
		profits = max(profits, dp[len(prices)][i])
	}
	return profits
}
