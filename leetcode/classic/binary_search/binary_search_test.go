package binary_search

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_74(t *testing.T) {
	t.Run("case", func(t *testing.T) {
		assert.Equal(t, true, searchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 50}}, 10))
	})
}

func Test_153(t *testing.T) {
	t.Run("153", func(t *testing.T) {
		assert.Equal(t, 1, findMin([]int{3, 4, 5, 1, 2}))
	})

	t.Run("153_2", func(t *testing.T) {
		assert.Equal(t, 0, findMin([]int{4, 5, 6, 7, 0, 1, 2}))
	})

	t.Run("153_3", func(t *testing.T) {
		assert.Equal(t, 1, findMin([]int{2, 1}))
	})
}

func Test_4(t *testing.T) {
	type testCase struct {
		nums1    []int
		nums2    []int
		expected float64
	}

	for _, tc := range []testCase{
		{
			nums1:    []int{1, 3},
			nums2:    []int{2},
			expected: float64(2),
		},
		{
			nums1:    []int{1, 2},
			nums2:    []int{3, 4},
			expected: 2.5,
		},
		{
			nums1:    []int{1, 2, 3},
			nums2:    []int{4, 5, 6, 7, 8, 9},
			expected: float64(5),
		},
		{
			nums1:    []int{1, 4, 5, 7},
			nums2:    []int{2, 3, 6},
			expected: float64(4),
		},
		{
			nums1:    []int{1},
			nums2:    []int{2, 3, 4, 5, 6},
			expected: 3.5,
		},
	} {
		t.Run("4", func(t *testing.T) {
			assert.Equal(t, tc.expected, findMedianSortedArrays(tc.nums1, tc.nums2))
		})
	}
}

func Test_33(t *testing.T) {
	t.Run("rotate search in distinct array", func(t *testing.T) {
		assert.Equal(t, 4, rotateSearch_noDistinct([]int{4, 5, 6, 7, 0, 1, 2}, 0))
	})

	t.Run("can't find element", func(t *testing.T) {
		assert.Equal(t, -1, rotateSearch_noDistinct([]int{4, 5, 6, 7, 0, 1, 2}, 3))
	})

	t.Run("can't find element", func(t *testing.T) {
		assert.Equal(t, -1, rotateSearch_noDistinct([]int{1}, 3))
	})
}

func Test_81(t *testing.T) {
	t.Run("rotate search in duplicated array", func(t *testing.T) {
		assert.Equal(t, true, searchII([]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 1, 1, 1, 1}, 2))
	})
}
