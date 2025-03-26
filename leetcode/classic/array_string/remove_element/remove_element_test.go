package remove_element

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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

	checkFunc := func(nums []int, val int) bool {
		for _, num := range nums {
			if num == val {
				return false
			}
		}
		return true
	}

	t.Run("case1", func(t *testing.T) {
		case1 := []int{3, 2, 2, 3}
		k1 := removeElementV3(case1, 3)
		assert.Equal(t, 2, k1)
		assert.Equal(t, []int{2, 2, 2, 3}, case1)
		assert.Equal(t, true, checkFunc(case1[:k1], 3))
	})

	t.Run("case2", func(t *testing.T) {
		case2 := []int{0, 1, 2, 2, 3, 0, 4, 2}
		k2 := removeElementV3(case2, 2)
		assert.Equal(t, 5, k2)
		assert.Equal(t, []int{0, 1, 3, 0, 4, 0, 4, 2}, case2)
		assert.Equal(t, true, checkFunc(case2[:k2], 2))
	})

	t.Run("special_case", func(t *testing.T) {
		special_case := []int{3}
		k3 := removeElementV3(special_case, 3)
		assert.Equal(t, 0, k3)
		assert.Equal(t, []int{3}, special_case)
		assert.Equal(t, true, checkFunc(special_case[:k3], 3))
	})

}

func Test_80(t *testing.T) {
	t.Run("case1", func(t *testing.T) {
		case1 := []int{1, 1, 1, 2, 2, 3}
		k1 := removeDuplicates_atMostTwice(case1)
		assert.Equal(t, 5, k1)
		assert.Equal(t, []int{1, 1, 2, 2, 3, 3}, case1)
		case11 := []int{1, 1, 1, 2, 2, 3}
		k11 := removeDuplicates_atMostTwice_standard(case11)
		assert.Equal(t, 5, k11)
		assert.Equal(t, []int{1, 1, 2, 2, 3, 3}, case11)
	})

	t.Run("case2", func(t *testing.T) {
		case2 := []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
		k2 := removeDuplicates_atMostTwice(case2)
		assert.Equal(t, 7, k2)
		assert.Equal(t, []int{0, 0, 1, 1, 2, 3, 3, 3, 3}, case2)
	})

	t.Run("special_case1", func(t *testing.T) {
		special_case1 := []int{1}
		k3 := removeDuplicates_atMostTwice(special_case1)
		assert.Equal(t, 1, k3)
		assert.Equal(t, []int{1}, special_case1)
	})

	t.Run("special_case2", func(t *testing.T) {
		special_case2 := []int{2, 2}
		k4 := removeDuplicates_atMostTwice(special_case2)
		assert.Equal(t, 2, k4)
		assert.Equal(t, []int{2, 2}, special_case2)
	})

	t.Run("special_case3", func(t *testing.T) {
		special_case3 := []int{2, 3}
		k5 := removeDuplicates_atMostTwice(special_case3)
		assert.Equal(t, 2, k5)
		assert.Equal(t, []int{2, 3}, special_case3)
	})

}
