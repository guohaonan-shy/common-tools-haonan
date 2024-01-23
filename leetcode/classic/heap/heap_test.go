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
