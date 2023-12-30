package classic

import "strings"

var dict = map[int]string{
	1000: "M",
	500:  "D",
	100:  "C",
	50:   "L",
	10:   "X",
	5:    "V",
	1:    "I",
}

var divdict = []int{1000, 100, 10, 1}

func intToRoman(num int) string {
	ans := ""
	for num != 0 {
		divend, div := 0, 0
		for _, d := range divdict {
			div = num / d
			if div != 0 {
				divend = d
				break
			}
		}

		var str = ""
		if div > 5 {
			if div == 9 {
				str = dict[divend] + dict[divend*10]
			} else {
				str = dict[divend*5] + strings.Repeat(dict[divend], div-5)
			}
			ans += str
		} else if div < 5 {
			if div == 4 {
				str = dict[divend] + dict[divend*5]
			} else {
				str = strings.Repeat(dict[divend], div)
			}
			ans += str
		} else {
			ans += dict[divend*5]
		}
		num = num % divend
	}
	return ans
}
