package matrix

var (
	Order = [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
)

func spiralOrder(matrix [][]int) []int {

	checked := make([][]bool, len(matrix))
	for i := range checked {
		checked[i] = make([]bool, len(matrix[0]))
	}
	res := make([]int, 0)
	x, y := 0, 0
	cur, cnt := 0, 0
	for {
		if cnt == 4 {
			break
		}

		if !checked[x][y] {
			res = append(res, matrix[x][y])
			checked[x][y] = true
		}

		// direction
		if x+Order[cur][0] < 0 || x+Order[cur][0] >= len(matrix) || y+Order[cur][1] < 0 || y+Order[cur][1] >= len(matrix[0]) || checked[x+Order[cur][0]][y+Order[cur][1]] {
			cur = (cur + 1) % 4
			cnt += 1
		} else {
			x += Order[cur][0]
			y += Order[cur][1]
			cnt = 0
		}
	}
	return res
}
