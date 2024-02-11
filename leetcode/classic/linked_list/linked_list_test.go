package linked_list

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

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

func Test_146(t *testing.T) {
	type testCases struct {
		cmds   []string
		value  [][]int
		output []int
	}

	cases := []testCases{
		{
			cmds:   []string{"LRUCache", "put", "get", "put", "get", "get"},
			value:  [][]int{{1}, {2, 1}, {2}, {3, 2}, {2}, {3}},
			output: []int{-1, -1, 1, -1, -1, 2},
		},
		{
			cmds:   []string{"LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"},
			value:  [][]int{{2}, {1, 1}, {2, 2}, {1}, {3, 3}, {2}, {4, 4}, {1}, {3}, {4}},
			output: []int{-1, -1, -1, 1, -1, -1, -1, -1, 3, 4},
		},
		{
			cmds:   []string{"LRUCache", "put", "put", "get", "put", "put"},
			value:  [][]int{{2}, {2, 1}, {2, 2}, {2}, {1, 1}, {4, 1}},
			output: []int{-1, -1, -1, 2, -1, -1},
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
