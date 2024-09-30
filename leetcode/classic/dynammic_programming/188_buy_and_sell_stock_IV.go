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

func maxProfitV2_kTransactions(k int, prices []int) int {
	dp := make([][]int, len(prices)+1)
	for i := range dp {
		dp[i] = make([]int, 2*k+1)
	}

	for day := 1; day <= len(prices); day++ {
		for status := 0; status <= day && status < 2*k+1; status++ {

			if status == 0 {
				dp[day][0] = dp[day-1][0]
				continue
			}

			profit := prices[day-1]
			if status%2 == 1 {
				profit = -profit
			}

			if day == status {
				dp[day][status] = dp[day-1][status-1] + profit
			} else {
				dp[day][status] = max(dp[day-1][status], dp[day-1][status-1]+profit)
			}

		}
	}

	maxVal := 0
	for i := 0; i < 2*k+1; i++ {
		maxVal = max(maxVal, dp[len(prices)][i])
	}
	return maxVal
}
