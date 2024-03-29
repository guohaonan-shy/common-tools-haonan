package dynammic_programming

import (
	"github.com/common-tools-haonan/leetcode/classic/graph"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_44(t *testing.T) {
	t.Run("wildcard matching", func(t *testing.T) {
		assert.Equal(t, true, isMatch("adceb", "*a*b"))
	})
}

func Test_97(t *testing.T) {
	t.Run("dp97", func(t *testing.T) {
		assert.Equal(t, true, isInterleave("aabcc", "dbbca", "aadbbcbcac"))
	})

	t.Run("dp97", func(t *testing.T) {
		assert.Equal(t, false, isInterleave("aabcc", "dbbca", "aadbbbaccc"))
	})
}

func Test_123(t *testing.T) {
	t.Run("buy_stock_3", func(t *testing.T) {
		assert.Equal(t, 4, maxProfit([]int{1, 2, 3, 4, 5}))
	})
}

func Test_188(t *testing.T) {
	t.Run("general_buy_stock", func(t *testing.T) {
		assert.Equal(t, 2, maxProfit_general(2, []int{2, 4, 1}))
	})

	t.Run("general_buy_stock", func(t *testing.T) {
		assert.Equal(t, 7, maxProfit_general(2, []int{3, 2, 6, 5, 0, 3}))
	})
}

func Test_221(t *testing.T) {
	t.Run("case1", func(t *testing.T) {
		assert.Equal(t, 4, maximalSquare(graph.BuildGraph("[[\"1\",\"0\",\"1\",\"0\",\"0\"],[\"1\",\"0\",\"1\",\"1\",\"1\"],[\"1\",\"1\",\"1\",\"1\",\"1\"],[\"1\",\"0\",\"0\",\"1\",\"0\"]]")))
	})
}
