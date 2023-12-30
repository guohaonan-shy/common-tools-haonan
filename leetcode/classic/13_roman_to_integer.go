package classic

func romanToInt(s string) int {
	var dict = map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	if len(s) == 1 {
		return dict[s]
	}

	ans := 0
	for i := 0; i < len(s); {
		if i+1 < len(s) && dict[string(s[i])] < dict[string(s[i+1])] { // 不满足常规的从大到小，即存在4、9这样的特殊
			num1, num2 := dict[string(s[i])], dict[string(s[i+1])]
			num := num2 - num1
			ans += num
			i += 2
		} else {
			num := dict[string(s[i])]
			ans += num
			i++
		}
	}
	return ans
}
