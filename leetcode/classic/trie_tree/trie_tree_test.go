package trie_tree

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func Test_TrieTree(t *testing.T) {
	type TestCase struct {
		Cmds     []string
		Inputs   []string
		Expected []bool
	}

	for _, tc := range []TestCase{
		{
			Cmds:     []string{"Trie", "insert", "search", "search", "startsWith", "insert", "search"},
			Inputs:   []string{"", "apple", "apple", "app", "app", "app", "app"},
			Expected: []bool{false, false, true, false, true, false, true},
		},
	} {
		t.Run(strings.Join(tc.Cmds, " "), func(t *testing.T) {
			var TrieTree Trie
			for i, cmd := range tc.Cmds {

				if cmd == "Trie" {
					TrieTree = Constructor()
				} else if cmd == "insert" {
					TrieTree.Insert(tc.Inputs[i])
				} else if cmd == "search" {
					assert.Equal(t, tc.Expected[i], TrieTree.Search(tc.Inputs[i]))
				} else {
					assert.Equal(t, tc.Expected[i], TrieTree.StartWith(tc.Inputs[i]))
				}
			}
		})
	}
}
