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

func solveV2(board [][]byte) {
	rowCnt, colCnt := len(board), len(board[0])

	var dfs func(i, j int)

	dfs = func(i, j int) {

		if i < 0 || i > rowCnt-1 || j < 0 || j > colCnt-1 {
			return
		}

		if board[i][j] == 'X' || board[i][j] == 'A' {
			return
		}

		board[i][j] = 'A'

		dfs(i-1, j)
		dfs(i+1, j)
		dfs(i, j-1)
		dfs(i, j+1)
		return
	}

	// start from the edge of the board to annotate 'O' that don't need to replace
	for i := 0; i < rowCnt; i++ {
		dfs(i, 0)
		dfs(i, colCnt-1)
	}

	for i := 0; i < colCnt; i++ {
		dfs(0, i)
		dfs(rowCnt-1, i)
	}

	for i, row := range board {
		for j := range row {

			if board[i][j] == 'A' {
				board[i][j] = 'O'
			} else if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
		}
	}
	return
}
