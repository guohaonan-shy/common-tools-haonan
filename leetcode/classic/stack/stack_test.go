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

func Test_84(t *testing.T) {
	t.Run("case1", func(t *testing.T) {
		case1 := []int{2, 1, 5, 6, 2, 3}
		assert.Equal(t, 10, largestRectangleArea(case1))
	})

	t.Run("case2", func(t *testing.T) {
		case2 := []int{2, 4}
		assert.Equal(t, 4, largestRectangleArea(case2))
	})

	t.Run("case3", func(t *testing.T) {
		case3 := []int{0, 0}
		assert.Equal(t, 0, largestRectangleArea(case3))
	})

	t.Run("case4", func(t *testing.T) {
		case4 := []int{4, 2, 0, 3, 2, 4, 3, 4}
		assert.Equal(t, 10, largestRectangleArea(case4))
	})
}

func Test_224(t *testing.T) {
	t.Run("case1", func(t *testing.T) {
		case1 := "1 + 1"
		assert.Equal(t, 2, calculate(case1))
	})

	t.Run("case2", func(t *testing.T) {
		case2 := " 2-1 + 2 "
		assert.Equal(t, 3, calculate(case2))
	})

	t.Run("case3", func(t *testing.T) {
		case3 := "(1+(4+5+2)-3)+(6+8)"
		assert.Equal(t, 23, calculate(case3))
	})
	t.Run("case4", func(t *testing.T) {
		case4 := "2147483647"
		assert.Equal(t, 2147483647, calculate(case4))
	})

	for _, test := range []struct {
		case1    string
		expected int
	}{{"1 + 1", 2}, {" 2-1 + 2 ", 3}, {"(1+(4+5+2)-3)+(6+8)", 23}, {"2147483647", 2147483647}, {"0", 0}} {
		t.Run("case", func(t *testing.T) {
			assert.Equal(t, test.expected, calculate_standard(test.case1))
		})
	}

}
