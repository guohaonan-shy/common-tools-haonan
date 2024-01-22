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
