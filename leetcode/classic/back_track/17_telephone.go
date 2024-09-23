package back_track

var telephone = map[byte][]string{
	'2': {"a", "b", "c"},
	'3': {"d", "e", "f"},
	'4': {"g", "h", "i"},
	'5': {"j", "k", "l"},
	'6': {"m", "n", "o"},
	'7': {"p", "q", "r", "s"},
	'8': {"t", "u", "v"},
	'9': {"w", "x", "y", "z"},
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	vals := letterCombinations(digits[1:])
	if len(vals) == 0 {
		vals = append(vals, "")
	}

	options := telephone[digits[0]]
	ans := make([]string, 0)
	for _, option := range options {
		for _, val := range vals {
			ans = append(ans, option+val)
		}
	}
	return ans
}

func letterCombinationsV2(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	if len(digits) == 1 {
		return telephone[digits[0]]
	}

	cur := digits[0]
	combinations := letterCombinationsV2(digits[1:])

	res := make([]string, 0)
	for _, curChoice := range telephone[cur] {
		for _, combination := range combinations {
			res = append(res, curChoice+combination)
		}
	}
	return res
}
