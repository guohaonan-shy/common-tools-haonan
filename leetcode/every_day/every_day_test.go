package every_day

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_7(t *testing.T) {
	type args struct {
		input int
	}

	testCases := []struct {
		name   string
		args   args
		wanted int
	}{
		{
			name: "test case 1",
			args: args{
				input: 123,
			},
			wanted: 321,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, testCase.wanted, reverse(testCase.args.input))
		})
	}
}

func Test_43(t *testing.T) {
	t.Run("case1", func(t *testing.T) {
		res := multiply("2", "3")
		assert.Equal(t, "6", res)
	})

	t.Run("case2", func(t *testing.T) {
		res := multiply("123", "456")
		assert.Equal(t, "56088", res)
	})

	t.Run("case3", func(t *testing.T) {
		res := multiply("9", "9")
		assert.Equal(t, "81", res)
	})

	t.Run("case4", func(t *testing.T) {
		res := multiply("498828660196", "840477629533")
		assert.Equal(t, "419254329864656431168468", res)
	})

	t.Run("case5", func(t *testing.T) {
		res := multiply("140", "721")
		assert.Equal(t, "100940", res)
	})
}

func Test_115(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		res := numDistinct("rabbbit", "rabbit")
		assert.Equal(t, 3, res)
	})

	t.Run("case 2", func(t *testing.T) {
		res := numDistinct("babgbag", "bag")
		assert.Equal(t, 5, res)
	})

	t.Run("case 3", func(t *testing.T) {
		res := numDistinct("aabb", "abb")
		assert.Equal(t, 2, res)
	})
}

func Test_134(t *testing.T) {
	t.Run("gas station::has solution::case1", func(t *testing.T) {
		gas, cost := []int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2}
		assert.Equal(t, 3, canCompleteCircuit(gas, cost))
	})

	t.Run("gas station::no solution::case1", func(t *testing.T) {
		gas, cost := []int{2, 3, 4}, []int{3, 4, 3}
		assert.Equal(t, -1, canCompleteCircuit(gas, cost))
	})
}

func Test_166(t *testing.T) {
	t.Run("fraction case1", func(t *testing.T) {
		numerator, denominator := 1, 2
		assert.Equal(t, "0.5", fractionToDecimal(numerator, denominator))
	})

	t.Run("fraction case2", func(t *testing.T) {
		numerator, denominator := 2, 1
		assert.Equal(t, "2", fractionToDecimal(numerator, denominator))
	})

	t.Run("fraction case3", func(t *testing.T) {
		numerator, denominator := 4, 333
		assert.Equal(t, "0.(012)", fractionToDecimal(numerator, denominator))
	})

	t.Run("fraction case4", func(t *testing.T) {
		numerator, denominator := 1, 6
		assert.Equal(t, "0.1(6)", fractionToDecimal(numerator, denominator))
	})

	t.Run("fraction case5", func(t *testing.T) {
		numerator, denominator := -50, 8
		assert.Equal(t, "-6.25", fractionToDecimal(numerator, denominator))
	})

	t.Run("fraction case6", func(t *testing.T) {
		numerator, denominator := 7, -12
		assert.Equal(t, "-0.58(3)", fractionToDecimal(numerator, denominator))
	})

	t.Run("corner case 1", func(t *testing.T) {
		numerator, denominator := -2147483648, 1
		assert.Equal(t, "-2147483648", fractionToDecimal(numerator, denominator))
	})
}

func Test_286(t *testing.T) {
	t.Run("case1", func(t *testing.T) {
		rooms := [][]int{
			{2147483647, -1, 0, 2147483647},
			{2147483647, 2147483647, 2147483647, -1},
			{2147483647, -1, 2147483647, -1},
			{0, -1, 2147483647, 2147483647},
		}
		wallsAndGates(rooms)
		assert.Equal(t, [][]int{
			{3, -1, 0, 1},
			{2, 2, 1, -1},
			{1, -1, 2, -1},
			{0, -1, 3, 4},
		}, rooms)
	})
}

func Test_650(t *testing.T) {
	t.Run("650", func(t *testing.T) {
		assert.Equal(t, 0, minSteps(1))
		assert.Equal(t, 6, minSteps(9))
		assert.Equal(t, 3, minSteps(3))
	})
}

func Test_688(t *testing.T) {
	t.Run("knight", func(t *testing.T) {
		assert.Equal(t, 0.06250, knightProbability(3, 2, 0, 0))
	})
}

