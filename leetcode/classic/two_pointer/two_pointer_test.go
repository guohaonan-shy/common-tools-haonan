package two_pointer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_11(t *testing.T) {
	t.Run("case1", func(t *testing.T) {
		case1 := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
		assert.Equal(t, 49, maxArea(case1))
	})

	t.Run("case2", func(t *testing.T) {
		case2 := []int{1, 2, 4, 3}
		assert.Equal(t, 4, maxArea(case2))
	})
}

func Test_15(t *testing.T) {
	t.Run("case1", func(t *testing.T) {
		case1 := []int{-1, 0, 1, 2, -1, -4}
		assert.Equal(t, [][]int{{-1, -1, 2}, {-1, 0, 1}}, threeSum(case1))
	})

	t.Run("case2", func(t *testing.T) {
		case2 := []int{0, 1, 1}
		assert.Equal(t, [][]int{}, threeSum(case2))
	})

	t.Run("case3", func(t *testing.T) {
		case3 := []int{0, 0, 0}
		assert.Equal(t, [][]int{{0, 0, 0}}, threeSum(case3))
	})
}

func Test_84(t *testing.T) {
	t.Run("case1", func(t *testing.T) {
		case1 := []int{2, 1, 5, 6, 2, 3}
		assert.Equal(t, 10, largestRectangleArea(case1))
	})

	t.Run("case2", func(t *testing.T) {
		case2 := []int{2, 4}
		assert.Equal(t, 4, largestRectangleArea(case2))
	})

	t.Run("case3", func(t *testing.T) {
		case3 := []int{0, 0}
		assert.Equal(t, 0, largestRectangleArea(case3))
	})

	t.Run("case4", func(t *testing.T) {
		case4 := []int{4, 2, 0, 3, 2, 4, 3, 4}
		assert.Equal(t, 10, largestRectangleArea(case4))
	})
}

func Test_125(t *testing.T) {
	t.Run("case1", func(t *testing.T) {
		case1 := "A man, a plan, a canal: Panama"
		assert.Equal(t, true, isPalindrome(case1))
	})

	t.Run("case2", func(t *testing.T) {
		case2 := "0P"
		assert.Equal(t, false, isPalindrome(case2))
	})
}
