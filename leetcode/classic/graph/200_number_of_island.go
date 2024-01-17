package graph

type Pos struct {
	X, Y int
}

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
