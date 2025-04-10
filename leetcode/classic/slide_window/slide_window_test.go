package slide_window

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_3(t *testing.T) {
	t.Run("case1", func(t *testing.T) {
		case1 := "abcabcbb"
		assert.Equal(t, 3, lengthOfLongestSubstring(case1))
	})
	t.Run("case2", func(t *testing.T) {
		case2 := "bbbbb"
		assert.Equal(t, 1, lengthOfLongestSubstring(case2))
	})
	t.Run("case3", func(t *testing.T) {
		case3 := "pwwkew"
		assert.Equal(t, 3, lengthOfLongestSubstring(case3))
	})

	t.Run("wrong_case1", func(t *testing.T) {
		wrong_case1 := "tmmzuxt"
		assert.Equal(t, 5, lengthOfLongestSubstring(wrong_case1))
	})
}

func Test_30(t *testing.T) {
	t.Run("case1", func(t *testing.T) {
		case1, words := "barfoothefoobarman", []string{"foo", "bar"}
		assert.Equal(t, []int{0, 9}, findSubstring(case1, words))
	})

	t.Run("case2", func(t *testing.T) {
		case2, words := "wordgoodgoodgoodbestword", []string{"word", "good", "best", "word"}
		assert.Equal(t, []int{}, findSubstring(case2, words))
	})

	t.Run("case3", func(t *testing.T) {
		case3, words := "a", []string{"a"}
		assert.Equal(t, []int{0}, findSubstring(case3, words))
	})

	t.Run("case4", func(t *testing.T) {
		case4, words := "wordgoodgoodgoodbestword", []string{"word", "good", "best", "good"}
		assert.Equal(t, []int{8}, findSubstringV2(case4, words))
	})
}

func Test_76(t *testing.T) {
	t.Run("case1", func(t *testing.T) {
		s1, t1 := "ADOBECODEBANC", "ABC"
		assert.Equal(t, "BANC", minWindowV2(s1, t1))
	})
}

func Test_209(t *testing.T) {
	t.Run("case1", func(t *testing.T) {
		target1, case1 := 7, []int{2, 3, 1, 2, 4, 3}
		assert.Equal(t, 2, minSubArrayLen(target1, case1))
		assert.Equal(t, 2, minSubArrayLenV2(target1, case1))
	})

	t.Run("case2", func(t *testing.T) {
		target2, case2 := 4, []int{1, 4, 4}
		assert.Equal(t, 1, minSubArrayLen(target2, case2))
		assert.Equal(t, 1, minSubArrayLenV2(target2, case2))
	})

	t.Run("case3", func(t *testing.T) {
		target3, case3 := 11, []int{1, 1, 1, 1, 1, 1, 1, 1}
		assert.Equal(t, 0, minSubArrayLen(target3, case3))
		assert.Equal(t, 0, minSubArrayLenV2(target3, case3))
	})

	t.Run("wrong_case4", func(t *testing.T) {
		target4, case4 := 11, []int{1, 2, 3, 4, 5}
		assert.Equal(t, 3, minSubArrayLen(target4, case4))
		assert.Equal(t, 3, minSubArrayLenV2(target4, case4))
	})

	t.Run("corner case", func(t *testing.T) {
		target5, case5 := 15, []int{1, 2, 3, 4, 5}
		assert.Equal(t, 5, minSubArrayLenV2(target5, case5))
	})
}

func Test_159(t *testing.T) {
	t.Run("longest substring with at most two distinct characters", func(t *testing.T) {
		assert.Equal(t, 3, lengthOfLongestSubstringTwoDistinct("eceba"))
		assert.Equal(t, 3, lengthOfLongestSubstringTwoDistinctV2("eceba"))
	})

	t.Run("longest substring with at most two distinct characters", func(t *testing.T) {
		assert.Equal(t, 5, lengthOfLongestSubstringTwoDistinct("ccaabbb"))
		assert.Equal(t, 5, lengthOfLongestSubstringTwoDistinctV2("ccaabbb"))
	})
}

func Test_340(t *testing.T) {
	t.Run("with k distinct", func(t *testing.T) {
		assert.Equal(t, 3, lengthOfLongestSubstringKDistinct("eceba", 2))
		assert.Equal(t, 3, lengthOfLongestSubstringKDistinctV2("eceba", 2))
		assert.Equal(t, 2, lengthOfLongestSubstringKDistinctV2("aa", 1))
	})

}

func Test_239(t *testing.T) {
	t.Run("maximum window", func(t *testing.T) {
		assert.Equal(t, []int{3, 3, 5, 5, 6, 7}, maxSlidingWindow_heap([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
	})
}

func Test_32(t *testing.T) {
	t.Run("longest valid parentheses", func(t *testing.T) {
		assert.Equal(t, 2, longestValidParentheses("(()"))
	})

	t.Run("longest valid parentheses", func(t *testing.T) {
		assert.Equal(t, 4, longestValidParentheses(")()())"))
	})
}
