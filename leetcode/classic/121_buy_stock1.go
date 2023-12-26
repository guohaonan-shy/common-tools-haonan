package classic

import "math"

func maxProfit(prices []int) int {
	minPrice := math.MaxInt32 // 关键定义，当卖出交易日为第i天时，第i天之前的最低买入价
	profit := 0
	for _, price := range prices {
		profit = max(profit, price-minPrice)
		minPrice = min(minPrice, price)
	}
	return profit
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
