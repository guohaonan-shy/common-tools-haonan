package classic

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_88(t *testing.T) {
	num1 := []int{1, 2, 3, 0, 0, 0}
	merge(num1, 3, []int{2, 5, 6}, 3)
	assert.Equal(t, []int{1, 2, 2, 3, 5, 6}, num1)
}

func Test_26(t *testing.T) {

	t.Run("case1", func(t *testing.T) {
		case1 := []int{1, 1, 2}
		k1 := removeDuplicates(case1)
		assert.Equal(t, 2, k1)
		assert.Equal(t, []int{1, 2, 2}, case1)
	})

	t.Run("case2", func(t *testing.T) {
		case2 := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
		k2 := removeDuplicates(case2)
		assert.Equal(t, 5, k2)
		assert.Equal(t, []int{0, 1, 2, 3, 4, 2, 2, 3, 3, 4}, case2)
	})

}

func Test_27(t *testing.T) {

	t.Run("case1", func(t *testing.T) {
		case1 := []int{3, 2, 2, 3}
		k1 := removeElement2(case1, 3)
		assert.Equal(t, 2, k1)
		assert.Equal(t, []int{2, 2, 2, 3}, case1)
	})

	t.Run("case2", func(t *testing.T) {
		case2 := []int{0, 1, 2, 2, 3, 0, 4, 2}
		k2 := removeElement2(case2, 2)
		assert.Equal(t, 5, k2)
		assert.Equal(t, []int{0, 1, 3, 0, 4, 0, 4, 2}, case2)
	})

}
