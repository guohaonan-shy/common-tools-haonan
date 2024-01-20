package back_track

func exist(board [][]byte, word string) bool {
	var dfs func(i, j int, idx int) bool
	dfs = func(i, j int, idx int) bool {

		if idx == len(word) {
			return true
		}

		if i < 0 || j < 0 || i >= len(board) || j >= len(board[0]) || board[i][j] == '#' {
			return false
		}

		if board[i][j] != word[idx] {
			return false
		}

		temp := board[i][j]
		board[i][j] = '#'

		if dfs(i-1, j, idx+1) {
			return true
		}

		if dfs(i+1, j, idx+1) {
			return true
		}

		if dfs(i, j-1, idx+1) {
			return true
		}

		if dfs(i, j+1, idx+1) {
			return true
		}

		board[i][j] = temp
		return false
	}

	for i, row := range board {
		for j := range row {
			if dfs(i, j, 0) {
				return true
			}
		}
	}
	return false
}
