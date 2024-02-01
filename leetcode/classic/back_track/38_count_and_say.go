package back_track

import "strconv"

func countAndSay(n int) string {

	if n == 1 {
		return "1"
	}

	if n == 2 {
		return "11"
	}

	str := countAndSay(n - 1)

	cnt := 1
	prev := str[0]
	res := ""
	for i := 1; i < len(str); i++ {
		if prev != str[i] {
			res += strconv.Itoa(cnt) + string(prev)
			cnt = 1
			prev = str[i]
		} else {
			cnt++
		}
	}
	res += strconv.Itoa(cnt) + string(prev)
	return res
}
