package every_day

var directions = [][2]int{
	{-2, 1}, {-1, 2},
	{1, 2}, {2, 1},
	{2, -1}, {1, -2},
	{-1, -2}, {-2, -1},
}

type Pos struct {
	X, Y   int
	Step   int
	isStop bool
}

// 复杂度高，且oom了
func knightProbability(n int, k int, row int, column int) float64 {
	if k == 0 {
		return float64(1)
	}

	queue := make([]Pos, 0)
	queue = append(queue, Pos{
		X:      row,
		Y:      column,
		Step:   0,
		isStop: false,
	})

	inBoard := 0
	notInBoard := 0

	for len(queue) > 0 {
		pos := queue[0]
		queue = queue[1:]
		if pos.Step == k {
			if pos.X >= 0 && pos.X < n && pos.Y >= 0 && pos.Y < n && !pos.isStop { // 到位
				inBoard++
			} else {
				notInBoard++
			}
			continue
		}

		isStop := false
		if pos.X < 0 || pos.X >= n || pos.Y < 0 || pos.Y >= n {
			isStop = true
		}

		for _, dir := range directions {
			newX, newY := pos.X+dir[0], pos.Y+dir[1]
			queue = append(queue, Pos{
				X:      newX,
				Y:      newY,
				Step:   pos.Step + 1,
				isStop: pos.isStop || isStop,
			})

		}
	}

	//return float64(inBoard) / math.Pow(8, float64(k))
	return float64(inBoard) / float64(notInBoard+inBoard)
}
