package graph

func solve(board [][]byte) {
	rowN, columnN := len(board), len(board[0])

	var f func(row, col int)

	f = func(row, col int) {
		if row >= rowN || row < 0 || col >= columnN || col < 0 {
			return
		}

		if board[row][col] != 'O' {
			return
		}

		board[row][col] = 'A'

		f(row+1, col)
		f(row-1, col)
		f(row, col+1)
		f(row, col-1)

		return
	}

	for i := 0; i < rowN; i++ {
		f(i, 0)
		f(i, columnN-1)
	}

	for i := 0; i < columnN; i++ {
		f(0, i)
		f(rowN-1, i)
	}

	for rowIdx, row := range board {
		for colIdx := range row {
			if board[rowIdx][colIdx] == 'A' {
				board[rowIdx][colIdx] = 'O'
			} else if board[rowIdx][colIdx] == 'O' {
				board[rowIdx][colIdx] = 'X'
			}
		}
	}
}