func Test_871(t *testing.T) {
	t.Run("refueling stop::case1", func(t *testing.T) {
		assert.Equal(t, 0, minRefuelStops(1, 1, [][]int{}))
		assert.Equal(t, 0, minRefuelStopDP(1, 1, [][]int{}))
	})

	t.Run("refueling stop::case2", func(t *testing.T) {
		assert.Equal(t, -1, minRefuelStops(100, 1, [][]int{{10, 100}}))
		assert.Equal(t, -1, minRefuelStopDP(100, 1, [][]int{{10, 100}}))
	})

	t.Run("refueling stop::case3", func(t *testing.T) {
		stations := [][]int{
			{10, 60},
			{20, 30},
			{30, 30},
			{60, 40},
		}
		assert.Equal(t, 2, minRefuelStops(100, 10, stations))
		assert.Equal(t, 2, minRefuelStopDP(100, 10, stations))
	})
}

func Test_887(t *testing.T) {
	t.Run("super_egg_drop case1", func(t *testing.T) {
		assert.Equal(t, 2, superEggDrop(1, 2))
	})

	t.Run("super_egg_drop case2", func(t *testing.T) {
		assert.Equal(t, 3, superEggDrop(2, 6))
	})

	t.Run("super_egg_drop case3", func(t *testing.T) {
		assert.Equal(t, 4, superEggDrop(3, 14))
	})
}

func Test_983(t *testing.T) {
	t.Run("minimum travel ticket case1", func(t *testing.T) {
		assert.Equal(t, 11, minimumTravelCost([]int{1, 4, 6, 7, 8, 20}, []int{2, 7, 15}))
	})

	t.Run("minimum travel ticket case2", func(t *testing.T) {
		assert.Equal(t, 17, minimumTravelCost([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 30, 31}, []int{2, 7, 15}))
	})

	t.Run("minimum travel ticket case2", func(t *testing.T) {
		assert.Equal(t, 6, minimumTravelCost([]int{1, 4, 6, 7, 8, 20}, []int{7, 2, 15}))
	})

	t.Run("minimum travel ticket case3", func(t *testing.T) {
		assert.Equal(t, 50, minimumTravelCost([]int{1, 2, 3, 4, 6, 8, 9, 10, 13, 14, 16, 17, 19, 21, 24, 26, 27, 28, 29}, []int{3, 14, 50}))
	})
}

func Test_1570(t *testing.T) {
	t.Run("dot product", func(t *testing.T) {
		v1 := Constructor([]int{1, 0, 0, 2, 3})
		v2 := Constructor([]int{0, 3, 0, 4, 0})

		assert.Equal(t, 8, v1.dotProduct(v2))
	})
}

func Test_1870(t *testing.T) {
	t.Run("minimum speed to take the trains case1", func(t *testing.T) {
		dist := []int{1, 3, 2}
		hour := 6.0
		assert.Equal(t, 1, minSpeedOnTime(dist, hour))
	})

	t.Run("minimum speed to take the trains case2", func(t *testing.T) {
		dist := []int{1, 3, 2}
		hour := 2.7
		assert.Equal(t, 3, minSpeedOnTime(dist, hour))
	})

	t.Run("minimum speed to take the trains case3", func(t *testing.T) {
		dist := []int{1, 3, 2}
		hour := 1.9
		assert.Equal(t, -1, minSpeedOnTime(dist, hour))
	})

	t.Run("minimum speed to take the trains: extreme case1", func(t *testing.T) {
		dist := []int{1, 1, 100000}
		hour := 2.01
		assert.Equal(t, 10000000, minSpeedOnTime(dist, hour))
	})
}

func Test_1928(t *testing.T) {
	t.Run("case1", func(t *testing.T) {
		edges := [][]int{
			{0, 1, 10},
			{1, 2, 10},
			{2, 5, 10},
			{0, 3, 1},
			{3, 4, 10},
			{4, 5, 15},
		}
		passingFees := []int{5, 1, 2, 20, 20, 3}
		assert.Equal(t, 11, minCostDP(30, edges, passingFees))
	})

	t.Run("case2", func(t *testing.T) {
		edges := [][]int{
			{0, 1, 10},
			{1, 2, 10},
			{2, 5, 10},
			{0, 3, 1},
			{3, 4, 10},
			{4, 5, 15},
		}
		passingFees := []int{5, 1, 2, 20, 20, 3}
		assert.Equal(t, 48, minCostDP(29, edges, passingFees))
	})

	t.Run("case3", func(t *testing.T) {
		edges := [][]int{
			{0, 1, 10},
			{1, 2, 10},
			{2, 5, 10},
			{0, 3, 1},
			{3, 4, 10},
			{4, 5, 15},
		}
		passingFees := []int{5, 1, 2, 20, 20, 3}
		assert.Equal(t, -1, minCostDP(25, edges, passingFees))
	})

}

