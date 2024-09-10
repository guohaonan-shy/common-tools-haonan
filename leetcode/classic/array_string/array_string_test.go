package array_string

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_12(t *testing.T) {
	t.Run("case1", func(t *testing.T) {
		assert.Equal(t, "III", intToRoman(3))
	})

	t.Run("case2", func(t *testing.T) {
		assert.Equal(t, "LVIII", intToRoman(58))
	})

	t.Run("case3", func(t *testing.T) {
		assert.Equal(t, "MCMXCIV", intToRoman(1994))
	})

	t.Run("v2 cases", func(t *testing.T) {
		assert.Equal(t, "III", intToRomanV2(3))
		assert.Equal(t, "LVIII", intToRomanV2(58))
		assert.Equal(t, "MCMXCIV", intToRomanV2(1994))
	})
}

func Test_13(t *testing.T) {
	t.Run("case1", func(t *testing.T) {
		assert.Equal(t, 3, romanToInt("III"))
	})

	t.Run("case2", func(t *testing.T) {
		assert.Equal(t, 58, romanToInt("LVIII"))
	})
}

func Test_26(t *testing.T) {

	t.Run("case1", func(t *testing.T) {
		case1 := []int{1, 1, 2}
		k1 := removeDuplicates(case1)
		assert.Equal(t, 2, k1)
		assert.Equal(t, []int{1, 2, 2}, case1)
	})

	t.Run("case2", func(t *testing.T) {
		case2 := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
		k2 := removeDuplicates(case2)
		assert.Equal(t, 5, k2)
		assert.Equal(t, []int{0, 1, 2, 3, 4, 2, 2, 3, 3, 4}, case2)
	})

}

func Test_27(t *testing.T) {

	t.Run("case1", func(t *testing.T) {
		case1 := []int{3, 2, 2, 3}
		k1 := removeElement2(case1, 3)
		assert.Equal(t, 2, k1)
		assert.Equal(t, []int{2, 2, 2, 3}, case1)
	})

	t.Run("case2", func(t *testing.T) {
		case2 := []int{0, 1, 2, 2, 3, 0, 4, 2}
		k2 := removeElement2(case2, 2)
		assert.Equal(t, 5, k2)
		assert.Equal(t, []int{0, 1, 3, 0, 4, 0, 4, 2}, case2)
	})

	t.Run("special_case", func(t *testing.T) {
		special_case := []int{3}
		k3 := removeElement2(special_case, 3)
		assert.Equal(t, 0, k3)
		assert.Equal(t, []int{3}, special_case)
	})

}

func Test_28(t *testing.T) {
	t.Run("case1", func(t *testing.T) {
		haystack1, needle1 := "mississippi", "issip"
		assert.Equal(t, 4, strStr(haystack1, needle1))
	})

	t.Run("case2", func(t *testing.T) {
		haystack2, needle2 := "sadbutsad", "sad"
		assert.Equal(t, 0, strStr(haystack2, needle2))
	})
}

func Test_42(t *testing.T) {
	t.Run("case1", func(t *testing.T) {
		case1 := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
		assert.Equal(t, 6, trap(case1))
		assert.Equal(t, 6, trap_stack(case1))
	})

	t.Run("case2", func(t *testing.T) {
		case2 := []int{4, 2, 0, 3, 2, 5}
		assert.Equal(t, 9, trap(case2))
		assert.Equal(t, 9, trap_stack(case2))
	})
}

func Test_65(t *testing.T) {
	t.Run(".1", func(t *testing.T) {
		assert.Equal(t, true, isNumber(".1"))
	})

	t.Run("4e+", func(t *testing.T) {
		assert.Equal(t, false, isNumber("4e+"))
	})

	t.Run(".-4", func(t *testing.T) {
		assert.Equal(t, false, isNumber(".-4"))
	})

	t.Run("+.8", func(t *testing.T) {
		assert.Equal(t, true, isNumber("+.8"))
	})
}

func Test_68(t *testing.T) {
	t.Run("case1", func(t *testing.T) {
		case1 := []string{"This", "is", "an", "example", "of", "text", "justification."}
		max_width1 := 16
		assert.Equal(t, []string{"This    is    an", "example  of text", "justification.  "}, fullJustify(case1, max_width1))
	})

	t.Run("case2", func(t *testing.T) {
		case2 := []string{"What", "must", "be", "acknowledgment", "shall", "be"}
		max_width2 := 16
		assert.Equal(t, []string{"What   must   be", "acknowledgment  ", "shall be        "}, fullJustify(case2, max_width2))
	})
}

