package dynammic_programming

import "math"

/*
space complexity is 0(N^2)
*/
func minimumTotal(triangle [][]int) int {
	dp := make([][]int, len(triangle))
	for i := range dp {
		dp[i] = make([]int, len(triangle[len(triangle)-1]))
		for j := range dp[i] {
			dp[i][j] = math.MaxInt32
		}
	}

	dp[0][0] = triangle[0][0]
	for row := 1; row < len(triangle); row++ {
		for column := 0; column < len(triangle[row]); column++ {
			if column == 0 {
				dp[row][column] = min(dp[row][column], dp[row-1][column]+triangle[row][column])
				continue
			}

			if column == len(triangle[row])-1 {
				dp[row][column] = min(dp[row][column], dp[row-1][column-1]+triangle[row][column])
				continue
			}

			minVal := min(dp[row-1][column], dp[row-1][column-1]) + triangle[row][column]
			dp[row][column] = min(dp[row][column], minVal)
		}
	}

	globalMin := math.MaxInt32
	for _, val := range dp[len(triangle)-1] {
		globalMin = min(globalMin, val)
	}
	return globalMin
}

func minimumTotalWithLessMemory(triangle [][]int) int {
	if len(triangle) == 1 {
		return triangle[0][0]
	}
	pre := make([]int, 0)
	pre = append(pre, triangle[0][0])

	for row := 1; row < len(triangle); row++ {
		cur := make([]int, len(triangle[row]))
		for column := 0; column < len(triangle[row]); column++ {
			if column == 0 {
				cur[0] = pre[0] + triangle[row][column]
				continue
			}
			if column == len(triangle[row])-1 {
				cur[column] = pre[len(pre)-1] + triangle[row][column]
				continue
			}
			minVal := min(pre[column], pre[column-1]) + triangle[row][column]
			cur[column] = minVal
		}

		pre = cur
	}
	globalMin := math.MaxInt32
	for _, val := range pre {
		globalMin = min(globalMin, val)
	}
	return globalMin
}
