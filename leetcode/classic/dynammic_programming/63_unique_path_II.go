package dynammic_programming

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	dp := make([][]int, len(obstacleGrid))

	for i := range dp {
		dp[i] = make([]int, len(obstacleGrid[0]))
	}

	if obstacleGrid[0][0] != 1 { // corner case: start point is obstacle
		dp[0][0] = 1
	}

	for row := 0; row < len(obstacleGrid); row++ {
		for column := 0; column < len(obstacleGrid[row]); column++ {
			if row == 0 && column == 0 {
				continue
			}

			if obstacleGrid[row][column] == 1 {
				continue
			}

			if row == 0 {
				dp[row][column] = dp[row][column-1]
				continue
			}

			if column == 0 {
				dp[row][column] = dp[row-1][column]
				continue
			}

			dp[row][column] = dp[row][column-1] + dp[row-1][column]
		}
	}
	return dp[len(obstacleGrid)-1][len(obstacleGrid[0])-1]
}