func Test_1945(t *testing.T) {
	t.Run("case1", func(t *testing.T) {
		assert.Equal(t, 36, getLucky("iiii", 1))
	})

	t.Run("s[i] vs str", func(t *testing.T) {
		iterateString("aaa过好难")

		iterateStringByStr("aaa过好难")
	})
}

func Test_2024(t *testing.T) {
	t.Run("maximum consecutive confusion::case1", func(t *testing.T) {
		assert.Equal(t, 4, maxConsecutiveAnswers("TTFF", 2))
	})

	t.Run("maximum consecutive confusion::case2", func(t *testing.T) {
		assert.Equal(t, 3, maxConsecutiveAnswers("TTFF", 1))
	})

	t.Run("maximum consecutive confusion::case3", func(t *testing.T) {
		assert.Equal(t, 5, maxConsecutiveAnswers("TTFTTFTT", 1))
	})
}

func Test_2145(t *testing.T) {
	type args struct {
		differences  []int
		lower, upper int
	}

	testCases := []struct {
		name   string
		args   args
		wanted int
	}{
		{
			name: "case 1",
			args: args{
				differences: []int{1, -3, 4},
				lower:       1,
				upper:       6,
			},
			wanted: 2,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res := numberOfArrays(testCase.args.differences, testCase.args.lower, testCase.args.upper)
			assert.Equal(t, testCase.wanted, res)
		})
	}
}

func Test_2187(t *testing.T) {
	t.Run("minimum time::case1", func(t *testing.T) {
		assert.Equal(t, int64(3), minimumTime_standard([]int{1, 2, 3}, 5))
	})

	t.Run("minimum time::case2", func(t *testing.T) {
		assert.Equal(t, int64(2), minimumTime_standard([]int{2}, 1))
	})

	t.Run("minimum time::case3", func(t *testing.T) {
		assert.Equal(t, int64(25), minimumTime_standard([]int{5, 10, 10}, 9))
	})
}

func Test_2304(t *testing.T) {
	t.Run("minimum path in the grid case1", func(t *testing.T) {
		grid := [][]int{
			{5, 3},
			{4, 0},
			{2, 1},
		}

		moveCost := [][]int{
			{9, 8},
			{1, 5},
			{10, 12},
			{18, 6},
			{2, 4},
			{14, 3},
		}

		assert.Equal(t, 17, minimumPathInGridDP(grid, moveCost))
	})

	t.Run("minimum path in the grid case2", func(t *testing.T) {
		grid := [][]int{
			{5, 1, 2},
			{4, 0, 3},
		}

		moveCost := [][]int{
			{12, 10, 15},
			{20, 23, 8},
			{21, 7, 1},
			{8, 1, 13},
			{9, 10, 25},
			{5, 3, 2},
		}

		assert.Equal(t, 6, minimumPathInGridDP(grid, moveCost))
	})
}

func Test_3171(t *testing.T) {
	t.Run("case1", func(t *testing.T) {
		case1 := []int{1, 2, 3, 4}
		k := 3
		assert.Equal(t, 0, minimumDifference(case1, k))
	})

	t.Run("case2", func(t *testing.T) {
		case2 := []int{1, 3, 1, 3}
		k := 2
		assert.Equal(t, 1, minimumDifference(case2, k))
	})

	t.Run("case3", func(t *testing.T) {
		case3 := []int{1}
		k := 10
		assert.Equal(t, 9, minimumDifference(case3, k))
	})

	t.Run("case4", func(t *testing.T) {
		case4 := []int{6}
		k := 2
		assert.Equal(t, 4, minimumDifference(case4, k))
	})

	t.Run("case5", func(t *testing.T) {
		case5 := []int{1, 10}
		k := 6
		assert.Equal(t, 4, minimumDifference(case5, k))
	})

	t.Run("case6", func(t *testing.T) {
		case6 := []int{1, 8, 9}
		k := 7
		assert.Equal(t, 1, minimumDifference(case6, k))
	})
}
