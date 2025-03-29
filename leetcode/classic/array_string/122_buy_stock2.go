package array_string

// 首先，本题可以求解的是总收益，即经过多次的股票交易的最大收益；但限定手上只能有一张股票
// 因此首先想到的思路是dynamic programming 或者贪心
func maxProfit_DP(prices []int) int {
	dp := make([][2]int, len(prices)) // dp[i][0]表示第i天结束后，手上没有股票时的最大收益；dp[i][1]表示第i天结束后，手上有股票时的最大收益
	dp[0][0], dp[0][1] = 0, -prices[0]

	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i]) // 如果第i天手上没有股票，两种情况:1.第i-1天没有股票且第i天没有买入；2.第i-1天手上有股票且第i天卖掉；即第i天的收益只与第i-1天的收益有关
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i]) // 同上
	}

	return max(dp[len(prices)-1][0], dp[len(prices)-1][1])
}

// 贪心 - 即多次交易的收益总和，多笔交易price[sell] - price[buy]，一个区间可被拆分成该区间多个子区间
// prices[sell] - price[buy] = price[sell]-price[sell-1] + price[sell-1]-price[sell-2] + ...... + price[buy+1]-price[buy]
// 则最大收益即为所有上升区间的总和；答案是最优解，但并非最优方案
func maxProfit_Greedy(prices []int) int {
	ans := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			ans += prices[i] - prices[i-1]
		}
	}
	return ans
}

func maxProfitDPV2(prices []int) int {
	dp := make([][2]int, len(prices)+1)

	for i := range dp {
		dp[i] = [2]int{}
	}

	dp[1][0], dp[1][1] = 0, -prices[0]

	for i := 2; i <= len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i-1])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i-1])
	}

	return dp[len(prices)][0] // 最后不持有股票一定比持有股票赚钱
}
