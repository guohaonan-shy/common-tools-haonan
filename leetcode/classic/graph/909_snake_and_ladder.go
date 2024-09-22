package graph

type Dest struct {
	DestVal         int
	Path            int
	IsSnakeOrLadder bool
}

func NewDest(dest, path int, isSnakeOrLadder bool) *Dest {
	return &Dest{
		DestVal:         dest,
		Path:            path,
		IsSnakeOrLadder: isSnakeOrLadder,
	}
}

func snakesAndLadders(board [][]int) int {

	queue := make([]*Dest, 0)
	queue = append(queue, NewDest(1, 0, false))
	reached := make(map[int]map[bool]int, 0)
	row := len(board)

	for len(queue) > 0 {
		pos := queue[0]
		queue = queue[1:]

		if val, ok := reached[pos.DestVal]; ok {
			if value, ok1 := val[pos.IsSnakeOrLadder]; ok1 {
				if value > pos.Path {
					reached[pos.DestVal][pos.IsSnakeOrLadder] = pos.Path
				} else {
					continue
				}
			} else {
				reached[pos.DestVal][pos.IsSnakeOrLadder] = pos.Path
			}
		} else {
			reached[pos.DestVal] = map[bool]int{
				pos.IsSnakeOrLadder: pos.Path,
			}
		}
		// cal pos in board
		div, remain := pos.DestVal/row, pos.DestVal%row
		x := row - 1 - div
		if remain == 0 {
			x += 1
			div -= 1
		}

		var offset int
		if div%2 == 1 {
			offset = (row - remain) % row
		} else {
			offset = (remain - 1 + row) % row
		}

		y := offset
		//
		if board[x][y] == -1 || (pos.IsSnakeOrLadder && board[x][y] != -1) {
			for i := pos.DestVal + 1; i <= min(pos.DestVal+6, row*row); i++ {
				queue = append(queue, NewDest(i, pos.Path+1, false))
			}
		} else {
			queue = append(queue, NewDest(board[x][y], pos.Path, true))
		}
	}

	if _, ok := reached[row*row]; !ok {
		return -1
	} else {
		if reached[row*row][true] != 0 && reached[row*row][false] != 0 {
			return min(reached[row*row][true], reached[row*row][false])
		} else {
			if reached[row*row][true] == 0 {
				return reached[row*row][false]
			}
			return reached[row*row][true]
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func snakesAndLadders_standard(board [][]int) int {
	queue := make([][]int, 0)
	queue = append(queue, []int{1, 0})
	row := len(board)
	reached := make(map[int]bool, 0)

	for len(queue) > 0 {

		cur := queue[0]
		queue = queue[1:]

		for i := cur[0] + 1; i <= min(cur[0]+6, row*row); i++ {

			// cal pos in board
			div, remain := i/row, i%row
			x := row - 1 - div
			if remain == 0 {
				x += 1
				div -= 1
			}

			var offset int
			if div%2 == 1 {
				offset = (row - remain) % row
			} else {
				offset = (remain - 1 + row) % row
			}

			y := offset

			next := i
			if board[x][y] != -1 {
				next = board[x][y]
			}

			if next == row*row {
				return cur[1] + 1
			}

			if !reached[next] {
				reached[next] = true
				queue = append(queue, []int{next, cur[1] + 1})
			}
		}
	}
	return -1
}

type BFSNode struct {
	label int
	step  int
}

func snakesAndLaddersV2(board [][]int) int {
	queue := make([]*BFSNode, 0)
	queue = append(queue, &BFSNode{
		label: 1,
		step:  0,
	})
	dimension := len(board)
	reachedMap := make(map[int]bool, 0)
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		label, step := cur.label, cur.step

		for i := label + 1; i < min(label+6, dimension*dimension)+1; i++ {
			rowIdx, colIdx := label2pos(i, dimension)
			next := i
			if board[rowIdx][colIdx] != -1 {
				// if the step has the snake or ladder, we walk further in this step; we can move further as fewer steps as possible
				next = board[rowIdx][colIdx]
			}

			if next == dimension*dimension {
				return step + 1
			}

			if !reachedMap[next] {
				reachedMap[next] = true
				queue = append(queue, &BFSNode{
					label: next,
					step:  step + 1,
				})
			}
		}
	}
	return -1
}

func label2pos(label int, dimension int) (int, int) {
	label--
	row, column := label/dimension, label%dimension

	rowIdx := dimension - 1 - row
	columnIdx := column
	if row%2 == 1 {
		columnIdx = dimension - 1 - column
	}
	return rowIdx, columnIdx
}