func Test_80(t *testing.T) {
	t.Run("case1", func(t *testing.T) {
		case1 := []int{1, 1, 1, 2, 2, 3}
		k1 := removeDuplicates_atMostTwice(case1)
		assert.Equal(t, 5, k1)
		assert.Equal(t, []int{1, 1, 2, 2, 3, 3}, case1)
		case11 := []int{1, 1, 1, 2, 2, 3}
		k11 := removeDuplicates_atMostTwice_standard(case11)
		assert.Equal(t, 5, k11)
		assert.Equal(t, []int{1, 1, 2, 2, 3, 3}, case11)
	})

	t.Run("case2", func(t *testing.T) {
		case2 := []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
		k2 := removeDuplicates_atMostTwice(case2)
		assert.Equal(t, 7, k2)
		assert.Equal(t, []int{0, 0, 1, 1, 2, 3, 3, 3, 3}, case2)
	})

	t.Run("special_case1", func(t *testing.T) {
		special_case1 := []int{1}
		k3 := removeDuplicates_atMostTwice(special_case1)
		assert.Equal(t, 1, k3)
		assert.Equal(t, []int{1}, special_case1)
	})

	t.Run("special_case2", func(t *testing.T) {
		special_case2 := []int{2, 2}
		k4 := removeDuplicates_atMostTwice(special_case2)
		assert.Equal(t, 2, k4)
		assert.Equal(t, []int{2, 2}, special_case2)
	})

	t.Run("special_case3", func(t *testing.T) {
		special_case3 := []int{2, 3}
		k5 := removeDuplicates_atMostTwice(special_case3)
		assert.Equal(t, 2, k5)
		assert.Equal(t, []int{2, 3}, special_case3)
	})

}

func Test_88(t *testing.T) {
	num1 := []int{1, 2, 3, 0, 0, 0}
	merge(num1, 3, []int{2, 5, 6}, 3)
	assert.Equal(t, []int{1, 2, 2, 3, 5, 6}, num1)
}

func Test_121(t *testing.T) {
	t.Run("case1", func(t *testing.T) {
		case1 := []int{7, 1, 5, 3, 6, 4}
		assert.Equal(t, 5, maxProfit(case1))
	})

	t.Run("case2", func(t *testing.T) {
		case2 := []int{7, 6, 4, 3, 1}
		assert.Equal(t, 0, maxProfit(case2))
	})
}

func Test_122(t *testing.T) {
	t.Run("case1", func(t *testing.T) {
		case1 := []int{7, 1, 5, 3, 6, 4}
		assert.Equal(t, 7, maxProfit_DP(case1))
	})

	t.Run("case2", func(t *testing.T) {
		case2 := []int{7, 6, 4, 3, 1}
		assert.Equal(t, 0, maxProfit_DP(case2))
	})

	t.Run("case3", func(t *testing.T) {
		case3 := []int{1, 2, 3, 4, 5}
		assert.Equal(t, 4, maxProfit_DP(case3))
	})
}

func Test_134(t *testing.T) {
	t.Run("case1", func(t *testing.T) {
		gas1, cost1 := []int{6, 1, 4, 3, 5}, []int{3, 8, 2, 4, 2}
		assert.Equal(t, 2, canCompleteCircuit(gas1, cost1))
	})
}

func Test_135(t *testing.T) {
	t.Run("candyV1", func(t *testing.T) {
		rating1 := []int{1, 0, 2}
		assert.Equal(t, 5, candy(rating1))

		rating2 := []int{1, 2, 2}
		assert.Equal(t, 4, candy(rating2))
	})
	t.Run("candyV2", func(t *testing.T) {
		rating1 := []int{1, 0, 2}
		assert.Equal(t, 5, candyV2(rating1))

		rating2 := []int{1, 2, 2}
		assert.Equal(t, 4, candyV2(rating2))

		rating3 := []int{1, 3, 2, 2, 1}
		assert.Equal(t, 7, candyV2(rating3))

		rating4 := []int{1, 2, 3, 2, 2, 1}
		assert.Equal(t, 10, candyV2(rating4))

		rating5 := []int{1, 2, 3, 3, 2, 2, 1}
		assert.Equal(t, 12, candyV2(rating5))

		rating6 := []int{1, 6, 10, 8, 7, 3, 2}
		assert.Equal(t, 18, candyV2(rating6))
	})
}

