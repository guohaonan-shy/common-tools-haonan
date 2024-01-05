package matrix

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_54(t *testing.T) {
	t.Run("case1", func(t *testing.T) {
		case1 := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
		assert.Equal(t, []int{1, 2, 3, 6, 9, 8, 7, 4, 5}, spiralOrder(case1))
	})
}
