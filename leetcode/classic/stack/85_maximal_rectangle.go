package stack

func maximalRectangle(matrix [][]byte) int {
	length := make([][]int, len(matrix))
	for i := range length {
		length[i] = make([]int, len(matrix[0]))
	}

	for i, row := range matrix {
		for j, v := range row {
			if v == '0' {
				continue
			}

			if j == 0 {
				length[i][j] = 1
			} else {
				length[i][j] = length[i][j-1] + 1
			}
		}
	}

	maxVal := -1
	for column := 0; column < len(matrix[0]); column++ {
		up, down := make([]int, len(matrix)), make([]int, len(matrix))
		stack := make([]int, 0)

		for i := 0; i < len(matrix); i++ {
			for len(stack) > 0 && length[stack[len(stack)-1]][column] >= length[i][column] {
				stack = stack[:len(stack)-1]
			}

			if len(stack) == 0 {
				up[i] = -1
			} else {
				up[i] = stack[len(stack)-1]
			}
			stack = append(stack, i)
		}

		stack = []int{}

		for i := len(matrix) - 1; i >= 0; i-- {
			for len(stack) > 0 && length[stack[len(stack)-1]][column] >= length[i][column] {
				stack = stack[:len(stack)-1]
			}

			if len(stack) == 0 {
				down[i] = len(matrix)
			} else {
				down[i] = stack[len(stack)-1]
			}
			stack = append(stack, i)
		}

		for i, leng := range length {
			maxVal = max(maxVal, (down[i]-up[i]-1)*leng[column])
		}

	}
	return maxVal
}
