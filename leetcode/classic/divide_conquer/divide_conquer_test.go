package divide_conquer

import (
	"testing"

	. "github.com/common-tools-haonan/leetcode/classic/binary_tree"
	"github.com/common-tools-haonan/leetcode/classic/linked_list"
	"github.com/stretchr/testify/assert"
)

func Test_108(t *testing.T) {
	t.Run("case1", func(t *testing.T) {
		root := convert([]int{-10, -3, 0, 5, 9})
		assert.Equal(t, []int{-10, -3, 0, 5, 9}, PreorderTraversal(root))
	})
}

func Test_148(t *testing.T) {
	t.Run("case1", func(t *testing.T) {
		head := linked_list.BuildLinkedList([]int{4, 2, 1, 3})
		assert.Equal(t, []int{1, 2, 3, 4}, linked_list.ConvertToList(sortList(head)))
	})
}
