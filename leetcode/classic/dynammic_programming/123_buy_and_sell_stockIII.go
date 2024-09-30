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

// noticed: this practice limits us to have only two transactions, therefore, we need to maintain the status
// But 122 limit us to hold only one stock at the same time but we can trade for unlimited times.
func maxProfitV2(prices []int) int {
	/*
		this practice requires at most two transactions, so we will have five status:
			1. no any transaction (0)
			2. first buy-in (1)
			3. first sell-out (2)
			4. second buy-in (3)
			5. second sell-out (4)
		then, we just figure out the max profits in the last day.
	*/
	dp := make([][5]int, len(prices)+1)

	for day := 1; day <= len(prices); day++ {
		for i := 0; i <= day && i <= 4; i++ {
			// for the day i, like day 1, we can only have two status: no any transaction (0) or first buy-in (1)

			// for status 0, we can only acquire status 0 from previous day' status 0
			if i == 0 {
				dp[day][0] = dp[day-1][0]
				continue
			}

			profit := prices[day-1]
			if i%2 == 1 {
				profit = -profit
			}

			if day == i {
				// corner case: for day i's status i, like 2nd day, we can only reach the status 2, even if we make a transaction every day before.
				dp[day][i] = dp[day-1][i-1] + profit
			} else {
				dp[day][i] = max(dp[day-1][i], dp[day-1][i-1]+profit)
			}

		}
	}

	maxVal := 0
	for i := 0; i < 5; i++ {
		maxVal = max(maxVal, dp[len(prices)][i])
	}
	return maxVal
}
