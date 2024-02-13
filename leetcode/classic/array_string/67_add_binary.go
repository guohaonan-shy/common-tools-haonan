package array_string

import (
	"strconv"
	"strings"
)

func addBinary(a string, b string) string {
	res := make([]string, 0)

	p1, p2 := len(a)-1, len(b)-1
	step := 0
	for p1 >= 0 || p2 >= 0 {
		v1, v2 := 0, 0
		if p1 >= 0 {
			v1 = int(a[p1] - '0')
			p1--
		}
		if p2 >= 0 {
			v2 = int(b[p2] - '0')
			p2--
		}

		digit := v1 + v2 + step
		step = digit / 2
		val := digit % 2
		res = append(res, strconv.Itoa(val))
	}
	if step == 1 {
		res = append(res, "1")
	}

	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res)-1-i] = res[len(res)-1-i], res[i]
	}

	return strings.Join(res, "")
}
