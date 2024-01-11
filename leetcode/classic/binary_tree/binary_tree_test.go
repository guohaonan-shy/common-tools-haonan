package binary_tree

import (
	"github.com/stretchr/testify/assert"
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
