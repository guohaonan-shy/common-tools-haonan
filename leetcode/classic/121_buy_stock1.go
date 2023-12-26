package classic

import "math"

// 首先买入和卖出的日期不相同，其次整个过程只售卖一次
// 关键的解题点，对于每一天来说，如果第i天出售手中的股票，那么要使得手中股票利益最大一定是以第i天之前的最低价买入，因此寻找的是出售的日子-第i天，可以使利益大于其他日子出售的收益
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
