package matrix

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_36(t *testing.T) {
	t.Run("case1", func(t *testing.T) {
		case1 := [][]byte{
			{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
			{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
			{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
			{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
			{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
			{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
			{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
			{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
			{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
		}
		assert.Equal(t, true, isValidSudokuV2(case1))
	})

	t.Run("case2", func(t *testing.T) {
		case2 := [][]byte{
			{'.', '.', '.', '.', '5', '.', '.', '1', '.'},
			{'.', '4', '.', '3', '.', '.', '.', '.', '.'},
			{'.', '.', '.', '.', '.', '3', '.', '.', '1'},
			{'8', '.', '.', '.', '.', '.', '.', '2', '.'},
			{'.', '.', '2', '.', '7', '.', '.', '.', '.'},
			{'.', '1', '5', '.', '.', '.', '.', '.', '.'},
			{'.', '.', '.', '.', '.', '2', '.', '.', '.'},
			{'.', '2', '.', '9', '.', '.', '.', '.', '.'},
			{'.', '.', '4', '.', '.', '.', '.', '.', '.'},
		}
		assert.Equal(t, false, isValidSudokuV2(case2))
	})
}

func Test_54(t *testing.T) {
	t.Run("case1", func(t *testing.T) {
		case1 := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
		assert.Equal(t, []int{1, 2, 3, 6, 9, 8, 7, 4, 5}, spiralOrder(case1))
	})
}

func Test_79(t *testing.T) {
	type args struct {
		board [][]byte
		word  string
	}

	for _, testCase := range []struct {
		name string
		args args
		res  bool
	}{
		{
			name: "case 1",
			args: args{
				board: [][]byte{
					{'A', 'B', 'C', 'E'},
					{'S', 'F', 'C', 'S'},
					{'A', 'D', 'E', 'E'},
				},
				word: "ABCCED",
			},
			res: true,
		},
		{
			name: "case 2",
			args: args{
				board: [][]byte{
					{'A', 'B', 'C', 'E'},
					{'S', 'F', 'C', 'S'},
					{'A', 'D', 'E', 'E'},
				},
				word: "SEE",
			},
			res: true,
		},
		{
			name: "case 3",
			args: args{
				board: [][]byte{
					{'A', 'B', 'C', 'E'},
					{'S', 'F', 'C', 'S'},
					{'A', 'D', 'E', 'E'},
				},
				word: "ABCB",
			},
			res: false,
		},
		{
			name: "case 4",
			args: args{
				board: [][]byte{
					{'a', 'b'},
					{'c', 'd'},
				},
				word: "acdb",
			},
			res: true,
		},
		{
			name: "case 5",
			args: args{
				board: [][]byte{
					{'C', 'A', 'A'},
					{'A', 'A', 'A'},
					{'B', 'C', 'D'},
				},
				word: "AAB",
			},
			res: true,
		},
	} {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, testCase.res, existDFS(testCase.args.board, testCase.args.word))
		})
	}
}
