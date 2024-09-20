package binary_tree

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"strings"
	"testing"
)

func Test_treeify(t *testing.T) {
	type testCase struct {
		intput string
	}

	cases := []*testCase{
		{
			intput: "[3,9,20,null,null,15,7]",
		},
	}

	for _, tc := range cases {
		t.Run("case", func(t *testing.T) {

			root := treeify(tc.intput)
			assert.Equal(t, 3, root.Val)
			assert.Equal(t, 9, root.Left.Val)
			assert.Equal(t, 20, root.Right.Val)

			assert.Equal(t, (*TreeNode)(nil), root.Left.Left)
			assert.Equal(t, (*TreeNode)(nil), root.Left.Right)

			assert.Equal(t, 15, root.Right.Left.Val)
			assert.Equal(t, 7, root.Right.Right.Val)

			assert.Equal(t, (*TreeNode)(nil), root.Right.Right.Left)
		})
	}
}

func Test_100(t *testing.T) {
	type testCase struct {
		inputA   string
		inputB   string
		expected bool
	}

	tcs := []*testCase{
		{
			inputA:   "[1,2,3]",
			inputB:   "[1,2,3]",
			expected: true,
		},
		{
			inputA:   "[1,2]",
			inputB:   "[1,null,2]",
			expected: false,
		},
		{
			inputA:   "[1,2,1]",
			inputB:   "[1,1,2]",
			expected: false,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.inputA+string('\t')+string('\t')+tc.inputB, func(t *testing.T) {
			assert.Equal(t, tc.expected, isSameTree(treeify(tc.inputA), treeify(tc.inputB)))
		})
	}
}

func Test_104(t *testing.T) {
	type testCase struct {
		input    string
		expected int
	}

	tcs := []*testCase{
		{
			input:    "[3,9,20,null,null,15,7]",
			expected: 3,
		},
		{
			input:    "[1,null,2]",
			expected: 2,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.input, func(t *testing.T) {
			assert.Equal(t, tc.expected, maxDepth(treeify(tc.input)))
		})
	}
}

func Test_114(t *testing.T) {
	t.Run("case1", func(t *testing.T) {
		input := "[1,2,5,3,4,null,6]"
		tree := treeify(input)
		flatten(tree)
		flattenV2(tree)
		print("done")
	})
}

func Test_124(t *testing.T) {
	for _, tc := range []struct {
		input    string
		expected int
	}{
		{
			input:    "[1,2,3]",
			expected: 6,
		},
		{
			input:    "[-10,9,20,null,null,15,7]",
			expected: 42,
		},
	} {

		t.Run(tc.input, func(t *testing.T) {
			res := maxPathSum(treeify(tc.input))
			assert.Equal(t, tc.expected, res)
		})

	}
}

func Test_129(t *testing.T) {
	for _, tc := range []struct {
		input    string
		expected int
	}{
		{
			input:    "[1,2,3]",
			expected: 25,
		},
		{
			input:    "[4,9,0,5,1]",
			expected: 1026,
		},
	} {

		t.Run(tc.input, func(t *testing.T) {
			res := sumNumbers(treeify(tc.input))
			assert.Equal(t, tc.expected, res)
		})

	}
}

func Test_173(t *testing.T) {
	for _, tc := range []struct {
		cmds     []string
		tree     string
		expected []string
	}{
		{
			cmds:     []string{"BSTIterator", "next", "next", "hasNext", "next", "hasNext", "next", "hasNext", "next", "hasNext"},
			tree:     "[7, 3, 15, null, null, 9, 20]",
			expected: []string{"null", "3", "7", "true", "9", "true", "15", "true", "20", "false"},
		},
	} {
		t.Run("case", func(t *testing.T) {
			var bst BSTIterator
			for i, cmd := range tc.cmds {
				if cmd == "BSTIterator" {
					bst = Constructor(treeify(tc.tree))
				} else if cmd == "next" {
					expect, _ := strconv.Atoi(tc.expected[i])
					assert.Equal(t, bst.Next(), expect)
				} else {
					expect, _ := strconv.ParseBool(tc.expected[i])
					assert.Equal(t, bst.HasNext(), expect)
				}
			}
		})
	}
}

func Test_222(t *testing.T) {
	for _, tc := range []struct {
		input    string
		expected int
	}{
		{
			input:    "[1,2,3,4,5,6]",
			expected: 6,
		},
		{
			input:    "[]",
			expected: 0,
		},
	} {

		t.Run(tc.input, func(t *testing.T) {
			res := countNodes(treeify(tc.input))
			assert.Equal(t, tc.expected, res)
		})

	}
}

func Test_226(t *testing.T) {
	testCases := []struct {
		input string
	}{
		{input: "[4,2,7,1,3,6,9]"},
	}

	for _, tc := range testCases {
		formetInput := strings.TrimLeft(tc.input, "[")
		formetInput = strings.TrimRight(formetInput, "]")
		elements := strings.Split(formetInput, ",")
		t.Run(tc.input, func(t *testing.T) {
			inverted := invertTree(treeify(tc.input))

			assert.Equal(t, elements[0], strconv.Itoa(inverted.Val))
			assert.Equal(t, elements[2], strconv.Itoa(inverted.Left.Val))
			assert.Equal(t, elements[1], strconv.Itoa(inverted.Right.Val))

			assert.Equal(t, elements[3], strconv.Itoa(inverted.Right.Right.Val))
			assert.Equal(t, elements[4], strconv.Itoa(inverted.Right.Left.Val))
		})
	}
}
