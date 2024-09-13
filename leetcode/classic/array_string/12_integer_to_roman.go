package array_string

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

var divList = []int{1000, 500, 100, 50, 10, 5, 1}

func intToRomanV2(num int) string {
	res := ""
	val := 0
	for i := 0; i < len(divList); i++ {
		div := divList[i]

		val, num = num/div, num%div

		if val == 0 {
			continue
		}
		// because the value of num cannot exceed 4000
		if (div == 100 || div == 10 || div == 1) && val == 4 {
			res += dict[div]
			res += dict[divList[i-1]]
		} else if val == 1 && (div == 500 || div == 50 || div == 5) { // 9 or 90 or 900
			res += dict[divList[i+1]]
			res += dict[divList[i-1]]
			num = num % divList[i+1]
			i++
		} else {
			res += strings.Repeat(dict[div], val)
		}
	}
	return res
}
