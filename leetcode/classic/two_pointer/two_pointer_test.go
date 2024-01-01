package two_pointer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_11(t *testing.T) {
	t.Run("case1", func(t *testing.T) {
		case1 := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
		assert.Equal(t, 49, maxArea(case1))
	})

	t.Run("case2", func(t *testing.T) {
		case2 := []int{1, 2, 4, 3}
		assert.Equal(t, 4, maxArea(case2))
	})
}

func Test_125(t *testing.T) {
	t.Run("case1", func(t *testing.T) {
		case1 := "A man, a plan, a canal: Panama"
		assert.Equal(t, true, isPalindrome(case1))
	})

	t.Run("case2", func(t *testing.T) {
		case2 := "0P"
		assert.Equal(t, false, isPalindrome(case2))
	})
}
