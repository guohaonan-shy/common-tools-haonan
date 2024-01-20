package trie_tree

import (
	"github.com/common-tools-haonan/leetcode/classic/graph"
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
					assert.Equal(t, tc.Expected[i], TrieTree.StartsWith(tc.Inputs[i]))
				}
			}
		})
	}
}

func Test_WordDictionary(t *testing.T) {
	type TestCase struct {
		Cmds     []string
		Inputs   []string
		Expected []bool
	}

	for _, tc := range []TestCase{
		{
			Cmds:     []string{"WordDictionary", "addWord", "addWord", "addWord", "search", "search", "search", "search"},
			Inputs:   []string{"", "bad", "dad", "mad", "pad", "bad", ".ad", "b.."},
			Expected: []bool{false, false, false, false, false, true, true, true},
		},
	} {
		t.Run(strings.Join(tc.Cmds, " "), func(t *testing.T) {
			var dict WordDictionary
			for i, cmd := range tc.Cmds {

				if cmd == "WordDictionary" {
					dict = WordConstructor()
				} else if cmd == "addWord" {
					dict.AddWord(tc.Inputs[i])
				} else if cmd == "search" {
					assert.Equal(t, tc.Expected[i], dict.Search(tc.Inputs[i]))
				}
			}
		})
	}
}

func Test_212(t *testing.T) {
	type TestCase struct {
		Board    [][]byte
		Words    []string
		Expected []string
	}

	for _, tc := range []TestCase{
		//{
		//	Board:    graph.BuildGraph("[[\"o\",\"a\",\"a\",\"n\"],[\"e\",\"t\",\"a\",\"e\"],[\"i\",\"h\",\"k\",\"r\"],[\"i\",\"f\",\"l\",\"v\"]]"),
		//	Words:    []string{"oath", "pea", "eat", "rain"},
		//	Expected: []string{"eat", "oath"},
		//},
		{
			Board:    graph.BuildGraph("[[\"a\",\"b\",\"c\",\"e\"],[\"x\",\"x\",\"c\",\"d\"],[\"x\",\"x\",\"b\",\"a\"]]"),
			Words:    []string{"abc", "abcd"},
			Expected: []string{"abc", "abcd"},
		},
	} {
		t.Run(strings.Join(tc.Expected, "+"), func(t *testing.T) {
			assert.Equal(t, tc.Expected, findWords(tc.Board, tc.Words))
		})
	}

}
