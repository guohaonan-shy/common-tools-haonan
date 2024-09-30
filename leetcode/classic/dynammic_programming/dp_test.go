package dynammic_programming

import (
	"testing"

	"github.com/common-tools-haonan/leetcode/classic/graph"
	"github.com/stretchr/testify/assert"
)

func Test_5(t *testing.T) {
	t.Run("case1", func(t *testing.T) {
		case1 := "babad"
		assert.Equal(t, "bab", longestPalindrome(case1))
	})

	t.Run("case2", func(t *testing.T) {
		case2 := "cbbd"
		assert.Equal(t, "bb", longestPalindrome(case2))
	})

	t.Run("corner case 1", func(t *testing.T) {
		cornerCase1 := "a"
		assert.Equal(t, "a", longestPalindrome(cornerCase1))
	})
}

func Test_44(t *testing.T) {
	t.Run("wildcard matching", func(t *testing.T) {
		assert.Equal(t, true, isMatch("adceb", "*a*b"))
	})
}

func Test_63(t *testing.T) {
	t.Run("case1", func(t *testing.T) {
		grid := [][]int{
			{0, 0, 0},
			{0, 1, 0},
			{0, 0, 0},
		}

		assert.Equal(t, 2, uniquePathsWithObstacles(grid))
	})

	t.Run("case1", func(t *testing.T) {
		grid := [][]int{
			{0, 1},
			{0, 0},
		}

		assert.Equal(t, 1, uniquePathsWithObstacles(grid))
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

func Test_120(t *testing.T) {
	t.Run("case1", func(t *testing.T) {
		triangle := [][]int{
			{2},
			{3, 4},
			{6, 5, 7},
			{4, 1, 8, 3},
		}
		actual := minimumTotal(triangle)
		assert.Equal(t, 11, actual)
	})

	t.Run("case2", func(t *testing.T) {
		triangle := [][]int{
			{-10},
		}
		actual := minimumTotal(triangle)
		assert.Equal(t, -10, actual)
	})

	t.Run("case1 for less-memory method", func(t *testing.T) {
		triangle := [][]int{
			{2},
			{3, 4},
			{6, 5, 7},
			{4, 1, 8, 3},
		}
		actual := minimumTotalWithLessMemory(triangle)
		assert.Equal(t, 11, actual)
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

func Test_300(t *testing.T) {
	t.Run("case1", func(t *testing.T) {
		assert.Equal(t, 4, longestIncreasingSequence([]int{10, 9, 2, 5, 3, 7, 101, 18}))
	})

	t.Run("case2", func(t *testing.T) {
		assert.Equal(t, 4, longestIncreasingSequence([]int{5, 6, 7, 1, 2, 3, 4}))
	})

	t.Run("case3", func(t *testing.T) {
		assert.Equal(t, 4, longestIncreasingSequence([]int{0, 1, 0, 3, 2, 3}))
	})

	t.Run("case4", func(t *testing.T) {
		assert.Equal(t, 1, longestIncreasingSequence([]int{7, 7, 7, 7, 7, 7, 7}))
	})

	t.Run("binary search case1", func(t *testing.T) {
		assert.Equal(t, 4, longestIncreasingSequence_binarySearch([]int{10, 9, 2, 5, 3, 7, 101, 18}))
	})

	t.Run("binary search case2", func(t *testing.T) {
		assert.Equal(t, 4, longestIncreasingSequence_binarySearch([]int{5, 6, 7, 1, 2, 3, 4}))
	})

	t.Run("binary search case3", func(t *testing.T) {
		assert.Equal(t, 4, longestIncreasingSequence_binarySearch([]int{0, 1, 0, 3, 2, 3}))
	})

	t.Run("binary search case4", func(t *testing.T) {
		assert.Equal(t, 1, longestIncreasingSequence_binarySearch([]int{7, 7, 7, 7, 7, 7, 7}))
	})

	t.Run("binary search case5", func(t *testing.T) {
		assert.Equal(t, 3, longestIncreasingSequence_binarySearch([]int{10, 9, 2, 5, 3, 4}))
	})
}

func Test_322(t *testing.T) {
	t.Run("case1", func(t *testing.T) {
		coinPackage := []int{1, 2, 5}
		target := 11
		assert.Equal(t, 3, coinChange(coinPackage, target))
	})

	t.Run("case2", func(t *testing.T) {
		coinPackage := []int{2}
		target := 3
		assert.Equal(t, -1, coinChange(coinPackage, target))
	})

	t.Run("case3", func(t *testing.T) {
		coinPackage := []int{1}
		target := 0
		assert.Equal(t, 0, coinChange(coinPackage, target))
	})
}
