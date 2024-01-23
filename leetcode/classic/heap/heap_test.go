package heap

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_215(t *testing.T) {
	t.Run("heap", func(t *testing.T) {
		assert.Equal(t, 5, findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2))
	})
}

func Test_502(t *testing.T) {
	t.Run("heap", func(t *testing.T) {
		profit := []int{1, 2, 3}
		capital := []int{0, 1, 1}
		k, w := 2, 0
		assert.Equal(t, 4, findMaximizedCapital(k, w, profit, capital))
	})
}
