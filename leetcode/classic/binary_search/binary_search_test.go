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
