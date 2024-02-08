package back_track

import (
	"strconv"
	"strings"
)

func getPermutation(n int, k int) string {
	temp := make([]string, 0)
	state := make([]bool, n+1)
	num := 1
	var dfs func() string
	dfs = func() string {
		if len(temp) == n {
			// is the kth element
			if num == k {
				return strings.Join(temp, "")
			}
			num++
			return ""
		}

		for i := 1; i <= n; i++ {
			if !state[i] {
				state[i] = true
				temp = append(temp, strconv.Itoa(i))
				res := dfs()
				if res != "" {
					return res
				}
				temp = temp[:len(temp)-1]
				state[i] = false
			}
		}
		return ""
	}
	return dfs()
}

func getPermutation_standard(n, k int) string {

	factorialCache := make(map[int]int, n)
	pre := 1
	for i := 1; i <= n; i++ {
		val := pre * i
		factorialCache[i] = val
		pre = val
	}

	state := make([]bool, n+1)
	res := ""
	for i := 1; i < n && k > 1; i++ {

		choices := factorialCache[n-i]
		seq, subseq := k/choices, k%choices
		val, idx := 1, 1
		if subseq == 0 {
			seq--
			subseq = choices
		}

		for seq >= 0 {
			if !state[idx] {
				val = idx
				seq--
			}
			idx++
		}
		res += strconv.Itoa(val)
		state[val] = true
		k = subseq
	}

	for i := 1; i <= n; i++ {
		if !state[i] {
			res += strconv.Itoa(i)
		}
	}

	return res
}
