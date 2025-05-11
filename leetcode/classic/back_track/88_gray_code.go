package back_track

func grayCode(n int) []int {
	temp := []int{0}
	res := make([][]int, 0)
	var dfs func(cur int)
	reached := make(map[int]bool, 0)
	dfs = func(cur int) {
		if len(temp) == 1<<n {
			// check diff between first and last element
			if checkLast(cur, n) {
				ans := make([]int, len(temp))
				copy(ans, temp)
				res = append(res, ans)
			}
			return
		}

		bits := trans2bit(cur, n)

		for i, bit := range bits {
			if bit == 0 {
				bits[i] = 1
			} else {
				bits[i] = 0
			}

			val := trans2integer(bits)
			if !reached[val] && len(res) < 1 {
				reached[val] = true
				temp = append(temp, val)
				dfs(val)
				temp = temp[:len(temp)-1]
				reached[val] = false
			}
			bits[i] = bit
		}
		return
	}
	reached[0] = true
	dfs(0)
	return res[0]
}

func trans2bit(num int, base int) []int {
	res := make([]int, base)
	idx := 0
	for num > 0 {
		ans := num % 2
		res[idx] = ans
		num = num / 2
		idx++
	}
	return res
}

func trans2integer(bits []int) int {
	res := 0
	for i, bit := range bits {
		if bit == 1 {
			res += 1 << i
		}
	}
	return res
}

func checkLast(num int, base int) bool {
	for i := 0; i <= base-1; i++ {
		if num == 1<<i {
			return true
		}
	}
	return false
}
