package back_track

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_17(t *testing.T) {
	t.Run("case1", func(t *testing.T) {
		assert.Equal(t, []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}, letterCombinations("23"))
	})
}
