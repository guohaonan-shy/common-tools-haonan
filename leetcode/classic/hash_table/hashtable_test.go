package hash_table

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_128(t *testing.T) {
	t.Run("case1", func(t *testing.T) {
		case1 := []int{100, 4, 200, 1, 3, 2}
		assert.Equal(t, 4, longestConsecutive(case1))
	})

	t.Run("case2", func(t *testing.T) {
		case2 := []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}
		assert.Equal(t, 9, longestConsecutive(case2))
	})

}
