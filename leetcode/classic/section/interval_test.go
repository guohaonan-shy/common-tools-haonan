package section

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_57(t *testing.T) {
	t.Run("case1", func(t *testing.T) {
		case1 := [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}
		newInterval := []int{4, 8}

		assert.Equal(t, [][]int{{1, 2}, {3, 10}, {12, 16}}, insert(case1, newInterval))
	})

	t.Run("case2", func(t *testing.T) {
		case2 := [][]int{{1, 5}}
		newInterval := []int{6, 8}
		assert.Equal(t, [][]int{{1, 5}, {6, 8}}, insert(case2, newInterval))
	})

	t.Run("case3", func(t *testing.T) {
		case3 := [][]int{{1, 3}, {6, 9}}
		newInterval := []int{2, 5}
		assert.Equal(t, [][]int{{1, 5}, {6, 9}}, insert(case3, newInterval))
	})
}

func Test_452(t *testing.T) {
	t.Run("case1", func(t *testing.T) {
		case1 := [][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}}
		assert.Equal(t, 2, findMinArrowShots(case1))
	})

	t.Run("case2", func(t *testing.T) {
		case2 := [][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}}
		assert.Equal(t, 4, findMinArrowShots(case2))
	})

	t.Run("case3", func(t *testing.T) {
		case3 := [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}}
		assert.Equal(t, 2, findMinArrowShots(case3))
	})
}
