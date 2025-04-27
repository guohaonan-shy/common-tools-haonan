package two_pointer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_11(t *testing.T) {
	t.Run("case1", func(t *testing.T) {
		case1 := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
		assert.Equal(t, 49, maxArea(case1))
		assert.Equal(t, 49, maxAreaV2(case1))
	})

	t.Run("case2", func(t *testing.T) {
		case2 := []int{1, 2, 4, 3}
		assert.Equal(t, 4, maxArea(case2))
		assert.Equal(t, 4, maxAreaV2(case2))
	})
}

func Test_15(t *testing.T) {
	t.Run("case1", func(t *testing.T) {
		case1 := []int{-1, 0, 1, 2, -1, -4}
		assert.Equal(t, [][]int{{-1, -1, 2}, {-1, 0, 1}}, threeSum(case1))

		assert.Equal(t, [][]int{{-1, -1, 2}, {-1, 0, 1}}, threeSumV2(case1))
	})

	t.Run("case2", func(t *testing.T) {
		case2 := []int{0, 1, 1}
		assert.Equal(t, [][]int{}, threeSum(case2))
		assert.Equal(t, [][]int{}, threeSumV2(case2))
	})

	t.Run("case3", func(t *testing.T) {
		case3 := []int{0, 0, 0}
		assert.Equal(t, [][]int{{0, 0, 0}}, threeSum(case3))
		assert.Equal(t, [][]int{{0, 0, 0}}, threeSumV2(case3))
	})
}

func Test_31(t *testing.T) {
	for _, testCase := range []struct {
		name   string
		input  []int
		wanted []int
	}{
		{
			name:   "test case 1",
			input:  []int{1, 2, 3},
			wanted: []int{1, 3, 2},
		},
		{
			name:   "test case 2",
			input:  []int{3, 2, 1},
			wanted: []int{1, 2, 3},
		},
		{
			name:   "test case 3",
			input:  []int{1, 1, 5},
			wanted: []int{1, 5, 1},
		},
	} {
		t.Run(testCase.name, func(t *testing.T) {
			nextPermutation(testCase.input)
			assert.Equal(t, testCase.wanted, testCase.input)
		})
	}
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
