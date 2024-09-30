package dynammic_programming

func minPathSum(grid [][]int) int {
	dp := make([][]int, len(grid))
	for i := range dp {
		dp[i] = make([]int, len(grid[0]))
		for j := range dp[i] {
			dp[i][j] = 0
		}
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if i == 0 && j == 0 {
				dp[0][0] = grid[0][0]
				continue
			}

			if i == 0 {
				dp[i][j] = dp[i][j-1] + grid[i][j]
				continue
			}

			if j == 0 {
				dp[i][j] = dp[i-1][j] + grid[i][j]
				continue
			}

			minVal := min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
			dp[i][j] = minVal
		}
	}
	return dp[len(grid)-1][len(grid[0])-1]
}
