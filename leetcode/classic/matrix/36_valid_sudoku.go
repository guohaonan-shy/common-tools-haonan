package matrix

func isValidSudoku(board [][]byte) bool {

	rows := make([]bool, len(board))
	column := make([]bool, len(board))
	grid := make([][3]bool, 3)

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {

			if board[i][j] == '.' {
				continue
			}

			x, y := i/3, j/3
			if rows[i] && column[j] && grid[x][y] {
				continue
			}

			rowContent := map[byte]struct{}{}
			for _, v := range board[i] {
				if v == '.' {
					continue
				}
				if _, ok := rowContent[v]; ok {
					return false
				} else {
					rowContent[v] = struct{}{}
				}
			}
			rows[i] = true

			columnContent := map[byte]struct{}{}
			for a := 0; a < len(board); a++ {
				if board[a][j] == '.' {
					continue
				}
				if _, ok := columnContent[board[a][j]]; ok {
					return false
				} else {
					columnContent[board[a][j]] = struct{}{}
				}
			}
			column[j] = true

			gridContent := map[byte]struct{}{}
			for a := x * 3; a < (x+1)*3; a++ {
				for b := y * 3; b < (y+1)*3; b++ {
					if board[a][b] == '.' {
						continue
					}

					if _, ok := gridContent[board[a][b]]; ok {
						return false
					} else {
						gridContent[board[a][b]] = struct{}{}
					}
				}
			}
			grid[x][y] = true
		}
	}
	return true

}
