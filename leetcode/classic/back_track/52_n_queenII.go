package back_track

var directions = [][]int{{1, -1}, {1, 1}, {-1, 1}, {-1, -1}}

func totalNQueens(n int) int {
	board := make([][]int, n)
	for i := range board {
		board[i] = make([]int, n)
	}
	cnt := 0

	var dfs func(int)
	dfs = func(i int) {
		if i == n {
			cnt++
			return
		}

		for j := 0; j < n; j++ {
			valid := true

			for row := 0; row < n; row++ {
				if board[row][j] == 1 {
					valid = false
					break
				}
			}

			if !valid {
				continue
			}

			for _, direction := range directions {
				tempX, tempY := i, j
				for tempX >= 0 && tempY >= 0 && tempX < n && tempY < n {
					if board[tempX][tempY] == 1 {
						valid = false
						break
					}
					tempX, tempY = tempX+direction[0], tempY+direction[1]
				}
				if !valid {
					break
				}
			}

			if !valid {
				continue
			}

			board[i][j] = 1
			dfs(i + 1)
			board[i][j] = 0
		}
	}

	dfs(0)

	return cnt
}
