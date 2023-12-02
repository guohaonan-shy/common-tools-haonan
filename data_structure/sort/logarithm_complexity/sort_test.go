package logarithm_complexity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_MergeSort(t *testing.T) {
	t.Run("merge_sort_case1", func(t *testing.T) {
		test_case := []int{5, 4, 3, 2, 1}
		result := MergeSort(test_case)
		assert.Equal(t, []int{1, 2, 3, 4, 5}, result)
	})

	t.Run("merge_sort_case2", func(t *testing.T) {
		test_case := []int{1, 2, 3, 4, 5}
		result := MergeSort(test_case)
		assert.Equal(t, []int{1, 2, 3, 4, 5}, result)
	})
	//
	t.Run("merge_sort_case3", func(t *testing.T) {
		test_case := []int{5, 1, 2, 4, 3}
		result := MergeSort(test_case)
		assert.Equal(t, []int{1, 2, 3, 4, 5}, result)
	})

}
