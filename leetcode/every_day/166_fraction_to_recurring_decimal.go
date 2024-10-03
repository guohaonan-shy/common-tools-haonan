package every_day

import "strconv"

func fractionToDecimal(numerator int, denominator int) string {
	// can divide with no remain
	if numerator%denominator == 0 {
		return strconv.Itoa(numerator / denominator)
	}

	// 1. generate the sign
	res := ""
	if (numerator < 0) != (denominator < 0) { // this writing style seems like expression
		res += "-"
	}
	// all the steps are same with calculation between two positive number
	if denominator < 0 {
		denominator = -denominator
	}

	if numerator < 0 {
		numerator = -numerator
	}

	// 2. generate the integer part
	integer := numerator / denominator
	numerator = numerator % denominator

	res += strconv.Itoa(integer) + "."

	set := make(map[int]int, 0)

	for numerator != 0 && set[numerator] == 0 {
		set[numerator] = len(res)
		numerator *= 10
		digit := numerator / denominator
		res += strconv.Itoa(digit)
		numerator = numerator % denominator
	}

	if set[numerator] != 0 {
		last := set[numerator]
		// reconstruct in repeated
		res = res[:last] + "(" + res[last:] + ")"
	}
	return res
}
