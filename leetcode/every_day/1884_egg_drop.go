package every_day

import "math"

func twoEggDrop(n int) int {
	dp := make([]int, n+1)

	for i := 1; i < len(dp); i++ {
		dp[i] = math.MaxInt32
		for k := 1; k <= i; k++ {
			// two cases:
			// 1. on the kth floor, the egg is broken => f is in [1, k-1], we need k-1 times to check
			// 2. the egg is not broken, f is in [k, i] => do[i-k]

			dp[i] = min(dp[i], max(k-1, dp[i-k])+1)
		}
	}
	return dp[n]
}

func twoEggDropV2(n int) int {
	local := 1
	for ; n > local; local++ {
		n -= local
	}
	return local
}
