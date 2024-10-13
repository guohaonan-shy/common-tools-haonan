package every_day

import (
	"fmt"
	"strconv"
)

func getLucky(s string, k int) int {
	convertRes := convert2integer(s)

	for i := 0; i < k; i++ {
		convertRes = transform(convertRes)
	}
	res, _ := strconv.Atoi(convertRes)
	return res
}

func convert2integer(s string) string {
	res := ""
	for i := range s {
		local := s[i]
		res += strconv.Itoa(int(local - 'a' + 1))
	}
	return res
}

func transform(s string) string {
	res := 0
	for i := range s {
		res += int(s[i] - '0')
	}
	return strconv.Itoa(res)
}

func iterateString(s string) {
	for i, str := range s {
		println(fmt.Sprintf("interate by index %v", i))
		println(s[i])
		println(str)
		println("####################")
	}
}

func iterateStringByStr(s string) {
	for _, str := range s {
		println("interate by str")
		println(str)
		println("####################")
	}
}
