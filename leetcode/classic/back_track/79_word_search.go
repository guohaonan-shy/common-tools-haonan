package back_track

var dirs = [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func exist(board [][]byte, word string) bool {
	var dfs func(i, j int, idx int) bool
	dfs = func(i, j int, idx int) bool {

		if idx == len(word) { // match finished
			return true
		}

		// board[i][j] == '# means this element has been covered in this branch
		if i < 0 || j < 0 || i >= len(board) || j >= len(board[0]) || board[i][j] == '#' {
			return false
		}

		if board[i][j] != word[idx] {
			return false
		}

		temp := board[i][j]
		board[i][j] = '#'

		for _, dir := range dirs {
			dx, dy := i+dir[0], j+dir[1]
			if dfs(dx, dy, idx+1) {
				return true
			}
		}
		// four direction is not match, it means that the word[idx+1:] doesn't match with target
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
