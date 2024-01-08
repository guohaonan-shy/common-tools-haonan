package stack

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_20(t *testing.T) {
	t.Run("case1", func(t *testing.T) {
		case1 := "()"
		assert.Equal(t, true, isValid(case1))
	})
}

func Test_71(t *testing.T) {
	t.Run("case1", func(t *testing.T) {
		case1 := "/home/"
		assert.Equal(t, "/home", simplifyPath(case1))
	})

	t.Run("case2", func(t *testing.T) {
		case2 := "/../"
		assert.Equal(t, "/", simplifyPath(case2))
	})

	t.Run("case3", func(t *testing.T) {
		case3 := "/home//foo/"
		assert.Equal(t, "/home/foo", simplifyPath(case3))
	})
}
