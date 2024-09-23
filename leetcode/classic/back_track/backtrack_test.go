package back_track

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_17(t *testing.T) {
	t.Run("case1", func(t *testing.T) {
		assert.Equal(t, []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}, letterCombinations("23"))
	})
}

func Test_22(t *testing.T) {
	t.Run("case", func(t *testing.T) {
		assert.Equal(t, []string{"((()))", "(()())", "(())()", "()(())", "()()()"}, generateParenthesis(3))
	})
}

func Test_77(t *testing.T) {
	t.Run("case1", func(t *testing.T) {
		assert.Equal(t, [][]int{{1, 2}, {1, 3}, {1, 4}, {2, 3}, {2, 4}, {3, 4}}, combine(4, 2))
		assert.Equal(t, [][]int{{1, 2}, {1, 3}, {1, 4}, {2, 3}, {2, 4}, {3, 4}}, combine_standard(4, 2))
	})

	t.Run("case2", func(t *testing.T) {
		assert.Equal(t, [][]int{{1}}, combine(1, 1))
	})

	t.Run("case3", func(t *testing.T) {
		assert.Equal(t, [][]int{{1, 2, 3}, {1, 2, 4}, {1, 3, 4}, {2, 3, 4}}, combine(4, 3))
	})
}

func Test_46(t *testing.T) {
	t.Run("case1", func(t *testing.T) {
		assert.Equal(t, [][]int{{0, 1}, {1, 0}}, permuteV2([]int{0, 1}))
	})
}

func Test_47(t *testing.T) {
	t.Run("permutation2", func(t *testing.T) {
		assert.Equal(t, [][]int{{1, 1, 2}, {1, 2, 1}, {2, 1, 1}}, permuteUniqueV2([]int{1, 1, 2}))
	})
}

func Test_51(t *testing.T) {
	t.Run("n_queen", func(t *testing.T) {
		assert.Equal(t, [][]string{{".Q..", "...Q", "Q...", "..Q."}, {"..Q.", "Q...", "...Q", ".Q.."}}, solveNQueens(4))
	})
}

func Test_52(t *testing.T) {
	t.Run("case1", func(t *testing.T) {
		assert.Equal(t, 2, totalNQueens(4))
		assert.Equal(t, 1, totalNQueens(1))
	})
}

func Test_38(t *testing.T) {
	t.Run("count and say", func(t *testing.T) {
		assert.Equal(t, "1211", countAndSay(4))
	})
}

func Test_60(t *testing.T) {
	t.Run("get kth permutation", func(t *testing.T) {
		assert.Equal(t, "213", getPermutation(3, 3))
		assert.Equal(t, "213", getPermutation_standard(3, 3))
	})

	t.Run("get kth permutation", func(t *testing.T) {
		assert.Equal(t, "2314", getPermutation(4, 9))
		assert.Equal(t, "2314", getPermutation_standard(4, 9))
	})

	t.Run("get kth permutation", func(t *testing.T) {
		assert.Equal(t, "123", getPermutation(3, 1))
		assert.Equal(t, "123", getPermutation_standard(3, 1))
	})
}

func Test_93(t *testing.T) {
	t.Run("restore ip", func(t *testing.T) {
		//_ = restoreIpAddresses("25525511135")
		_ = restoreIpAddresses("101023")
	})
}
