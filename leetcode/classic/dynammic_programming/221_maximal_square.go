package dynammic_programming

import "strconv"

func maximalSquare(matrix [][]byte) int {
	isAllZero := true

	dp := make([][]int, len(matrix))
	for i := range matrix {
		dp[i] = make([]int, len(matrix[0]))
		for j := range matrix[i] {
			if matrix[i][j] == '1' {
				isAllZero = false
			}
			dp[i][j], _ = strconv.Atoi(string(matrix[i][j]))
		}
	}

	if isAllZero {
		return 0
	}

	globalMaxSquare := 1
	for i := 1; i < min(len(matrix), len(matrix[0])); i++ {
		square := i * i
		isChange := false
		for row := i; row < len(matrix); row++ {
			for column := i; column < len(matrix[0]); column++ {
				if dp[row-1][column] >= square && dp[row][column-1] >= square && dp[row-1][column-1] >= square && matrix[row][column] == '1' {
					dp[row][column] = (i + 1) * (i + 1)
					globalMaxSquare = max(globalMaxSquare, dp[row][column])
					isChange = true
				}
			}
		}
		if !isChange {
			break
		}

	}
	return globalMaxSquare
}

func maximalSquareV2(matrix [][]byte) int {
	dp := make([][]int, len(matrix))
	for i := range dp {
		dp[i] = make([]int, len(matrix[0]))
	}

	globalMaxVal := 0
	for row := 0; row < len(matrix); row++ {
		for column := 0; column < len(matrix[0]); column++ {
			if matrix[row][column] == '0' {
				dp[row][column] = 0
				continue
			}

			if column-1 < 0 || row-1 < 0 {
				// we only need to consider left and upper;
				// and no need on right and down: eg. column+1 >= len(matrix[0]) || row+1 >= len(matrix)
				dp[row][column] = 1
			} else if matrix[row-1][column] == '0' || matrix[row][column-1] == '0' || matrix[row-1][column-1] == '0' {
				dp[row][column] = 1
			} else {
				minWidth := min(dp[row-1][column-1], min(dp[row-1][column], dp[row][column-1]))
				dp[row][column] = minWidth + 1
			}

			globalMaxVal = max(globalMaxVal, dp[row][column]*dp[row][column])

		}
	}
	return globalMaxVal
}
