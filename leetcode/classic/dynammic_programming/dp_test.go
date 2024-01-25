package dynammic_programming

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_97(t *testing.T) {
	t.Run("dp97", func(t *testing.T) {
		assert.Equal(t, true, isInterleave("aabcc", "dbbca", "aadbbcbcac"))
	})

	t.Run("dp97", func(t *testing.T) {
		assert.Equal(t, false, isInterleave("aabcc", "dbbca", "aadbbbaccc"))
	})
}
