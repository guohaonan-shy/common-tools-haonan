package linked_list

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

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
