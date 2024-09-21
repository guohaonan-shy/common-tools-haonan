package graph

type Pos struct {
	X, Y int
}

// 使用bfs
func numIslands(grid [][]byte) int {

	reached := make([][]bool, len(grid))
	rowNum, colNum := len(grid), len(grid[0])

	for i := range grid {
		reached[i] = make([]bool, len(grid[0]))
	}

	f := func(row, column int) {
		queue := make([]*Pos, 0)
		queue = append(queue, &Pos{X: row, Y: column})
		for len(queue) > 0 {
			p := queue[0]
			queue = queue[1:]

			if reached[p.X][p.Y] {
				continue
			}

			if p.X-1 >= 0 && !reached[p.X-1][p.Y] && grid[p.X-1][p.Y] == '1' {
				queue = append(queue, &Pos{X: p.X - 1, Y: p.Y})
			}

			if p.X+1 < rowNum && !reached[p.X+1][p.Y] && grid[p.X+1][p.Y] == '1' {
				queue = append(queue, &Pos{X: p.X + 1, Y: p.Y})
			}

			if p.Y-1 >= 0 && !reached[p.X][p.Y-1] && grid[p.X][p.Y-1] == '1' {
				queue = append(queue, &Pos{X: p.X, Y: p.Y - 1})
			}

			if p.Y+1 < colNum && !reached[p.X][p.Y+1] && grid[p.X][p.Y+1] == '1' {
				queue = append(queue, &Pos{X: p.X, Y: p.Y + 1})
			}
			reached[p.X][p.Y] = true
		}
	}

	ans := 0
	for rowIdx, row := range grid {
		for columnIdx, val := range row {

			if val == '0' || reached[rowIdx][columnIdx] {
				continue
			}

			f(rowIdx, columnIdx)
			ans++

		}
	}
	return ans
}

// 使用dfs
func numIslands_DFS(grid [][]byte) int {

	rowN, colN := len(grid), len(grid[0])
	//reached := make([][]int, len(grid))
	//for i := range reached {
	//	reached[i] = make([]int, len(grid[0]))
	//}

	var dfs func(rowIdx int, colIdx int)
	dfs = func(rowIdx int, colIdx int) {

		//if reached[rowIdx][colIdx] != 0 {
		//	return
		//}

		//reached[rowIdx][colIdx] = 1

		//if rowIdx-1 >= 0 && grid[rowIdx-1][colIdx] == '1' && reached[rowIdx-1][colIdx] == 0 {
		//	dfs(rowIdx-1, colIdx)
		//}
		//
		//if rowIdx+1 < rowN && grid[rowIdx+1][colIdx] == '1' && reached[rowIdx+1][colIdx] == 0 {
		//	dfs(rowIdx+1, colIdx)
		//}
		//
		//if colIdx-1 >= 0 && grid[rowIdx][colIdx-1] == '1' && reached[rowIdx][colIdx-1] == 0 {
		//	dfs(rowIdx, colIdx-1)
		//}
		//
		//if colIdx+1 < colN && grid[rowIdx][colIdx+1] == '1' && reached[rowIdx][colIdx+1] == 0 {
		//	dfs(rowIdx, colIdx+1)
		//}

		if rowIdx < 0 || rowIdx >= rowN || colIdx < 0 || colIdx >= colN || grid[rowIdx][colIdx] != '1' {
			return
		}
		grid[rowIdx][colIdx] = '0'
		dfs(rowIdx-1, colIdx)
		dfs(rowIdx+1, colIdx)
		dfs(rowIdx, colIdx-1)
		dfs(rowIdx, colIdx+1)
		//reached[rowIdx][colIdx] = 2
	}

	ans := 0
	for rowIdx, row := range grid {
		for colIdx, _ := range row {
			if grid[rowIdx][colIdx] == '1' {
				dfs(rowIdx, colIdx)
				ans++
			}
		}
	}
	return ans
}

func numIslandsV2(grid [][]byte) int {
	row, column := len(grid), len(grid[0])
	status := make([][]bool, row)
	for i := range status {
		status[i] = make([]bool, column)
	}

	cnt := 0
	var dfs func(i, j int)
	dfs = func(i, j int) {

		if i < 0 || i > row-1 || j < 0 || j > column-1 {
			return
		}

		if status[i][j] {
			return
		}

		if grid[i][j] == '0' {
			return
		}

		status[i][j] = true

		dfs(i-1, j)
		dfs(i+1, j)
		dfs(i, j-1)
		dfs(i, j+1)
		return
	}

	for i, rows := range grid {
		for j := range rows {
			if grid[i][j] == '0' {
				continue
			}

			if status[i][j] {
				continue
			}

			dfs(i, j)
			cnt++
		}
	}
	return cnt
}
