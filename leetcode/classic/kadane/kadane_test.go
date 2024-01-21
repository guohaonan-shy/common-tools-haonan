package kadane

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func Test_918(t *testing.T) {

	type testCase struct {
		Input    []int
		Expected int
	}

	for _, tc := range []testCase{
		{
			Input:    []int{5, -3, 5},
			Expected: 10,
		},
		{
			Input:    []int{1, -2, 3, -2},
			Expected: 3,
		},
		{
			Input:    []int{-3, -2, -3},
			Expected: -2,
		},
	} {
		input := ""
		for _, val := range tc.Input {
			input = input + strconv.Itoa(val) + "_"
		}
		input = input[:len(input)-1]
		t.Run(input, func(t *testing.T) {
			assert.Equal(t, tc.Expected, maxSubarraySumCircular(tc.Input))
		})
	}

}
