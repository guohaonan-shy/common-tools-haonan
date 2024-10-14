package every_day

import "math"

/*
timeout => time complexity: O(n^2*k)
*/
func superEggDrop(k int, n int) int {
	dp := make([][]int, n+1) // the first dimension is the number of floors

	for i := range dp {
		dp[i] = make([]int, k+1) // the second dimension is the number of remaining eggs
	}

	for floor := 1; floor <= n; floor++ {

		for eggs := 1; eggs <= k; eggs++ {
			if eggs == 1 {
				dp[floor][1] = floor
				continue
			}

			//for try := 1; try <= floor; try++ {
			//	/*
			//		supposed we have a n-level building, and we drop the egg from ith floor:
			//		1. if broken, we can use k-1 eggs to identify from 1~i-1
			//		2. if not, we can use k eggs to check from i ~ n
			//	*/
			//	if dp[floor][eggs] == 0 {
			//		dp[floor][eggs] = math.MaxInt32
			//	}
			//
			//	dp[floor][eggs] = min(dp[floor][eggs], max(dp[try-1][eggs-1], dp[floor-try][eggs])+1)
			//}

			// we can use binary search to find the minimum of maximum among dp[try-1][eggs-1](increasing as try increases) and dp[floor-try][eggs] (decreasing as try increases)
			left, right := 1, floor
			for left < right {
				mid := (left + right) / 2
				if dp[mid-1][eggs-1] < dp[floor-mid][eggs] {
					left = mid + 1
				} else {
					right = mid
				}
			}
			if dp[floor][eggs] == 0 {
				dp[floor][eggs] = math.MaxInt32
			}
			dp[floor][eggs] = min(dp[floor][eggs], max(dp[left-1][eggs-1], dp[floor-left][eggs])+1)
		}
	}
	return dp[n][k]
}
