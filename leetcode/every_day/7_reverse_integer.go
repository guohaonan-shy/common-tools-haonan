package every_day

func reverse(x int) int {

	res := 0

	for x != 0 {
		val := x % 10

		if val != 0 || res != 0 {
			res = 10*res + val
		}

		x = x / 10
	}

	if x < 0 {
		res = -res
	}
	a, b := (1<<31)-1, -(1 << 31)
	if res > a || res < b {
		res = 0
	}

	return res

}
