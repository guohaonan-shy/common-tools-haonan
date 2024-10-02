package every_day

import "strconv"

func fractionToDecimal(numerator int, denominator int) string {
	// 1. generate the sign
	res := ""
	if numerator*denominator < 0 {
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

	if numerator == 0 {
		return res + strconv.Itoa(integer)
	}

	res += strconv.Itoa(integer) + "."
	posOfDot := len(res)

	set := make(map[int]int, 0)

	for numerator != 0 {
		numerator *= 10
		if last, ok := set[numerator]; ok {
			// reconstruct in repeated
			res = res[:last] + "(" + res[last:] + ")"
			break
		}

		digit := numerator / denominator
		res += strconv.Itoa(digit)
		set[numerator] = posOfDot

		posOfDot++
		numerator = numerator % denominator
	}
	return res
}
