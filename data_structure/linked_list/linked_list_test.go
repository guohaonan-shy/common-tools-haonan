package linked_list

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_LinkedList(t *testing.T) {
	t.Run("linked_list_case1", func(t *testing.T) {
		case1 := NewLinkedList()
		case1.Insert(0, 1)
		case1.Insert(1, 2)
		case1.Insert(2, 3)
		case1.Insert(3, 4)
		case1.Insert(4, 5)
		result := make([]int, 0)
		for head := case1.head; head != nil; head = head.next {
			result = append(result, head.item)
		}
		assert.Equal(t, []int{1, 2, 3, 4, 5}, result)
	})

	t.Run("linked_list_case2", func(t *testing.T) {
		case2 := NewLinkedList()
		case2.Insert(0, 1)
		case2.Insert(0, 2)
		case2.Insert(2, 3)
		case2.Insert(2, 4)
		case2.Insert(5, 5)
		result := make([]int, 0)
		for head := case2.head; head != nil; head = head.next {
			result = append(result, head.item)
		}
		assert.Equal(t, []int{2, 1, 4, 3, 5}, result)
	})

	t.Run("linked_list_case3", func(t *testing.T) {
		case3 := NewLinkedList()
		case3.Insert(0, 1)
		case3.Insert(1, 2)
		case3.Insert(2, 3)
		case3.Insert(3, 4)
		case3.Insert(4, 5)
		result1 := make([]int, 0)
		for head := case3.head; head != nil; head = head.next {
			result1 = append(result1, head.item)
		}
		assert.Equal(t, []int{1, 2, 3, 4, 5}, result1)

		case3.Remove(0)
		result2 := make([]int, 0)
		for head := case3.head; head != nil; head = head.next {
			result2 = append(result2, head.item)
		}
		assert.Equal(t, []int{2, 3, 4, 5}, result2)

		case3.Remove(3)
		result3 := make([]int, 0)
		for head := case3.head; head != nil; head = head.next {
			result3 = append(result3, head.item)
		}
		assert.Equal(t, []int{2, 3, 4}, result3)

		case3.Remove(1)
		result4 := make([]int, 0)
		for head := case3.head; head != nil; head = head.next {
			result4 = append(result4, head.item)
		}
		assert.Equal(t, []int{2, 4}, result4)

	})
}
