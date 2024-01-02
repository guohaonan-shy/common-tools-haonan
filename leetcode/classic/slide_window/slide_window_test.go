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

func Test_209(t *testing.T) {
	t.Run("case1", func(t *testing.T) {
		target1, case1 := 7, []int{2, 3, 1, 2, 4, 3}
		assert.Equal(t, 2, minSubArrayLen(target1, case1))
	})

	t.Run("case2", func(t *testing.T) {
		target2, case2 := 4, []int{1, 4, 4}
		assert.Equal(t, 1, minSubArrayLen(target2, case2))
	})

	t.Run("case3", func(t *testing.T) {
		target3, case3 := 11, []int{1, 1, 1, 1, 1, 1, 1, 1}
		assert.Equal(t, 0, minSubArrayLen(target3, case3))
	})

	t.Run("wrong_case4", func(t *testing.T) {
		target4, case4 := 11, []int{1, 2, 3, 4, 5}
		assert.Equal(t, 3, minSubArrayLen(target4, case4))
	})
}
