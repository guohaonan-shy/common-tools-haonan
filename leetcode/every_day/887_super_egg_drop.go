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

/*
	Finally, let us make a conclusion about this solution:

	After we decided to use dynamic programming to solve this problem, we try to abstract the sub problem that is similar with the problem.
	First - analyze the problem and split into sub problem:
		We have n floors and k eggs, supposed we try to drop an egg to shrink the scope, we have two cases:
			1. if the egg is broken, it means that the threshold is in [0, t-1], and now we have k-1 eggs => sub problem: we have k-1 eggs to check under t-1 floor
			2. if not, it means that the threshold is in [t, n], and now we still have k eggs => sub problem: we have k eggs to check under n-t floor
	Second - optimize the time complexity:
		if we iterate by n and k, then try from 1 to current highest floor, the time complexity is O(n^2*k) => timeout
		we know:
			1. dp[n-t][eggs] is decreasing as t increases
			2. dp[t-1][eggs-1] is increasing as t increases
		when we try the drop floor from 1 to n, the maximum of dp[n-t][eggs] and dp[t-1][eggs-1] is the moves we can determine with certainty
		therefore, the position that are closer to the mid is possible to minimum value => binary search to find the minimum value
*/
