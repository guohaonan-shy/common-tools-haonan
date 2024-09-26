package dynammic_programming

import (
	"testing"

	"github.com/common-tools-haonan/leetcode/classic/graph"
	"github.com/stretchr/testify/assert"
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

func Test_139(t *testing.T) {
	t.Run("split_words case1", func(t *testing.T) {
		assert.Equal(t, true, splitWord("leetcode", []string{"leet", "code"}))
	})

	t.Run("split_words case2", func(t *testing.T) {
		assert.Equal(t, true, splitWord("applepenapple", []string{"apple", "pen"}))
	})

	t.Run("split_words case3", func(t *testing.T) {
		assert.Equal(t, false, splitWord("catsandog", []string{"cats", "dog", "sand", "and", "cat"}))
	})

	t.Run("split_words case4", func(t *testing.T) {
		assert.Equal(t, false, splitWord("a", []string{"b"}))
	})

	t.Run("split_words case5", func(t *testing.T) {
		assert.Equal(t, true, splitWord("b", []string{"b", "bbb", "bbbb"}))
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
