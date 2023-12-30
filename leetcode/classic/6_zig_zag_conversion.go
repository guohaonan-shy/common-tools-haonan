package classic

import "strings"

func convert(s string, numRows int) string {
	temp := make([]string, numRows)
	cur, step := 0, 1

	for _, v := range s {
		str := string(v)
		temp[cur] += str
		if cur+step >= numRows || cur+step < 0 {
			step = -step
		}
		cur += step
	}

	return strings.Join(temp, "")
}
