package array_string

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

func convertV2(s string, numRows int) string {
	// numRows = 1 -> div = 0; extreme case we need to process specifically
	if numRows == 1 {
		return s
	}
	div := 2*numRows - 2

	res := make([]string, numRows)
	for i := range s {
		idx := min(i%div, div-i%div)
		res[idx] += string(s[i])
	}
	return strings.Join(res, "")
}