func Test_151(t *testing.T) {
	t.Run("case1", func(t *testing.T) {
		case1 := "the sky is blue"
		assert.Equal(t, "blue is sky the", reverseWords(case1))
	})
}

func Test_189(t *testing.T) {
	t.Run("case1", func(t *testing.T) {
		case1 := []int{1, 2, 3, 4, 5, 6, 7}
		k := 3
		rotate(case1, k)
		assert.Equal(t, []int{5, 6, 7, 1, 2, 3, 4}, case1)
	})

	t.Run("case2", func(t *testing.T) {
		case2 := []int{-1}
		k := 2
		rotate(case2, k)
		assert.Equal(t, []int{-1}, case2)
	})

	t.Run("case3", func(t *testing.T) {
		case3 := []int{1, 2}
		k := 3
		rotate(case3, k)
		assert.Equal(t, []int{2, 1}, case3)
	})

	t.Run("case4", func(t *testing.T) {
		case4 := []int{1, 2, 3, 4, 5, 6}
		k := 1
		rotate(case4, k)
		assert.Equal(t, []int{6, 1, 2, 3, 4, 5}, case4)
	})

	t.Run("gcd_test_case", func(t *testing.T) {
		a, b := 15, 9
		c := gcd(a, b)
		assert.Equal(t, 3, c)

		d, e := 9, 15
		f := gcd(d, e)
		assert.Equal(t, 3, f)
	})

	t.Run("cycle_method_test_case1", func(*testing.T) {
		case1 := []int{1, 2, 3, 4, 5, 6, 7}
		k := 3
		rotate_cycle(case1, k)
		assert.Equal(t, []int{5, 6, 7, 1, 2, 3, 4}, case1)
	})

	t.Run("cycle_method_test_case2", func(t *testing.T) {
		case2 := []int{-1}
		k := 2
		rotate_cycle(case2, k)
		assert.Equal(t, []int{-1}, case2)
	})

	t.Run("cycle_method_test_case3", func(t *testing.T) {
		case3 := []int{1, 2}
		k := 3
		rotate_cycle(case3, k)
		assert.Equal(t, []int{2, 1}, case3)
	})

	t.Run("reverse_test_case1", func(t *testing.T) {
		case1 := []int{1, 2, 3, 4, 5, 6, 7}
		k := 3
		rotate_reverse(case1, k)
		assert.Equal(t, []int{5, 6, 7, 1, 2, 3, 4}, case1)
	})

	t.Run("reverse_test_case2", func(t *testing.T) {
		case2 := []int{-1}
		k := 2
		rotate_reverse(case2, k)
		assert.Equal(t, []int{-1}, case2)
	})

	t.Run("reverse_test_case2", func(t *testing.T) {
		case3 := []int{1, 2}
		k := 3
		rotate_reverse(case3, k)
		assert.Equal(t, []int{2, 1}, case3)
	})

}

func Test_274(t *testing.T) {
	t.Run("case1", func(t *testing.T) {
		case1 := []int{3, 0, 6, 1, 5}
		assert.Equal(t, 3, hIndex(case1))
	})

	t.Run("case2", func(t *testing.T) {
		case2 := []int{1, 3, 1}
		assert.Equal(t, 1, hIndex(case2))
	})

	t.Run("case3", func(t *testing.T) {
		case3 := []int{100}
		assert.Equal(t, 1, hIndex(case3))
	})

	t.Run("case4", func(t *testing.T) {
		case4 := []int{11, 15}
		assert.Equal(t, 2, hIndex(case4))
	})
}

func Test_31(t *testing.T) {
	t.Run("next permutation", func(t *testing.T) {
		nums := []int{1, 2, 3}
		nextPermutation(nums)
		assert.Equal(t, []int{1, 3, 2}, nums)
	})

	t.Run("next permutation", func(t *testing.T) {
		nums := []int{3, 2, 1}
		nextPermutation(nums)
		assert.Equal(t, []int{1, 2, 3}, nums)
	})

	t.Run("next permutation", func(t *testing.T) {
		nums := []int{1, 1, 5}
		nextPermutation(nums)
		assert.Equal(t, []int{1, 5, 1}, nums)
	})
}
