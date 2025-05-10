package matrix

func existDFS(board [][]byte, word string) bool {
	start := word[0]

	input := [][2]int{}

	for i := range board {
		for j := range board[i] {
			if board[i][j] == start {
				input = append(input, [2]int{i, j})
			}
		}
	}

	if len(input) == 0 {
		return false
	}

	var dfs func(cur [2]int, curIdx int) bool
	reached := make(map[[2]int]bool, 0)
	directions := [4][2]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	}
	dfs = func(cur [2]int, curIdx int) bool {

		x, y := cur[0], cur[1]

		if board[x][y] != word[curIdx] {
			return false
		}
		if curIdx == len(word)-1 {
			return true
		}

		reached[[2]int{x, y}] = true
		for _, direction := range directions {
			xstep, ystep := direction[0], direction[1]
			if x+xstep < 0 || x+xstep >= len(board) || y+ystep < 0 || y+ystep >= len(board[0]) {
				continue
			}

			if reached[[2]int{x + xstep, y + ystep}] {
				continue
			}

			if dfs([2]int{x + xstep, y + ystep}, curIdx+1) {
				return true
			}
		}
		reached[[2]int{x, y}] = false
		return false
	}

	for _, pos := range input {
		if dfs(pos, 0) {
			return true
		}
	}
	return false
}
