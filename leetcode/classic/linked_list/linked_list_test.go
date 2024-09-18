package linked_list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_2(t *testing.T) {
	t.Run("case1", func(t *testing.T) {
		l1 := buildLinkedList([]int{2, 4, 3})
		l2 := buildLinkedList([]int{5, 6, 4})

		sum1 := addTwoNumber(l1, l2)
		assert.Equal(t, []int{7, 0, 8}, convertToList(sum1))
	})
}

func Test_25(t *testing.T) {
	t.Run("case1", func(t *testing.T) {
		l1 := buildLinkedList([]int{1, 2, 3, 4, 5})
		res1 := reverseKGroup(l1, 2)

		list1 := convertToList(res1)
		assert.Equal(t, []int{2, 1, 4, 3, 5}, list1)
	})
}

func Test_61(t *testing.T) {
	t.Run("rotate_list", func(t *testing.T) {
		case1 := buildLinkedList([]int{1, 2, 3, 4, 5})
		expected1 := buildLinkedList([]int{4, 5, 1, 2, 3})
		assert.Equal(t, expected1, rotateRight(case1, 2))
	})

	t.Run("k > length", func(t *testing.T) {
		case2 := buildLinkedList([]int{0, 1, 2})
		expected2 := buildLinkedList([]int{2, 0, 1})
		assert.Equal(t, expected2, rotateRight(case2, 4))
	})
}

func Test_82(t *testing.T) {
	t.Run("82", func(t *testing.T) {
		case1 := buildLinkedList([]int{1, 2, 3, 3, 4, 4, 5})
		expected1 := buildLinkedList([]int{1, 2, 5})
		assert.Equal(t, expected1, deleteDuplicates(case1))
	})
}

func Test_92(t *testing.T) {
	t.Run("case1", func(t *testing.T) {
		case1 := buildLinkedList([]int{1, 2, 3, 4, 5})
		left, right := 2, 4
		assert.Equal(t, []int{1, 4, 3, 2, 5}, convertToList(reverseBetween(case1, left, right)))
	})

	t.Run("extreme corner case1", func(t *testing.T) {
		case2 := buildLinkedList([]int{1, 2, 3, 4, 5})
		left, right := 2, 5
		assert.Equal(t, []int{1, 5, 4, 3, 2}, convertToList(reverseBetween(case2, left, right)))
	})

	t.Run("extreme corner case2", func(t *testing.T) {
		case3 := buildLinkedList([]int{1, 2, 3, 4, 5})
		left, right := 1, 5
		assert.Equal(t, []int{5, 4, 3, 2, 1}, convertToList(reverseBetween(case3, left, right)))
	})

	t.Run("extreme corner case3", func(t *testing.T) {
		case3 := buildLinkedList([]int{1, 2, 3, 4, 5})
		left, right := 5, 5
		assert.Equal(t, []int{1, 2, 3, 4, 5}, convertToList(reverseBetween(case3, left, right)))
	})
}

func Test_146(t *testing.T) {
	type testCases struct {
		cmds   []string
		value  [][]int
		output []int
	}

	cases := []testCases{
		//{
		//	cmds:   []string{"LRUCache", "put", "get", "put", "get", "get"},
		//	value:  [][]int{{1}, {2, 1}, {2}, {3, 2}, {2}, {3}},
		//	output: []int{-1, -1, 1, -1, -1, 2},
		//},
		//{
		//	cmds:   []string{"LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"},
		//	value:  [][]int{{2}, {1, 1}, {2, 2}, {1}, {3, 3}, {2}, {4, 4}, {1}, {3}, {4}},
		//	output: []int{-1, -1, -1, 1, -1, -1, -1, -1, 3, 4},
		//},
		//{
		//	cmds:   []string{"LRUCache", "put", "put", "get", "put", "put"},
		//	value:  [][]int{{2}, {2, 1}, {2, 2}, {2}, {1, 1}, {4, 1}},
		//	output: []int{-1, -1, -1, 2, -1, -1},
		//},
		//{
		//	cmds:   []string{"LRUCache", "put", "put", "put", "put", "get", "get"},
		//	value:  [][]int{{2}, {2, 1}, {1, 1}, {2, 3}, {4, 1}, {1}, {2}},
		//	output: []int{-1, -1, -1, -1, -1, -1, 3},
		//},
		{
			cmds:   []string{"LRUCache", "get", "put", "get", "put", "put", "get", "get"},
			value:  [][]int{{2}, {2}, {2, 6}, {1}, {1, 5}, {1, 2}, {1}, {2}},
			output: []int{-1, -1, -1, -1, -1, -1, 2, 6},
		},
	}

	for i := range cases {
		t.Run("case1", func(t *testing.T) {
			temp := cases[i]
			length := len(temp.cmds)

			if temp.cmds[0] != "LRUCache" {
				t.Error("cmd wrong")
			}

			cache := Constructor(temp.value[0][0])

			for j := 1; j < length; j++ {
				value := temp.value[j]
				if temp.cmds[j] == "put" {
					cache.Put(value[0], value[1])
					assert.Equal(t, temp.output[j], -1)
				} else {
					result := cache.Get(value[0])
					assert.Equal(t, temp.output[j], result)
				}
			}

		})

	}

}
