package every_day

import "math"

/*
	we need to consider one important case, local min value might not on the path of the least path
*/

/*
backtrack will be time-consuming => timeout, both time complexity and space complexity are O(n^m)
*/
func minimumPathInGrid(grid [][]int, moveCost [][]int) int {
	globalVal := math.MaxInt32
	for i := 0; i < len(grid[0]); i++ {
		globalVal = min(globalVal, minPathProcessor(1, i, grid, moveCost, 0+grid[0][i]))
	}
	return globalVal
}

func minPathProcessor(nextRow, column int, grid [][]int, moveCost [][]int, pathSum int) int {
	if nextRow == len(grid) {
		return pathSum
	}
	cur := grid[nextRow-1][column]
	edges := moveCost[cur]

	globalVal := math.MaxInt32
	for i := 0; i < len(grid[0]); i++ {
		globalVal = min(globalVal, minPathProcessor(nextRow+1, i, grid, moveCost, pathSum+grid[nextRow][i]+edges[i]))
	}
	return globalVal
}

// time complexity: O(m*n^2)
func minimumPathInGridDP(grid [][]int, moveCost [][]int) int {
	dp := make([][]int, len(grid))
	for i := range dp {
		dp[i] = make([]int, len(grid[0]))
		for j := range dp[i] {
			dp[i][j] = math.MaxInt32
		}
	}
	for row := 0; row < len(grid); row++ {
		for column := 0; column < len(grid[0]); column++ {
			if row == 0 {
				dp[0][column] = grid[0][column]
				continue
			}

			for previousCol := 0; previousCol < len(grid[0]); previousCol++ {
				edges := moveCost[grid[row-1][previousCol]]
				dp[row][column] = min(dp[row][column], dp[row-1][previousCol]+edges[column]+grid[row][column])
			}
		}
	}

	globalVal := math.MaxInt32
	for i := range dp[len(grid)-1] {
		globalVal = min(globalVal, dp[len(grid)-1][i])
	}

	return globalVal
}
