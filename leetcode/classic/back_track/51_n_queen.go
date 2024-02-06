package back_track

func solveNQueens(n int) [][]string {
	temp := make([][]byte, n)
	for i := range temp {
		temp[i] = make([]byte, n)
	}
	res := make([][]string, 0)
	dirs := [][]int{{-1, -1}, {-1, 1}, {1, -1}, {1, 1}}

	var dfs func(idx int)

	dfs = func(idx int) {
		if idx == n {
			solution := make([]string, n)
			for i := range temp {
				newItem := make([]byte, n)
				copy(newItem, temp[i])
				solution[i] = string(newItem)
			}
			res = append(res, solution)
			return
		}
		defaultStr := make([]byte, n)
		for i := range defaultStr {
			defaultStr[i] = '.'
		}
		temp[idx] = defaultStr
		for i := 0; i < n; i++ {
			// row
			// column
			columnRes := false
			for row := 0; row < idx; row++ {
				if temp[row][i] == 'Q' {
					columnRes = true
					break
				}
			}

			if columnRes {
				continue
			}
			//
			arrowRes := false
			for _, dir := range dirs {
				dx, dy := dir[0], dir[1]
				x, y := idx, i
				x += dx
				y += dy
				for x >= 0 && y >= 0 && x < n && y < n {
					if temp[x][y] == 'Q' {
						arrowRes = true
						break
					}
					x += dx
					y += dy
				}
				if arrowRes {
					break
				}
			}
			if arrowRes {
				continue
			}

			temp[idx][i] = 'Q'
			dfs(idx + 1)
			temp[idx][i] = '.'
		}
	}
	dfs(0)
	return res
}
